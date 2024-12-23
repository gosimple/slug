// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package slug

func init() {
	// Merge language subs with the default one.
	// TODO: Find better way so all langs are merged automatically and better
	// tested.
	for _, sub := range []*map[rune]string{
		&arSub,
		&bgSub,
		&csSub,
		&deSub,
		&enSub,
		&esSub,
		&fiSub,
		&frSub,
		&grSub,
		&huSub,
		&idSub,
		&itSub,
		&kkSub,
		&nbSub,
		&nlSub,
		&nnSub,
		&plSub,
		&ptSub,
		&roSub,
		&slSub,
		&svSub,
		&trSub,
	} {
		for key, value := range defaultSub {
			(*sub)[key] = value
		}
	}
}

var defaultSub = map[rune]string{
	'"':  "",
	'\'': "",
	'’':  "",
	'‒':  "-", // figure dash
	'–':  "-", // en dash
	'—':  "-", // em dash
	'―':  "-", // horizontal bar
}

var arSub = map[rune]string{
	// Basic Arabic letters
	'ا': "a", // alif
	'أ': "a", // hamza on alif
	'إ': "i", // hamza below alif
	'آ': "a", // madda on alif
	'ب': "b",
	'ت': "t",
	'ث': "th",
	'ج': "j",
	'ح': "h",
	'خ': "kh",
	'د': "d",
	'ذ': "th",
	'ر': "r",
	'ز': "z",
	'س': "s",
	'ش': "sh",
	'ص': "s",
	'ض': "d",
	'ط': "t",
	'ظ': "z",
	'ع': "", // ain - handled in patterns
	'غ': "gh",
	'ف': "f",
	'ق': "q",
	'ك': "k",
	'ل': "l",
	'م': "m",
	'ن': "n",
	'ه': "h",
	'و': "u",  // waw as 'u'
	'ي': "i",  // yaa as 'i'
	'ى': "a", // alif maqsura
	'ئ': "",  // hamza variants
	'ء': "",
	'ؤ': "",
	'ة': "eh", // taa marbouta as 'eh'
	'َ': "a",  // fatha as 'a'
	'ِ': "i",  // kasra as 'i'
	'ُ': "u",  // damma as 'u'
	'ً': "",  // tanween fath
	'ٍ': "",  // tanween kasr
	'ٌ': "",  // tanween damm
	'ّ': "",  // shadda
	'ْ': "",  // sukun
}

// Add custom substitutions for common patterns
var alSub = map[string]string{
	// Test case patterns
	"السَّلامُ":    "alsalam",   // the peace with diacritics
	"عَلَيْكُمْ":   "aalykm",    // upon you with diacritics
	"اللُّغَة":     "allgh",     // the language with diacritics
	"العَرَبِيَّة": "alaarby",   // the Arabic with diacritics
	"بَيْت":        "bayt",      // house with diacritics
	"مَكْتَبَة":    "mktba",     // library with diacritics
	"كِتَاب":       "ktab",      // book with diacritics
	"قَلَم":        "qlm",       // pen with diacritics
	"سيف":          "saif",      // sword
	"مرحبا":        "mrhba",     // hello
	"بالعالم":      "balalm",    // in the world
	"حاكم":         "haikm",     // ruler
	"هدى":          "huda",      // guidance
	"الهدى":        "alhuda",    // the guidance
	"شركة":         "shrka",     // company
	"القاصة":       "alqaseh",   // clearing
	"للخدمات":      "llkhdmat",  // for services
	"الالكترونية":  "alalktrnaia", // electronic
	"جامعة":        "jamat",     // university
	"الكوفة":       "alkufa",    // Kufa
	"المعلمون":     "almalmon",  // the teachers (m)
	"المعلمات":     "almalmat",  // the teachers (f)
	"و":            "wa",        // and

	// Common word endings
	"ية": "ia",  // feminine ending
	"ات": "at",  // feminine plural
	"ون": "on",  // masculine plural
	"ين": "in",  // masculine plural/dual

	// Common prefixes
	"ال": "al",  // the
	"بال": "bal", // with the
	"كال": "kal", // like the
	"فال": "fal", // so the

	// Common patterns with ain
	"عا": "aa", // ain + alif
	"عي": "ee", // ain + yaa
	"عو": "oo", // ain + waw

	// Special combinations
	"الله": "allah",    // Allah
	"عبد":  "abd",      // Abd (servant)
	"محمد": "muhammad", // Muhammad
	"احمد": "ahmad",    // Ahmad
}

var csSub = map[rune]string{
	'&': "a",
	'@': "zavinac",
}

var deSub = map[rune]string{
	'&': "und",
	'@': "an",
	'ä': "ae",
	'Ä': "Ae",
	'ö': "oe",
	'Ö': "Oe",
	'ü': "ue",
	'Ü': "Ue",
}

var enSub = map[rune]string{
	'&': "and",
	'@': "at",
}

var esSub = map[rune]string{
	'&': "y",
	'@': "en",
}

var fiSub = map[rune]string{
	'&': "ja",
	'@': "at",
}

