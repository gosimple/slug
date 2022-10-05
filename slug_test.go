// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package slug

import (
	"testing"
)

//=============================================================================

func TestSlugMake(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{"DOBROSLAWZYBORT", "dobroslawzybort"},
		{"Dobroslaw Zybort", "dobroslaw-zybort"},
		{"  Dobroslaw     Zybort  ?", "dobroslaw-zybort"},
		{"Dobrosław Żybort", "dobroslaw-zybort"},
		{"Ala ma 6 kotów.", "ala-ma-6-kotow"},

		{"áÁàÀãÃâÂäÄąĄą̊Ą̊", "aaaaaaaaaaaaaa"},
		{"ćĆĉĈçÇčČ", "cccccccc"},
		{"éÉèÈẽẼêÊëËęĘěĚ", "eeeeeeeeeeeeee"},
		{"íÍìÌĩĨîÎïÏįĮ", "iiiiiiiiiiii"},
		{"łŁ", "ll"},
		{"ńŃ", "nn"},
		{"óÓòÒõÕôÔöÖǫǪǭǬø", "ooooooooooooooo"},
		{"śŚšŠ", "ssss"},
		{"řŘ", "rr"},
		{"ťŤ", "tt"},
		{"úÚùÙũŨûÛüÜųŲůŮ", "uuuuuuuuuuuuuu"},
		{"y̨Y̨ýÝ", "yyyy"},
		{"źŹżŹžŽ", "zzzzzz"},
		{"·/,:;`˜'\"", ""},
		{"2000–2013", "2000-2013"},
		{"style—not", "style-not"},
		{"test_slug", "test_slug"},
		{"_test_slug_", "test_slug"},
		{"-test-slug-", "test-slug"},
		{"Æ", "ae"},
		{"Ich heiße", "ich-heisse"},
		{"𐀀", ""}, // Bug #53
		{"% 5 @ 4 $ 3 / 2 & 1 & 2 # 3 @ 4 _ 5", "5-at-4-3-2-and-1-and-2-3-at-4-_-5"},

		{"This & that", "this-and-that"},
		{"fácil €", "facil-eu"},
		{"smile ☺", "smile"},
		{"Hellö Wörld хелло ворлд", "hello-world-khello-vorld"},
		{"\"C'est déjà l’été.\"", "cest-deja-lete"},
		{"jaja---lol-méméméoo--a", "jaja-lol-mememeoo-a"},
		{"影師", "ying-shi"},
		{"Đanković & Kožušček", "dankovic-and-kozuscek"},
		{"ĂăÂâÎîȘșȚț", "aaaaiisstt"},
	}

	for index, st := range testCases {
		got := Make(st.in)
		if got != st.want {
			t.Errorf(
				"%d. Make(%#v) = %#v; want %#v",
				index, st.in, got, st.want)
		}
	}
}

