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
		&csSub,
		&deSub,
		&enSub,
		&esSub,
		&fiSub,
		&frSub,
		&grSub,
		&huSub,
		&idSub,
		&kkSub,
		&nbSub,
		&nlSub,
		&nnSub,
		&plSub,
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
