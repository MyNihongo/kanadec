package kanadec

import (
	res "github.com/MyNihongo/kanadec/internal/resources"
)

// IsRuneHiragana checks whether the provided rune is a valid hiragana character
func IsRuneHiragana(value rune) bool {
	return value >= res.Hiragana_Start && value <= res.Hiragana_End
}

// IsRuneKatakana checks whether the provided rune is a valid katakana character
func IsRuneKatakana(value rune) bool {
	return value >= res.Katakana_Start && value <= res.Katakana_End
}

// IsRuneKana checks whether the provided rune is a valid hiragana or katakana character
func IsRuneKana(value rune) bool {
	return IsRuneHiragana(value) || IsRuneKatakana(value)
}

// IsRuneKanji checks whether the provided rune is a valid kanji character
func IsRuneKanji(value rune) bool {
	return value >= res.Kanji_Start && value <= res.Kanji_End ||
		value >= res.KanjiRare_Start && value <= res.KanjiRare_End ||
		value == res.Kanji_IterationMark
}

// IsRuneKanaOrKanji checks whether the provided rune is a valid hiragana, katakana or kanji character
func IsRuneKanaOrKanji(value rune) bool {
	return IsRuneKana(value) || IsRuneKanji(value)
}

// IsRuneRomaji checks whether the provided rune is a valid romaji character
func IsRuneRomaji(value rune) bool {
	if value >= res.English_Start && value <= res.English_End {
		return true
	} else {
		switch value {
		case res.Hepbun_CapitalA, res.Hepbun_SmallA,
			res.Hepbun_CapitalI, res.Hepbun_SmallI,
			res.Hepbun_CapitalU, res.Hepbun_SmallU,
			res.Hepbun_CapitalE, res.Hepbun_SmallE,
			res.Hepbun_CapitalO, res.Hepbun_SmallO:
			return true
		default:
			return false
		}
	}
}