func TestSlugMakeLang(t *testing.T) {
	testCases := []struct {
		lang      string
		in        string
		want      string
		lowercase bool
	}{
		{"bg", "АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЬЮЯабвгдежзийклмнопрстуфхцчшщъьюя", "abvgdezhziyklmnoprstufhtschshshayyuyaabvgdezhziyklmnoprstufhtschshshtayyuya", true},
		{"bg", "АБВГДЕЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЬЮЯабвгдежзийклмнопрстуфхцчшщъьюя", "ABVGDEZhZIYKLMNOPRSTUFHTsChShShAYYuYaabvgdezhziyklmnoprstufhtschshshtayyuya", false},
		{"cs", "ěščřžýáíéúůóňťĚŠČŘŽÝÁÍÉÚŮÓŇŤ", "escrzyaieuuontescrzyaieuuont", true},
		{"cs", "ěščřžýáíéúůóňťĚŠČŘŽÝÁÍÉÚŮÓŇŤ", "escrzyaieuuontESCRZYAIEUUONT", false},
		{"ces", "ěščřžýáíéúůóňťĚŠČŘŽÝÁÍÉÚŮÓŇŤ", "escrzyaieuuontescrzyaieuuont", true},
		{"ces", "ěščřžýáíéúůóňťĚŠČŘŽÝÁÍÉÚŮÓŇŤ", "escrzyaieuuontESCRZYAIEUUONT", false},
		{"de", "Wir mögen Bücher & Käse", "wir-moegen-buecher-und-kaese", true},
		{"de", "Wir mögen Bücher & Käse", "Wir-moegen-Buecher-und-Kaese", false},
		{"de", "Äpfel Über Österreich", "aepfel-ueber-oesterreich", true},
		{"de", "Äpfel Über Österreich", "Aepfel-Ueber-Oesterreich", false},
		{"en", "äÄäöÖöüÜü", "aaaooouuu", true},
		{"en", "äÄäöÖöüÜü", "aAaoOouUu", false},
		{"gr", "ϊχώΩϋ", "ichooy", true},
		{"gr", "ϊχώΩϋ", "ichoOy", false},
		{"Ell", "ϊχώΩϋ", "ichooy", true},  // Greek
		{"Ell", "ϊχώΩϋ", "ichoOy", false}, // Greek
		{"hu", "Árvíztűrő tükörfúrógép", "arvizturo-tukorfurogep", true},
		{"hu", "Árvíztűrő tükörfúrógép", "Arvizturo-tukorfurogep", false},
		{"hu", "SzÉlÜtÖtt ŰrÚjsÁgírÓnŐ", "SzElUtOtt-UrUjsAgirOnO", false},
		{"kk", "әғһіңөқұүӘҒҺІҢӨҚҰҮ", "aghinoquuaghinoquu", true},
		{"kk", "әғһіңөқұүӘҒҺІҢӨҚҰҮ", "aghinoquuAGHINOQUU", false},
		{"ro", "ĂăÂăÎîȘșȚț", "aaaaiisstt", true},
		{"ro", "ĂăÂăÎîȘșȚț", "AaAaIiSsTt", false},
		{"tr", "şüöğıçŞÜÖİĞÇ", "suogicsuoigc", true},
		{"tr", "şüöğıçŞÜÖİĞÇ", "suogicSUOIGC", false},

		// & fun.
		{"bg", "Това и онова", "tova-i-onova", true},
		{"cs", "Toto & Tamto", "toto-a-tamto", true},
		{"cs", "Toto & Tamto", "Toto-a-Tamto", false},
		{"cs", "Toto @ Tamto", "toto-zavinac-tamto", true},
		{"cs", "Toto @ Tamto", "Toto-zavinac-Tamto", false},
		{"ces", "Toto & Tamto", "toto-a-tamto", true},
		{"ces", "Toto & Tamto", "Toto-a-Tamto", false},
		{"ces", "Toto @ Tamto", "toto-zavinac-tamto", true},
		{"ces", "Toto @ Tamto", "Toto-zavinac-Tamto", false},
		{"de", "This & that", "this-und-that", true},
		{"en", "This & that", "this-and-that", true},
		{"es", "This & that", "this-y-that", true},
		{"fi", "This & that", "this-ja-that", true},
		{"fr", "This & that", "this-et-that", true},
		{"fr", "This @ that", "this-arobase-that", true},
		{"gr", "This & that", "this-kai-that", true},
		{"ell", "This & that", "this-kai-that", true}, // Greek
		{"Ell", "This & that", "this-kai-that", true}, // Greek
		{"id", "This & that", "this-dan-that", true},
		{"it", "This & that", "this-e-that", true},
		{"it", "This @ that", "this-chiocciola-that", true},
		{"kk", "This & that", "this-jane-that", true},
		{"kk", "This @ that", "this-that", true},
		{"nl", "This & that", "this-en-that", true},
		{"pl", "This & that", "this-i-that", true},
		{"pol", "This & that", "this-i-that", true},
		{"sv", "This & that", "this-och-that", true},
		{"sv", "This @ that", "this-snabel-a-that", true},
		{"swe", "This & that", "this-och-that", true},
		{"swe", "This @ that", "this-snabel-a-that", true},
		{"nb", "Ærlig, Østen, Åse", "aerlig-oesten-aase", true},
		{"nb", "This & that", "this-og-that", true},
		{"nb", "This @ that", "this-at-that", true},
		{"nn", "Ærlig, Østen, Åse", "aerlig-oesten-aase", true},
		{"nn", "This & that", "this-og-that", true},
		{"nn", "This @ that", "this-at-that", true},
		{"tr", "This & that", "this-ve-that", true},
		{"sl", "đanković & Kožušček", "dzankovic-in-kozuscek", true},
		{"sl", "ĐankoVIĆ & KOŽUŠČEK", "DZankoVIC-in-KOZUSCEK", false},
		{"test", "This & that", "this-and-that", true}, // unknown lang, fallback to "en"

		// Test defaultSub, when adding new lang copy/paste this line,
		// it contain special characters.
		{"bg", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"cs", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"de", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"en", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"es", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"fi", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"gr", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"kk", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"nb", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"nn", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"nl", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"pl", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"ro", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"sl", "1\"2'3’4-5–6—7―8", "1234-5-6-7-8", true},
		{"sv", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
		{"tr", "1\"2'3’4‒5–6—7―8", "1234-5-6-7-8", true},
	}

	for index, smlt := range testCases {
		Lowercase = smlt.lowercase
		got := MakeLang(smlt.in, smlt.lang)
		if got != smlt.want {
			t.Errorf(
				"%d. MakeLang(%#v, %#v) = %#v; want %#v",
				index, smlt.in, smlt.lang, got, smlt.want)
		}
	}
}

func TestSlugMakeUserSubstituteLang(t *testing.T) {
	testCases := []struct {
		cSub map[string]string
		lang string
		in   string
		want string
	}{
		{map[string]string{"'": " "}, "en", "That's great", "that-s-great"},
		{map[string]string{"&": "or"}, "en", "This & that", "this-or-that"},                   // by default "&" => "and"
		{map[string]string{"&": "or"}, "cs", "This & that", "this-or-that"},                   // by default "&" => "a"
		{map[string]string{"&": "or"}, "ces", "This & that", "this-or-that"},                  // by default "&" => "a"
		{map[string]string{"&": "or"}, "de", "This & that", "this-or-that"},                   // by default "&" => "und"
		{map[string]string{"&": "or"}, "DEU", "This & that", "this-or-that"},                  // by default "&" => "und"
		{map[string]string{"&": "or"}, "Fin", "This & that", "this-or-that"},                  // by default "&" => "ja"
		{map[string]string{"&": "or"}, "fr", "This & that", "this-or-that"},                   // by default "&" => "ja"
		{map[string]string{"&": "or"}, "kk", "This & that", "this-or-that"},                   // by default "&" => "jane"
		{map[string]string{"&": "or", "@": "the"}, "de", "@ This & that", "the-this-or-that"}, // by default "&" => "und", "@" => "an"
		{map[string]string{"&": "or", "@": "the"}, "sv", "@ This & that", "the-this-or-that"}, // by default "&" => "och", "@" => "snabel a"
	}

	for index, smust := range testCases {
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

func TestSlugMakeSubstituteOrderLang(t *testing.T) {
	// Always substitute runes first
	testCases := []struct {
		rSub map[rune]string
		sSub map[string]string
		in   string
		want string
	}{
		{map[rune]string{'o': "left"}, map[string]string{"o": "right"}, "o o", "left-left"},
		{map[rune]string{'o': "left", 'a': "r"}, map[string]string{"o": "right"}, "o a o", "left-r-left"},
		{map[rune]string{'o': "left"}, map[string]string{"o": "right", "a": "r"}, "a o a o", "r-left-r-left"},
		{map[rune]string{'&': "down"}, map[string]string{"&": "up"}, "&", "down"},
	}

	for index, smsot := range testCases {
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

func TestSubstituteLang(t *testing.T) {
	testCases := []struct {
		cSub map[string]string
		in   string
		want string
	}{
		{map[string]string{"o": "no"}, "o o o", "no no no"},
		{map[string]string{"o": "no", "a": "or"}, "o a o", "no nor no"},
		{map[string]string{"a": "or", "o": "no"}, "o a o", "no nor no"},
		{map[string]string{"'": " "}, "That's great", "That s great"},
		{map[string]string{"♥": "love"}, "Some ♥", "Some love"},
	}

	for index, sst := range testCases {
		got := Substitute(sst.in, sst.cSub)
		if got != sst.want {
			t.Errorf(
				"%d. Substitute(%#v, %#v) = %#v; want %#v",
				index, sst.in, sst.cSub, got, sst.want)
		}
	}
}

func TestSubstituteRuneLang(t *testing.T) {
	testCases := []struct {
		cSub map[rune]string
		in   string
		want string
	}{
		{map[rune]string{'o': "no"}, "o o o", "no no no"},
		{map[rune]string{'o': "no", 'a': "or"}, "o a o", "no or no"},
		{map[rune]string{'a': "or", 'o': "no"}, "o a o", "no or no"},
		{map[rune]string{'\'': " "}, "That's great", "That s great"},
	}

	for index, ssrt := range testCases {
		got := SubstituteRune(ssrt.in, ssrt.cSub)
		if got != ssrt.want {
			t.Errorf(
				"%d. SubstituteRune(%#v, %#v) = %#v; want %#v",
				index, ssrt.in, ssrt.cSub, got, ssrt.want)
		}
	}
}

func TestSlugMakeSmartTruncate(t *testing.T) {
	testCases := []struct {
		in            string
		maxLength     int
		want          string
		smartTruncate bool
	}{
		{"Dobrosław Żybort", 5, "dobro", true},
		{"Dobroslaw Zybort", 9, "dobroslaw", true},
		{"Dobroslaw Zybort", 12, "dobroslaw", true},
		{"Dobroslaw Zybort", 15, "dobroslaw", true},
		{"Dobroslaw Zybort", 15, "dobroslaw-zybor", false},
		{"Dobroslaw Zybort", 16, "dobroslaw-zybort", true},
		{"Dobroslaw Zybort", 17, "dobroslaw-zybort", true},
		{"Dobroslaw Zybort", 100, "dobroslaw-zybort", true},
		{"abc", 2, "ab", true},
		{"-abc", 2, "ab", true},
		{"abc-", 2, "ab", true},
		{"abc", 3, "abc", true},
		{"-abc", 3, "abc", true},
		{"abc", 4, "abc", true},
		{"abc-", 4, "abc", true},
		{"-abc-", 4, "abc", true},
		{"----abc----", 4, "abc", true},
		{"abc-de", 4, "abc", true},
		{"abc-de", 5, "abc", true},
		{"abc-de", 6, "abc-de", true},
		{"abc-de", 7, "abc-de", true},
		{"abc-de-fg", 6, "abc-de", true},
		{"abc-de-fg", 7, "abc-de", true},
		{"abc-de-fg", 8, "abc-de", true},
		{"abc-de-fg", 9, "abc-de-fg", true},
		{"abc-de-fg", 10, "abc-de-fg", true},

		{"DOBROSLAWZYBORT", 9, "dobroslaw", true},
		{"DOBROSLAWZYBORT", 15, "dobroslawzybort", true},
		{"DOBROSLAWZYBORT", 17, "dobroslawzybort", false},
		{"DOBROSLAWZYBORT", 100, "dobroslawzybort", true},
		{"  Dobroslaw     Zybort  ?", 12, "dobroslaw", true},
		{"Ala ma 6 kotów.", 10, "ala-ma-6", true},

		// No smart truncate
		{"Long branch-name", 14, "long-branch-na", false},
		{"Long branch-name", 12, "long-branch", false},
	}

	for index, smstt := range testCases {
		MaxLength = smstt.maxLength
		if smstt.smartTruncate {
			EnableSmartTruncate = true
		} else {
			EnableSmartTruncate = false
		}

		got := Make(smstt.in)
		if got != smstt.want {
			t.Errorf(
				"%d. MaxLength = %v; Make(%#v) = %#v; want %#v",
				index, smstt.maxLength, smstt.in, got, smstt.want)
		}
	}
}

func TestIsSlug(t *testing.T) {
	MaxLength = 0
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"some", args{"some"}, true},
		{"with -", args{"some-more"}, true},
		{"with _", args{"some_more"}, true},
		{"with numbers", args{"number-2"}, true},
		{"empty string", args{""}, false},
		{"upper case", args{"Some-more"}, false},
		{"space", args{"some more"}, false},
		{"starts with '-'", args{"-some"}, false},
		{"ends with '-'", args{"some-"}, false},
		{"starts with '_'", args{"_some"}, false},
		{"ends with '_'", args{"some_"}, false},
		{"outside ASCII", args{"Dobrosław Żybort"}, false},
		{"outside ASCII –", args{"2000–2013"}, false},
		{"smile ☺", args{"smile ☺"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSlug(tt.args.text); got != tt.want {
				t.Errorf("IsSlug() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("MaxLength", func(t *testing.T) {
		MaxLength = 4
		if got := IsSlug("012345"); got {
			t.Errorf("IsSlug() = %v, want %v", got, false)
		}
		MaxLength = 0
	})
}

func BenchmarkMakeShortAscii(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Make("Hello world")
	}
}

func BenchmarkMakeShort(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Make("хелло ворлд")
	}
}

func BenchmarkMakeShortSymbols(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Make("·/,:;`˜'\" &€￡￥")
	}
}

func BenchmarkMakeMediumAscii(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Make("ABCDE FGHIJ KLMNO PQRST UWXYZ ABCDE FGHIJ KLMNO PQRST UWXYZ ABCDE")
	}
}

func BenchmarkMakeMedium(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		Make("ｦｧｨｩｪ ｫｬｭｮｯ ｰｱｲｳｴ ｵｶｷｸｹ ｺｻｼｽｾ ｿﾀﾁﾂﾃ ﾄﾅﾆﾇﾈ ﾉﾊﾋﾌﾍ ﾎﾏﾐﾑﾒ ﾓﾔﾕﾖﾗ ﾘﾙﾚﾛﾜ")
	}
}

func BenchmarkMakeLongAscii(b *testing.B) {
	longStr := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi " +
		"pulvinar sodales ultrices. Nulla facilisi. Sed at vestibulum erat. Ut " +
		"sit amet urna posuere, sagittis eros ac, varius nisi. Morbi ullamcorper " +
		"odio at nunc pulvinar mattis. Vestibulum rutrum, ante eu dictum mattis, " +
		"elit risus finibus nunc, consectetur facilisis eros leo ut sapien. Sed " +
		"pulvinar volutpat mi. Cras semper mi ac eros accumsan, at feugiat massa " +
		"elementum. Morbi eget dolor sit amet purus condimentum egestas non ut " +
		"sapien. Duis feugiat magna vitae nisi lobortis, quis finibus sem " +
		"sollicitudin. Pellentesque eleifend blandit ipsum, ut porta arcu " +
		"ultricies et. Fusce vel ipsum porta, placerat diam ac, consectetur " +
		"magna. Nulla in porta sem. Suspendisse commodo, felis in molestie " +
		"ultricies, arcu ipsum aliquet turpis, elementum dapibus ipsum lorem a " +
		"nisl. Etiam varius imperdiet placerat. Aliquam euismod lacus arcu, " +
		"ultrices hendrerit est pellentesque vel. Aliquam sit amet laoreet leo. " +
		"Integer eros libero, mollis sed posuere."

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Make(longStr)
	}
}

