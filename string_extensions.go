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

// HasHiragana checks whether the provided value has a valid hiragana character
func HasHiragana(value string) bool {
	return hasCheck(value, IsRuneHiragana)
}

// HasKatakana checks whether the provided value has a valid katakana character
func HasKatakana(value string) bool {
	return hasCheck(value, IsRuneKatakana)
}

// HasKana checks whether the provided value has a valid hiragana or katakana character
func HasKana(value string) bool {
	return hasCheck(value, IsRuneKana)
}

// HasKanji checks whether the provided value has a valid kanji character
func HasKanji(value string) bool {
	return hasCheck(value, IsRuneKanji)
}

// HasKanaOrKanji checks whether the provided value has a valid hiragana, katakana or kanji character
func HasKanaOrKanji(value string) bool {
	return hasCheck(value, IsRuneKanaOrKanji)
}

// HasRomaji checks whether the provided value has a valid romaji character
func HasRomaji(value string) bool {
	return hasCheck(value, IsRuneRomaji)
}

func isCheck(value string, check func(rune) bool) bool {
	for _, r := range value {
		if !check(r) {
			return false
		}
	}

	return true
}

func hasCheck(value string, check func(rune) bool) bool {
	for _, r := range value {
		if check(r) {
			return true
		}
	}

	return false
}
