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
				"%d. Make(%v) => out = %v, want %v",
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
}

func TestSlugMakeLang(t *testing.T) {
	for index, smlt := range SlugMakeLangTests {
		slug := MakeLang(smlt.in, smlt.lang)
		if smlt.out != slug {
			t.Errorf(
				"%d. MakeLang(%v, \"%v\") => out = %v, want %v",
				index, smlt.in, smlt.lang, slug, smlt.out)
		}
	}
}

var SlugSubstituteTests = []struct {
	cSub map[string]string
	in   string
	out  string
}{
	{map[string]string{"water": "sand"}, "water is hot", "sand is hot"},
}

func TestSubstituteLang(t *testing.T) {
	for index, sst := range SlugSubstituteTests {
		text := Substitute(sst.in, sst.cSub)
		if sst.out != text {
			t.Errorf(
				"%d. Substitute(%v, %v) => out = %v, want %v",
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
}

func TestSubstituteRuneLang(t *testing.T) {
	for index, ssrt := range SlugSubstituteRuneTests {
		text := SubstituteRune(ssrt.in, ssrt.cSub)
		if ssrt.out != text {
			t.Errorf(
				"%d. SubstituteRune(%v, %v) => out = %v, want %v",
				index, ssrt.in, ssrt.cSub, text, ssrt.out)
		}
	}
}