func BenchmarkSubstituteRuneShort(b *testing.B) {
	shortStr := "Hello/Hi world"
	subs := map[rune]string{'o': "no", '/': "slash"}

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SubstituteRune(shortStr, subs)
	}
}

func BenchmarkSubstituteRuneLong(b *testing.B) {
	longStr := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi " +
		"pulvinar sodales ultrices. Nulla facilisi. Sed at vestibulum erat. Ut " +
		"sit amet urna posuere, sagittis eros ac, varius nisi. Morbi ullamcorper " +
		"odio at nunc pulvinar mattis. Vestibulum rutrum, ante eu dictum mattis, " +
		"elit risus finibus nunc, consectetur facilisis eros leo ut sapien. Sed " +
		"pulvinar volutpat mi. Cras semper mi ac eros accumsan, at feugiat massa " +
		"elementum. Morbi eget dolor sit amet purus condimentum egestas non ut " +
		"sapien. Duis feugiat magna vitae nisi lobortis, quis finibus sem " +
		"sollicitudin. Pellentesque eleifend blandit ipsum, ut porta arcu " +
		"ultricies et. Fusce vel ipsum porta, placerat diam ac, consectetur " +
		"magna. Nulla in porta sem. Suspendisse commodo, felis in molestie " +
		"ultricies, arcu ipsum aliquet turpis, elementum dapibus ipsum lorem a " +
		"nisl. Etiam varius imperdiet placerat. Aliquam euismod lacus arcu, " +
		"ultrices hendrerit est pellentesque vel. Aliquam sit amet laoreet leo. " +
		"Integer eros libero, mollis sed posuere."
	subs := map[rune]string{
		'o': "no",
		'/': "slash",
		'i': "done",
		'E': "es",
		'a': "ASD",
		'1': "one",
		'l': "onetwo",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SubstituteRune(longStr, subs)
	}
}

