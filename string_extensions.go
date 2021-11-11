package kanadec

// IsHiragana checks whether the provided value is a valid hiragana sequence
func IsHiragana(value string) bool {
	return isCheck(value, IsRuneHiragana)
}

// IsKatakana checks whether the provided value is a valid katakana sequence
func IsKatakana(value string) bool {
	return isCheck(value, IsRuneKatakana)
}

// IsKana checks whether the provided value is a valid hiragana / katakana sequence
func IsKana(value string) bool {
	return isCheck(value, IsRuneKana)
}

// IsKanji checks whether the provided value is a valid kanji sequence
func IsKanji(value string) bool {
	return isCheck(value, IsRuneKanji)
}

// IsKanaOrKanji checks whether the provided value is a valid hiragana / katakana / kanji sequence
func IsKanaOrKanji(value string) bool {
	return isCheck(value, IsRuneKanaOrKanji)
}

// IsRomaji checks whether the provided value is a valid romaji sequence
func IsRomaji(value string) bool {
	return isCheck(value, IsRuneRomaji)
}

func isCheck(value string, check func(rune) bool) bool {
	for _, r := range value {
		if !check(r) {
			return false
		}
	}

	return true
}
