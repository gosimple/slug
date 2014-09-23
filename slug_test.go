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
	in   string
	want string
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
		got := Make(st.in)
		if got != st.want {
			t.Errorf(
				"%d. Make(%#v) = %#v; want %#v",
				index, st.in, got, st.want)
		}
	}
}

var SlugMakeLangTests = []struct {
	lang string
	in   string
	want string
}{
	{"en", "This & that", "this-and-that"},
	{"de", "This & that", "this-und-that"},
	{"pl", "This & that", "this-i-that"},
	{"es", "This & that", "this-y-that"},
	{"test", "This & that", "this-and-that"}, // unknown lang, fallback to "en"
}

func TestSlugMakeLang(t *testing.T) {
	for index, smlt := range SlugMakeLangTests {
		got := MakeLang(smlt.in, smlt.lang)
		if got != smlt.want {
			t.Errorf(
				"%d. MakeLang(%#v, %#v) = %#v; want %#v",
				index, smlt.in, smlt.lang, got, smlt.want)
		}
	}
}

var SlugMakeUserSubstituteTests = []struct {
	cSub map[string]string
	lang string
	in   string
	want string
}{
	{map[string]string{"'": " "}, "en", "That's great", "that-s-great"},
	{map[string]string{"&": "or"}, "en", "This & that", "this-or-that"}, // by default "&" => "and"
	{map[string]string{"&": "or"}, "de", "This & that", "this-or-that"}, // by default "&" => "und"
}

func TestSlugMakeUserSubstituteLang(t *testing.T) {
	for index, smust := range SlugMakeUserSubstituteTests {
		CustomSub = smust.cSub
		got := MakeLang(smust.in, smust.lang)
		if got != smust.want {
			t.Errorf(
				"%d. %#v; MakeLang(%#v, %#v) = %#v; want %#v",
				index, smust.cSub, smust.in, smust.lang,
				got, smust.want)

		}
	}
}

// Always substitute runes first
var SlugMakeSubstituteOrderTests = []struct {
	rSub map[rune]string
	sSub map[string]string
	in   string
	want string
}{
	{map[rune]string{'o': "left"}, map[string]string{"o": "right"}, "o o", "left-left"},
	{map[rune]string{'&': "down"}, map[string]string{"&": "up"}, "&", "down"},
}

func TestSlugMakeSubstituteOrderLang(t *testing.T) {
	for index, smsot := range SlugMakeSubstituteOrderTests {
		CustomRuneSub = smsot.rSub
		CustomSub = smsot.sSub
		got := Make(smsot.in)
		if got != smsot.want {
			t.Errorf(
				"%d. %#v; %#v; Make(%#v) = %#v; want %#v",
				index, smsot.rSub, smsot.sSub, smsot.in,
				got, smsot.want)

		}
	}
}

var SlugSubstituteTests = []struct {
	cSub map[string]string
	in   string
	want string
}{
	{map[string]string{"o": "no"}, "o o o", "no no no"},
	{map[string]string{"'": " "}, "That's great", "That s great"},
}

func TestSubstituteLang(t *testing.T) {
	for index, sst := range SlugSubstituteTests {
		got := Substitute(sst.in, sst.cSub)
		if got != sst.want {
			t.Errorf(
				"%d. Substitute(%#v, %#v) = %#v; want %#v",
				index, sst.in, sst.cSub, got, sst.want)
		}
	}
}

var SlugSubstituteRuneTests = []struct {
	cSub map[rune]string
	in   string
	want string
}{
	{map[rune]string{'o': "no"}, "o o o", "no no no"},
	{map[rune]string{'\'': " "}, "That's great", "That s great"},
}

func TestSubstituteRuneLang(t *testing.T) {
	for index, ssrt := range SlugSubstituteRuneTests {
		got := SubstituteRune(ssrt.in, ssrt.cSub)
		if got != ssrt.want {
			t.Errorf(
				"%d. SubstituteRune(%#v, %#v) = %#v; want %#v",
				index, ssrt.in, ssrt.cSub, got, ssrt.want)
		}
	}
}

var SlugMakeSmartTruncateTests = []struct {
	in        string
	maxLength int
	want      string
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
		got := Make(smstt.in)
		if got != smstt.want {
			t.Errorf(
				"%d. MaxLength = %v; Make(%#v) = %#v; want %#v",
				index, smstt.maxLength, smstt.in, got, smstt.want)
		}
	}
}
