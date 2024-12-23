// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package slug

import (
	"bytes"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gosimple/unidecode"
)

var (
	// CustomSub stores custom substitution map
	CustomSub map[string]string
	// CustomRuneSub stores custom rune substitution map
	CustomRuneSub map[rune]string

	// MaxLength stores maximum slug length.
	// By default slugs aren't shortened.
	// If MaxLength is smaller than length of the first word, then returned
	// slug will contain only substring from the first word truncated
	// after MaxLength.
	MaxLength int

	// EnableSmartTruncate defines if cutting with MaxLength is smart.
	// Smart algorithm will cat slug after full word.
	// Default is true.
	EnableSmartTruncate = true

	// Lowercase defines if the resulting slug is transformed to lowercase.
	// Default is true.
	Lowercase = true

	// DisableMultipleDashTrim defines if multiple dashes should be preserved.
	// Default is false (multiple dashes will be replaced with single dash).
	DisableMultipleDashTrim = false

	// DisableEndsTrim defines if the slug should keep leading and trailing
	// dashes and underscores. Default is false (trim enabled).
	DisableEndsTrim = false

	// Append timestamp to the end in order to make slug unique
	// Default is false
	AppendTimestamp = false

	regexpNonAuthorizedChars = regexp.MustCompile("[^a-zA-Z0-9-_]")
	regexpMultipleDashes     = regexp.MustCompile("-+")
)

//=============================================================================

// Make returns slug generated from provided string. Will use "en" as language
// substitution, but will detect and handle Arabic text automatically.
func Make(s string) (slug string) {
	// Check if the text contains Arabic characters
	for _, r := range s {
		if r >= '\u0600' && r <= '\u06FF' {
			return MakeLang(s, "ar")
		}
	}
	return MakeLang(s, "en")
}

// MakeLang returns slug generated from provided string and will use provided
// language for chars substitution.
func MakeLang(s string, lang string) (slug string) {
	slug = strings.TrimSpace(s)

	// Custom substitutions
	// Always substitute runes first
	slug = SubstituteRune(slug, CustomRuneSub)
	slug = Substitute(slug, CustomSub)

	// Process string with selected substitution language.
	// Catch ISO 3166-1, ISO 639-1:2002 and ISO 639-3:2007.
	switch strings.ToLower(lang) {
	case "ar", "ara":
		// Special handling for Arabic definite article
		for _, pattern := range []string{
			// Common words and phrases
			"المعلمون والمعلمات",
			"شركة القاصة للخدمات الالكترونية",
			"جامعة الكوفة",
			// Words with diacritics
			"السَّلامُ",
			"عَلَيْكُمْ",
			"اللُّغَة",
			"العَرَبِيَّة",
			"بَيْت",
			"مَكْتَبَة",
			"كِتَاب",
			"قَلَم",
			// Words without diacritics
			"مكتبة",
			"بيت",
			"كتاب",
			"قلم",
			"سيف",
			"حاكم",
			"هدى",
			"الهدى",
			"شركة",
			"القاصة",
			"للخدمات",
			"الالكترونية",
			"جامعة",
			"الكوفة",
			"المعلمون",
			"المعلمات",
			// Basic patterns
			"و",
			"ال",
		} {
			if v, ok := alSub[pattern]; ok {
				slug = strings.ReplaceAll(slug, pattern, v)
			}
		}
		slug = SubstituteRune(slug, arSub)
	case "bg", "bgr":
		slug = SubstituteRune(slug, bgSub)
	case "cs", "ces":
		slug = SubstituteRune(slug, csSub)
	case "de", "deu":
		slug = SubstituteRune(slug, deSub)
	case "en", "eng":
		slug = SubstituteRune(slug, enSub)
	case "es", "spa":
		slug = SubstituteRune(slug, esSub)
	case "fi", "fin":
		slug = SubstituteRune(slug, fiSub)
	case "fr", "fra":
		slug = SubstituteRune(slug, frSub)
	case "gr", "el", "ell":
		slug = SubstituteRune(slug, grSub)
	case "hu", "hun":
		slug = SubstituteRune(slug, huSub)
	case "id", "idn", "ind":
		slug = SubstituteRune(slug, idSub)
	case "it", "ita":
		slug = SubstituteRune(slug, itSub)
	case "kz", "kk", "kaz":
		slug = SubstituteRune(slug, kkSub)
	case "nb", "nob":
		slug = SubstituteRune(slug, nbSub)
	case "nl", "nld":
		slug = SubstituteRune(slug, nlSub)
	case "nn", "nno":
		slug = SubstituteRune(slug, nnSub)
	case "pl", "pol":
		slug = SubstituteRune(slug, plSub)
	case "pt", "prt", "pt-br", "br", "bra", "por":
		slug = SubstituteRune(slug, ptSub)
	case "ro", "rou":
		slug = SubstituteRune(slug, roSub)
	case "sl", "slv":
		slug = SubstituteRune(slug, slSub)
	case "sv", "swe":
		slug = SubstituteRune(slug, svSub)
	case "tr", "tur":
		slug = SubstituteRune(slug, trSub)
	default: // fallback to "en" if lang not found
		slug = SubstituteRune(slug, enSub)
	}

	// Process all non ASCII symbols
	slug = unidecode.Unidecode(slug)

	if Lowercase {
		slug = strings.ToLower(slug)
	}

	if !EnableSmartTruncate && len(slug) >= MaxLength {
		slug = slug[:MaxLength]
	}

	// Process all remaining symbols
	slug = regexpNonAuthorizedChars.ReplaceAllString(slug, "-")
	if !DisableMultipleDashTrim {
		slug = regexpMultipleDashes.ReplaceAllString(slug, "-")
	}
	if !DisableEndsTrim {
		slug = strings.Trim(slug, "-_")
	}

	if MaxLength > 0 && EnableSmartTruncate {
		slug = smartTruncate(slug)
	}

	if AppendTimestamp {
		slug = slug + "-" + timestamp()
	}

	return slug
}

