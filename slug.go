// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package slug

import (
	"bytes"
	"regexp"
	"sort"
	"strings"

	"github.com/rainycape/unidecode"
)

var (
	// CustomSub stores custom substitution map
	CustomSub map[string]string
	// CustomRuneSub stores custom rune substitution map
	CustomRuneSub map[rune]string

	// MaxLength stores maximum slug length.
	// It's smart so it will cat slug after full word.
	// By default slugs aren't shortened.
	// If MaxLength is smaller than length of the first word, then returned
	// slug will contain only substring from the first word truncated
	// after MaxLength.
	MaxLength int

	// Separator configure default separator for all slugs.
	// By default set to '-'.
	Separator = '-'

	// BUG: Not possible to change separator from default.
	regexpNonAuthorizedChars = regexp.MustCompile("[^a-z0-9-_" + string(Separator) + "]")
	regexpMultipleSeparators = regexp.MustCompile(string(Separator) + "+")
)

//=============================================================================

// Make returns slug generated from provided string. Will use "en" as language
// substitution.
func Make(s string) (slug string) {
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
	case "de", "deu":
		slug = SubstituteRune(slug, deSub)
	case "en", "eng":
		slug = SubstituteRune(slug, enSub)
	case "es", "spa":
		slug = SubstituteRune(slug, esSub)
	case "fi", "fin":
		slug = SubstituteRune(slug, fiSub)
	case "gr", "el", "ell":
		slug = SubstituteRune(slug, grSub)
	case "nl", "nld":
		slug = SubstituteRune(slug, nlSub)
	case "pl", "pol":
		slug = SubstituteRune(slug, plSub)
	case "tr", "tur":
		slug = SubstituteRune(slug, trSub)
	default: // fallback to "en" if lang not found
		slug = SubstituteRune(slug, enSub)
	}

	// Process all non ASCII symbols
	slug = unidecode.Unidecode(slug)

	slug = strings.ToLower(slug)

	// Process all remaining symbols
	slug = regexpNonAuthorizedChars.ReplaceAllString(slug, string(Separator))
	slug = regexpMultipleSeparators.ReplaceAllString(slug, string(Separator))

	if MaxLength > 0 {
		slug = smartTruncate(slug)
	}

	slug = strings.Trim(slug, "-_"+string(Separator))

	return slug
}

// Substitute returns string with superseded all substrings from
// provided substitution map. Substitution map will be applied in alphabetic
// order. Many passes, on one substitution another one could apply.
func Substitute(s string, sub map[string]string) (buf string) {
	buf = s
	var keys []string
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
	if len(text) < MaxLength {
		return text
	}

	var truncated string
	words := strings.SplitAfter(text, "-")
	// If MaxLength is smaller than length of the first word return word
	// truncated after MaxLength.
	if len(words[0]) > MaxLength {
		return words[0][:MaxLength]
	}
	for _, word := range words {
		if len(truncated)+len(word)-1 <= MaxLength {
			truncated = truncated + word
		} else {
			break
		}
	}
	return truncated
}

// IsSlug returns True if provided text does not contain white characters,
// punctuation, all letters are lower case and only from ASCII range.
// It could contain `-` and `_` but not at the beginning or end of the text.
// It should be in range of the MaxLength var if specified.
// All output from slug.Make(text) should pass this test.
func IsSlug(text string) bool {
	if text == "" ||
		(MaxLength > 0 && len(text) > MaxLength) ||
		strings.HasPrefix(text, "-") || strings.HasPrefix(text, "_") ||
		strings.HasPrefix(text, string(Separator)) ||
		strings.HasSuffix(text, "-") || strings.HasSuffix(text, "_") ||
		strings.HasSuffix(text, string(Separator)) {
		return false
	}
	for _, c := range text {
		if (c < 'a' || c > 'z') && (c < '0' || c > '9') &&
			c != '-' && c != '_' && c != Separator {
			return false
		}
	}
	return true
}
