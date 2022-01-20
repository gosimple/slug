// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package slug

import (
	"bytes"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/gosimple/unidecode"
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

	// Lowercase defines if the resulting slug is transformed to lowercase.
	// Default is true.
	Lowercase = true

	regexpNonAuthorizedChars *regexp.Regexp
	regexpMultipleSeparators *regexp.Regexp
	regexpTrimSeparators     *regexp.Regexp
	regexpIsSlug             *regexp.Regexp

	separator string
)

// SetSeparator sets the separator that's used between "words"
func SetSeparator(s string) {
	separator = s
	quotedSeparator := regexp.QuoteMeta(s)

	regexpNonAuthorizedChars = regexp.MustCompile("[^a-zA-Z0-9-_]")

	// let the separator be 0 length, in the event of no separator being wanted
	// in this case we can set it to "nil" and check that it's not before use laster
	if len(s) != 0 {
		regexpMultipleSeparators = regexp.MustCompile(fmt.Sprintf("(%s)+", quotedSeparator))
		regexpTrimSeparators = regexp.MustCompile(fmt.Sprintf("^(%s)+|(%s)+$", quotedSeparator, quotedSeparator))
		regexpIsSlug = regexp.MustCompile(fmt.Sprintf("^(?:%s)|(?:%s)$|(?:%s)(?:%s)|([^a-z0-9-_]+)",
			quotedSeparator, quotedSeparator, quotedSeparator, quotedSeparator))
	} else {
		regexpMultipleSeparators = nil
		regexpTrimSeparators = nil
		regexpIsSlug = nil
	}
}

// GetSeparator returns the current separator
func GetSeparator() string {
	return separator
}

func init() {
	SetSeparator("-")
}

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

	// Process all remaining symbols
	// Break slug by separator so we don't replace sep chars
	// if the chars are not in regexpNonAuthorizedChars, while
	// also not considring all sep chars safe (make sure the
	// separator is only being considered as is)
	slugParts := strings.Split(slug, separator)
	for i, p := range slugParts {
		slugParts[i] = regexpNonAuthorizedChars.ReplaceAllString(p, separator)
	}
	slug = strings.Join(slugParts, separator)

	// if the separator is nothing then we don't need to remove/trim nothing
	if len(separator) != 0 {
		slug = regexpMultipleSeparators.ReplaceAllString(slug, separator)
		slug = regexpTrimSeparators.ReplaceAllString(slug, "")
	}
	slug = strings.Trim(slug, "-_")

	if MaxLength > 0 {
		slug = smartTruncate(slug)
	}

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

	// check if the separator is empty or not, if empty
	// we can do a simple slice of the string
	if len(separator) != 0 {
		var truncated string
		words := strings.SplitAfter(text, separator)

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
		truncated = regexpTrimSeparators.ReplaceAllString(truncated, "")
		return strings.Trim(truncated, "-_")
	}

	return text[:MaxLength]
}

// IsSlug returns True if provided text does not contain anything that
// isn't a-z, A-Z, 0-9, dashes, underscores, or the separator strings.
// It can't contain `-` and `_` at the beginning or end of the text.
// It could contain the separator, but not at the beginning or end of the text,
// or multiple consecutive separators.
// It should be in range of the MaxLength var if specified.
// All output from slug.Make(text) should pass this test.
func IsSlug(text string) bool {
	if text == "" ||
		(MaxLength > 0 && len(text) > MaxLength) ||
		text[0] == '-' || text[0] == '_' ||
		text[len(text)-1] == '-' || text[len(text)-1] == '_' {
		return false
	}

	if len(separator) == 0 {
		return true
	}

	matches := regexpIsSlug.FindAllStringSubmatch(text, -1)

	// no matches mean we didn't find anything we know to be bad
	if matches == nil {
		return true
	}

	// check if the 2nd, `m[1]`, match is equal to the separator
	// if it is, then we're fine, but if it isn't, then our
	// alphanum-dash-underscore check found something bad
	for _, m := range matches {
		if m[1] != separator {
			return false
		}
	}

	return true
}