func BenchmarkSmartTruncateShort(b *testing.B) {
	shortStr := "Hello-world"
	MaxLength = 8

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		smartTruncate(shortStr)
	}
}

func BenchmarkSmartTruncateLong(b *testing.B) {
	longStr := "Lorem-ipsum-dolor-sit-amet,-consectetur-adipiscing-elit.-Morbi-" +
		"pulvinar-sodales-ultrices.-Nulla-facilisi.-Sed-at-vestibulum-erat.-Ut-" +
		"sit-amet-urna-posuere,-sagittis-eros-ac,-varius-nisi.-Morbi-ullamcorper-" +
		"odio-at-nunc-pulvinar-mattis.-Vestibulum-rutrum,-ante-eu-dictum-mattis,-" +
		"elit-risus-finibus-nunc,-consectetur-facilisis-eros-leo-ut-sapien.-Sed-" +
		"pulvinar-volutpat-mi.-Cras-semper-mi-ac-eros-accumsan,-at-feugiat-massa-" +
		"elementum.-Morbi-eget-dolor-sit-amet-purus-condimentum-egestas-non-ut-" +
		"sapien.-Duis-feugiat-magna-vitae-nisi-lobortis,-quis-finibus-sem-" +
		"sollicitudin.-Pellentesque-eleifend-blandit-ipsum,-ut-porta-arcu-" +
		"ultricies-et.-Fusce-vel-ipsum-porta,-placerat-diam-ac,-consectetur-" +
		"magna.-Nulla-in-porta-sem.-Suspendisse-commodo,-felis-in-molestie-" +
		"ultricies,-arcu-ipsum-aliquet-turpis,-elementum-dapibus-ipsum-lorem-a-" +
		"nisl.-Etiam-varius-imperdiet-placerat.-Aliquam-euismod-lacus-arcu,-" +
		"ultrices-hendrerit-est-pellentesque-vel.-Aliquam-sit-amet-laoreet-leo.-" +
		"Integer-eros-libero,-mollis-sed-posuere."
	MaxLength = 256

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		smartTruncate(longStr)
	}
}

