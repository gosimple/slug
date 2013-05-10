// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package slug

import (
	"github.com/fiam/gounidecode/unidecode"
	"regexp"
	"strings"
)

var (
	// Custom substitution map
	CustomSub map[rune]string
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

	// Select substitution language
	switch lang {
	case "de":
		slug = substitute(slug, deSub)
	case "en":
		slug = substitute(slug, enSub)
	case "pl":
		slug = substitute(slug, plSub)
	default: // fallback to "en" if lang not found
		slug = substitute(slug, enSub)
	}

	slug = substitute(slug, defaultSub)
	slug = substitute(slug, CustomSub)

	slug = unidecode.Unidecode(slug)

	slug = strings.ToLower(slug)

	slug = regexp.MustCompile("[^a-z0-9-_]").ReplaceAllString(slug, "-")
	slug = regexp.MustCompile("-+").ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	return slug
}

// Substitute string chars with provided substitution map.
func substitute(s string, sub map[rune]string) (buf string) {
	for _, c := range s {
		if d, ok := sub[c]; ok {
			buf += d
		} else {
			buf += string(c)
		}
	}
	return
}