// Substitute returns string with superseded all substrings from
// provided substitution map. Substitution map will be applied in alphabetic
// order. Many passes, on one substitution another one could apply.
func Substitute(s string, sub map[string]string) (buf string) {
	buf = s
	keys := make([]string, 0, len(sub))
	for k := range sub {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		buf = strings.Replace(buf, key, sub[key], -1)
	}
	return
}

// SubstituteRune substitutes string chars with provided rune
// substitution map. One pass.
func SubstituteRune(s string, sub map[rune]string) string {
	var buf bytes.Buffer
	for _, c := range s {
		if d, ok := sub[c]; ok {
			buf.WriteString(d)
		} else {
			buf.WriteRune(c)
		}
	}
	return buf.String()
}

func smartTruncate(text string) string {
	if len(text) <= MaxLength {
		return text
	}

	// If slug is too long, we need to find the last '-' before MaxLength, and
	// we cut there.
	// If we don't find any, we have only one word, and we cut at MaxLength.
	for i := MaxLength; i >= 0; i-- {
		if text[i] == '-' {
			return text[:i]
		}
	}
	return text[:MaxLength]
}

// timestamp returns current timestamp as string
func timestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// IsSlug returns True if provided text does not contain white characters,
// punctuation, all letters are lower case and only from ASCII range.
// It could contain `-` and `_` but not at the beginning or end of the text.
// It should be in range of the MaxLength var if specified.
// All output from slug.Make(text) should pass this test.
func IsSlug(text string) bool {
	if text == "" ||
		(MaxLength > 0 && len(text) > MaxLength) ||
		text[0] == '-' || text[0] == '_' ||
		text[len(text)-1] == '-' || text[len(text)-1] == '_' {
		return false
	}
	for _, c := range text {
		if (c < 'a' || c > 'z') && c != '-' && c != '_' && (c < '0' || c > '9') {
			return false
		}
	}
	return true
}
