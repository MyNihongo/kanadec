package kanadec

import (
	"testing"

	res "github.com/MyNihongo/kanadec/internal/resources"
	"github.com/stretchr/testify/assert"
)

// IsRuneHiragana

func TestIsRuneHiragana_Hiragana(t *testing.T) {
	for _, v := range hiragana {
		got := IsRuneHiragana(v)
		assert.True(t, got)
	}
}

func TestIsRuneHiragana_Katakana(t *testing.T) {
	for _, v := range katakana {
		got := IsRuneHiragana(v)
		assert.False(t, got)
	}
}

func TestIsRuneHiragana_Kanji(t *testing.T) {
	for _, v := range kanji {
		got := IsRuneHiragana(v)
		assert.False(t, got)
	}
}

func TestIsRuneHiragana_Romaji(t *testing.T) {
	for _, v := range romaji {
		got := IsRuneHiragana(v)
		assert.False(t, got)
	}
}

// TestIsRuneKatakana

func TestIsRuneKatakana_Hiragana(t *testing.T) {
	for _, v := range hiragana {
		got := IsRuneKatakana(v)
		assert.False(t, got)
	}
}

func TestIsRuneKatakana_Katakana(t *testing.T) {
	for _, v := range katakana {
		got := IsRuneKatakana(v)
		assert.True(t, got)
	}
}

func TestIsRuneKatakana_Kanji(t *testing.T) {
	for _, v := range kanji {
		got := IsRuneKatakana(v)
		assert.False(t, got)
	}
}

func TestIsRuneKatakana_Romaji(t *testing.T) {
	for _, v := range romaji {
		got := IsRuneKatakana(v)
		assert.False(t, got)
	}
}

// TestIsRuneKana

func TestIsRuneKana_Hiragana(t *testing.T) {
	for _, v := range hiragana {
		got := IsRuneKana(v)
		assert.True(t, got)
	}
}

func TestIsRuneKana_Katakana(t *testing.T) {
	for _, v := range katakana {
		got := IsRuneKana(v)
		assert.True(t, got)
	}
}

func TestIsRuneKana_Kanji(t *testing.T) {
	for _, v := range kanji {
		got := IsRuneKana(v)
		assert.False(t, got)
	}
}

func TestIsRuneKana_Romaji(t *testing.T) {
	for _, v := range romaji {
		got := IsRuneKana(v)
		assert.False(t, got)
	}
}

// TestIsRuneKanji

func TestIsRuneKanji_Hiragana(t *testing.T) {
	for _, v := range hiragana {
		got := IsRuneKanji(v)
		assert.False(t, got)
	}
}

func TestIsRuneKanji_Katakana(t *testing.T) {
	for _, v := range katakana {
		got := IsRuneKanji(v)
		assert.False(t, got)
	}
}

func TestIsRuneKanji_Kanji(t *testing.T) {
	for _, v := range kanji {
		got := IsRuneKanji(v)
		assert.True(t, got)
	}
}

func TestIsRuneKanji_Romaji(t *testing.T) {
	for _, v := range romaji {
		got := IsRuneKanji(v)
		assert.False(t, got)
	}
}

func TestIsRuneKanji_BreakingPoints(t *testing.T) {
	inputs := []rune{
		res.Kanji_Start - 1,
		res.Kanji_End + 1,
		res.KanjiRare_Start - 1,
		res.KanjiRare_End + 1,
		res.Kanji_IterationMark - 1,
		res.Kanji_IterationMark + 1,
	}

	for _, v := range inputs {
		got := IsRuneKanji(v)
		assert.False(t, got)
	}
}

// TestIsRuneKanaOrKanji

func TestIsRuneKanaOrKanji_Hiragana(t *testing.T) {
	for _, v := range hiragana {
		got := IsRuneKanaOrKanji(v)
		assert.True(t, got)
	}
}

func TestIsRuneKanaOrKanji_Katakana(t *testing.T) {
	for _, v := range katakana {
		got := IsRuneKanaOrKanji(v)
		assert.True(t, got)
	}
}

func TestIsRuneKanaOrKanji_Kanji(t *testing.T) {
	for _, v := range kanji {
		got := IsRuneKanaOrKanji(v)
		assert.True(t, got)
	}
}

func TestIsRuneKanaOrKanji_Romaji(t *testing.T) {
	for _, v := range romaji {
		got := IsRuneKanaOrKanji(v)
		assert.False(t, got)
	}
}

func TestIsRuneKanaOrKanji_BreakingPoints(t *testing.T) {
	inputs := []rune{
		res.Kanji_Start - 1,
		res.Kanji_End + 1,
		res.KanjiRare_Start - 1,
		res.KanjiRare_End + 1,
		res.Kanji_IterationMark - 1,
		res.Kanji_IterationMark + 1,
	}

	for _, v := range inputs {
		got := IsRuneKanaOrKanji(v)
		assert.False(t, got)
	}
}

// TestIsRuneRomaji

func TestIsRuneRomaji_Hiragana(t *testing.T) {
	for _, v := range hiragana {
		got := IsRuneRomaji(v)
		assert.False(t, got)
	}
}

func TestIsRuneRomaji_Katakana(t *testing.T) {
	for _, v := range katakana {
		got := IsRuneRomaji(v)
		assert.False(t, got)
	}
}

func TestIsRuneRomaji_Kanji(t *testing.T) {
	for _, v := range kanji {
		got := IsRuneRomaji(v)
		assert.False(t, got)
	}
}

func TestIsRuneRomaji_Romaji(t *testing.T) {
	for _, v := range romaji {
		got := IsRuneRomaji(v)
		assert.True(t, got)
	}
}