func BenchmarkIsSlugShort(b *testing.B) {
	shortStr := "hello-world"

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		IsSlug(shortStr)
	}
}

func BenchmarkIsSlugLong(b *testing.B) {
	longStr := "lorem-ipsum-dolor-sit-amet-consectetur-adipiscing-elit-morbi-" +
		"pulvinar-sodales-ultrices-nulla-facilisi-sed-at-vestibulum-erat-ut-" +
		"sit-amet-urna-posuere-sagittis-eros-ac-varius-nisi-morbi-ullamcorper-" +
		"odio-at-nunc-pulvinar-mattis-vestibulum-rutrum-ante-eu-dictum-mattis,-" +
		"elit-risus-finibus-nunc-consectetur-facilisis-eros-leo-ut-sapien-sed-" +
		"pulvinar-volutpat-mi-cras-semper-mi-ac-eros-accumsan-at-feugiat-massa-" +
		"elementum-morbi-eget-dolor-sit-amet-purus-condimentum-egestas-non-ut-" +
		"sapien-duis-feugiat-magna-vitae-nisi-lobortis-quis-finibus-sem-" +
		"sollicitudin-pellentesque-eleifend-blandit-ipsum-ut-porta-arcu-" +
		"ultricies-et-fusce-vel-ipsum-porta-placerat-diam-ac-consectetur-" +
		"magna-nulla-in-porta-sem-suspendisse-commodo-felis-in-molestie-" +
		"ultricies-arcu-ipsum-aliquet-turpis-elementum-dapibus-ipsum-lorem-a-" +
		"nisl-etiam-varius-imperdiet-placerat-aliquam-euismod-lacus-arcu-" +
		"ultrices-hendrerit-est-pellentesque-vel-aliquam-sit-amet-laoreet-leo-" +
		"integer-eros-libero-mollis-sed-posuere"

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		IsSlug(longStr)
	}
}
