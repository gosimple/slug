// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package slug

import (
	"testing"
)

//=============================================================================

var SlugMakeTests = []struct {
	in  string
	out string
}{
	{"DOBROSLAWZYBORT", "dobroslawzybort"},
	{"Dobroslaw Zybort", "dobroslaw-zybort"},
	{"  Dobroslaw     Zybort  ?", "dobroslaw-zybort"},
	{"Dobrosław Żybort", "dobroslaw-zybort"},
	{"Ala ma 6 kotów.", "ala-ma-6-kotow"},

	{"áÁàÀãÃâÂäÄąĄą̊Ą̊", "aaaaaaaaaaaaaa"},
	{"ćĆĉĈçÇ", "cccccc"},
	{"éÉèÈẽẼêÊëËęĘ", "eeeeeeeeeeee"},
	{"íÍìÌĩĨîÎïÏįĮ", "iiiiiiiiiiii"},
	{"łŁ", "ll"},
	{"ńŃ", "nn"},
	{"óÓòÒõÕôÔöÖǫǪǭǬø", "ooooooooooooooo"},
	{"śŚ", "ss"},
	{"úÚùÙũŨûÛüÜųŲ", "uuuuuuuuuuuu"},
	{"y̨Y̨", "yy"},
	{"źŹżŹ", "zzzz"},
	{"·/,:;`˜'\"", ""},
	{"2000–2013", "2000-2013"},
	{"style—not", "style-not"},
	{"test_slug", "test_slug"},
	{"Æ", "ae"},
	{"Ich heiße", "ich-heisse"},

	{"This & that", "this-and-that"},
	{"fácil €", "facil-eu"},
	{"smile ☺", "smile"},
	{"Hellö Wörld хелло ворлд", "hello-world-khello-vorld"},
	{"\"C'est déjà l’été.\"", "cest-deja-lete"},
	{"jaja---lol-méméméoo--a", "jaja-lol-mememeoo-a"},
	{"影師", "ying-shi"},
}

func TestSlugMake(t *testing.T) {
	for index, st := range SlugMakeTests {
		slug := Make(st.in)
		if st.out != slug {
			t.Errorf(
				"%d. Make(%q) => out = %q, want %q",
				index, st.in, slug, st.out)
		}
	}
}

var SlugMakeLangTests = []struct {
	lang string
	in   string
	out  string
}{
	{"en", "This & that", "this-and-that"},
	{"de", "This & that", "this-und-that"},
	{"pl", "This & that", "this-i-that"},
	{"test", "This & that", "this-and-that"}, // unknown lang, fallback to "en"
}

func TestSlugMakeLang(t *testing.T) {
	for index, smlt := range SlugMakeLangTests {
		slug := MakeLang(smlt.in, smlt.lang)
		if smlt.out != slug {
			t.Errorf(
				"%d. MakeLang(%q, %q) => out = %q, want %q",
				index, smlt.in, smlt.lang, slug, smlt.out)
		}
	}
}

var SlugMakeUserSubstituteTests = []struct {
	cSub map[string]string
	lang string
	in   string
	out  string
}{
	{map[string]string{"'": " "}, "en", "That's great", "that-s-great"},
	{map[string]string{"&": "or"}, "en", "This & that", "this-or-that"}, // by default "&" => "and"
	{map[string]string{"&": "or"}, "de", "This & that", "this-or-that"}, // by default "&" => "und"
}

func TestSlugMakeUserSubstituteLang(t *testing.T) {
	for index, smust := range SlugMakeUserSubstituteTests {
		CustomSub = smust.cSub
		slug := MakeLang(smust.in, smust.lang)
		if smust.out != slug {
			t.Errorf(
				"%d. %q; MakeLang(%q, %q) => out = %q, want %q",
				index, smust.cSub, smust.in, smust.lang,
				slug, smust.out)

		}
	}
}

// Always substitute runes first
var SlugMakeSubstituteOrderTests = []struct {
	rSub map[rune]string
	sSub map[string]string
	in   string
	out  string
}{
	{map[rune]string{'o': "left"}, map[string]string{"o": "right"}, "o o", "left-left"},
	{map[rune]string{'&': "down"}, map[string]string{"&": "up"}, "&", "down"},
}

func TestSlugMakeSubstituteOrderLang(t *testing.T) {
	for index, smsot := range SlugMakeSubstituteOrderTests {
		CustomRuneSub = smsot.rSub
		CustomSub = smsot.sSub
		slug := Make(smsot.in)
		if smsot.out != slug {
			t.Errorf(
				"%d. %q; %q; Make(%q) => out = %q, want %q",
				index, smsot.rSub, smsot.sSub, smsot.in,
				slug, smsot.out)

		}
	}
}

var SlugSubstituteTests = []struct {
	cSub map[string]string
	in   string
	out  string
}{
	{map[string]string{"o": "no"}, "o o o", "no no no"},
	{map[string]string{"'": " "}, "That's great", "That s great"},
}

func TestSubstituteLang(t *testing.T) {
	for index, sst := range SlugSubstituteTests {
		text := Substitute(sst.in, sst.cSub)
		if sst.out != text {
			t.Errorf(
				"%d. Substitute(%q, %q) => out = %q, want %q",
				index, sst.in, sst.cSub, text, sst.out)
		}
	}
}

var SlugSubstituteRuneTests = []struct {
	cSub map[rune]string
	in   string
	out  string
}{
	{map[rune]string{'o': "no"}, "o o o", "no no no"},
	{map[rune]string{'\'': " "}, "That's great", "That s great"},
}

func TestSubstituteRuneLang(t *testing.T) {
	for index, ssrt := range SlugSubstituteRuneTests {
		text := SubstituteRune(ssrt.in, ssrt.cSub)
		if ssrt.out != text {
			t.Errorf(
				"%d. SubstituteRune(%q, %q) => out = %q, want %q",
				index, ssrt.in, ssrt.cSub, text, ssrt.out)
		}
	}
}

var SlugMakeSmartTruncateTests = []struct {
	in        string
	maxLength int
	out       string
}{
	{"DOBROSLAWZYBORT", 100, "dobroslawzybort"},
	{"Dobroslaw Zybort", 100, "dobroslaw-zybort"},
	{"Dobroslaw Zybort", 12, "dobroslaw"},
	{"  Dobroslaw     Zybort  ?", 12, "dobroslaw"},
	{"Ala ma 6 kotów.", 10, "ala-ma-6"},
	{"Dobrosław Żybort", 5, "dobro"},
}

func TestSlugMakeSmartTruncate(t *testing.T) {
	for index, smstt := range SlugMakeSmartTruncateTests {
		MaxLength = smstt.maxLength
		slug := Make(smstt.in)
		if smstt.out != slug {
			t.Errorf(
				"%d. MaxLength = %v; Make(%q) => out = %q, want %q",
				index, smstt.maxLength, smstt.in, slug, smstt.out)
		}
	}
}