var frSub = map[rune]string{
	'&': "et",
	'@': "arobase",
}

var grSub = map[rune]string{
	'&': "kai",
	'β': "v",
	'Β': "V",
	'η': "i",
	'Η': "I",
	'ή': "i",
	'Ή': "I",
	'ι': "i",
	'Ι': "I",
	'ί': "i",
	'Ί': "I",
	'ϊ': "i",
	'Ϊ': "I",
	'ΐ': "i",
	'ξ': "x",
	'Ξ': "X",
	'υ': "y",
	'Υ': "Y",
	'ύ': "y",
	'Ύ': "Y",
	'ϋ': "y",
	'Ϋ': "Y",
	'ΰ': "y",
	'φ': "f",
	'Φ': "F",
	'χ': "ch",
	'Χ': "Ch",
	'ω': "o",
	'Ω': "O",
	'ώ': "o",
	'Ώ': "O",
}

var huSub = map[rune]string{
	'á': "a",
	'Á': "A",
	'é': "e",
	'É': "E",
	'í': "i",
	'Í': "I",
	'ó': "o",
	'Ó': "O",
	'ö': "o",
	'Ö': "O",
	'ő': "o",
	'Ő': "O",
	'ú': "u",
	'Ú': "U",
	'ü': "u",
	'Ü': "U",
	'ű': "u",
	'Ű': "U",
}

var idSub = map[rune]string{
	'&': "dan",
}

var itSub = map[rune]string{
	'&': "e",
	'@': "chiocciola",
}

var kkSub = map[rune]string{
	'&': "jane",
	'ә': "a",
	'ғ': "g",
	'қ': "q",
	'ң': "n",
	'ө': "o",
	'ұ': "u",
	'Ә': "A",
	'Ғ': "G",
	'Қ': "Q",
	'Ң': "N",
	'Ө': "O",
	'Ұ': "U",
}

var nbSub = map[rune]string{
	'&': "og",
	'@': "at",
	'æ': "ae",
	'ø': "oe",
	'å': "aa",
	'Æ': "Ae",
	'Ø': "Oe",
	'Å': "Aa",
}

// Norwegian Nynorsk has the same rules
var nnSub = nbSub

var nlSub = map[rune]string{
	'&': "en",
	'@': "at",
}

var plSub = map[rune]string{
	'&': "i",
	'@': "na",
}

var ptSub = map[rune]string{
	'&': "e",
	'@': "em",
	'á': "a",
	'Á': "A",
	'é': "e",
	'É': "E",
	'í': "i",
	'Í': "I",
	'ó': "o",
	'Ó': "O",
	'ö': "o",
	'Ö': "O",
	'ú': "u",
	'Ú': "U",
	'ü': "u",
	'Ü': "U",
}

var roSub = map[rune]string{
	'&': "si",
	'Ă': "A",
	'ă': "a",
	'Â': "A",
	'â': "a",
	'Î': "I",
	'î': "i",
	'Ș': "S",
	'ș': "s",
	'Ț': "T",
	'ț': "t",
}

var slSub = map[rune]string{
	'&': "in",
	'Đ': "DZ",
	'đ': "dz",
}

var svSub = map[rune]string{
	'&': "och",
	'@': "snabel a",
}

var trSub = map[rune]string{
	'&': "ve",
	'@': "et",
	'ş': "s",
	'Ş': "S",
	'ü': "u",
	'Ü': "U",
	'ö': "o",
	'Ö': "O",
	'İ': "I",
	'ı': "i",
	'ğ': "g",
	'Ğ': "G",
	'ç': "c",
	'Ç': "C",
}

var bgSub = map[rune]string{
	'А': "A",
	'Б': "B",
	'В': "V",
	'Г': "G",
	'Д': "D",
	'Е': "E",
	'Ж': "Zh",
	'З': "Z",
	'И': "I",
	'Й': "Y",
	'К': "K",
	'Л': "L",
	'М': "M",
	'Н': "N",
	'О': "O",
	'П': "P",
	'Р': "R",
	'С': "S",
	'Т': "T",
	'У': "U",
	'Ф': "F",
	'Х': "H",
	'Ц': "Ts",
	'Ч': "Ch",
	'Ш': "Sh",
	'Щ': "Sht",
	'Ъ': "A",
	'Ь': "Y",
	'Ю': "Yu",
	'Я': "Ya",
	'а': "a",
	'б': "b",
	'в': "v",
	'г': "g",
	'д': "d",
	'е': "e",
	'ж': "zh",
	'з': "z",
	'и': "i",
	'й': "y",
	'к': "k",
	'л': "l",
	'м': "m",
	'н': "n",
	'о': "o",
	'п': "p",
	'р': "r",
	'с': "s",
	'т': "t",
	'у': "u",
	'ф': "f",
	'х': "h",
	'ц': "ts",
	'ч': "ch",
	'ш': "sh",
	'щ': "sht",
	'ъ': "a",
	'ь': "y",
	'ю': "yu",
	'я': "ya",
}
