package kanadec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// IsHiragana

func TestIsHiragana_Empty(t *testing.T) {
	got := IsHiragana("")
	assert.False(t, got)
}

func TestIsHiragana_Hiragana(t *testing.T) {
	got := IsHiragana(generateHiragana())
	assert.True(t, got)
}

func TestIsHiragana_Katakana(t *testing.T) {
	got := IsHiragana(generateHiragana() + generateKatakana())
	assert.False(t, got)
}

func TestIsHiragana_Kanji(t *testing.T) {
	got := IsHiragana(generateHiragana() + generateKanji())
	assert.False(t, got)
}

func TestIsHiragana_Romaji(t *testing.T) {
	got := IsHiragana(generateHiragana() + generateRomaji())
	assert.False(t, got)
}

// IsKatakana

func TestIsKatakana_Empty(t *testing.T) {
	got := IsKatakana("")
	assert.False(t, got)
}

func TestIsKatakana_Hiragana(t *testing.T) {
	got := IsKatakana(generateKatakana() + generateHiragana())
	assert.False(t, got)
}

func TestIsKatakana_Katakana(t *testing.T) {
	got := IsKatakana(generateKatakana())
	assert.True(t, got)
}

func TestIsKatakana_Kanji(t *testing.T) {
	got := IsKatakana(generateKatakana() + generateKanji())
	assert.False(t, got)
}

func TestIsKatakana_Romaji(t *testing.T) {
	got := IsKatakana(generateKatakana() + generateRomaji())
	assert.False(t, got)
}

// IsKana

func TestIsKana_Empty(t *testing.T) {
	got := IsKana("")
	assert.False(t, got)
}

func TestIsKana_Hiragana(t *testing.T) {
	got := IsKana(generateKatakana() + generateHiragana())
	assert.True(t, got)
}

func TestIsKana_Katakana(t *testing.T) {
	got := IsKana(generateKatakana() + generateHiragana())
	assert.True(t, got)
}

func TestIsKana_Kanji(t *testing.T) {
	got := IsKana(generateKatakana() + generateHiragana() + generateKanji())
	assert.False(t, got)
}

func TestIsKana_Romaji(t *testing.T) {
	got := IsKana(generateKatakana() + generateHiragana() + generateRomaji())
	assert.False(t, got)
}

// IsKanji

func TestIsKanji_Empty(t *testing.T) {
	got := IsKanji("")
	assert.False(t, got)
}

func TestIsKanji_Hiragana(t *testing.T) {
	got := IsKanji(generateKanji() + generateHiragana())
	assert.False(t, got)
}

func TestIsKanji_Katakana(t *testing.T) {
	got := IsKanji(generateKanji() + generateKatakana())
	assert.False(t, got)
}

func TestIsKanji_Kanji(t *testing.T) {
	got := IsKanji(generateKanji())
	assert.True(t, got)
}

func TestIsKanji_Romaji(t *testing.T) {
	got := IsKanji(generateKanji() + generateRomaji())
	assert.False(t, got)
}

// IsKanaOrKanji

func TestIsKanaOrKanji_Empty(t *testing.T) {
	got := IsKanaOrKanji("")
	assert.False(t, got)
}

func TestIsKanaOrKanji_Hiragana(t *testing.T) {
	got := IsKanaOrKanji(generateKanji() + generateHiragana() + generateKatakana())
	assert.True(t, got)
}

func TestIsKanaOrKanji_Katakana(t *testing.T) {
	got := IsKanaOrKanji(generateKanji() + generateHiragana() + generateKatakana())
	assert.True(t, got)
}

func TestIsKanaOrKanji_Kanji(t *testing.T) {
	got := IsKanaOrKanji(generateKanji() + generateHiragana() + generateKatakana())
	assert.True(t, got)
}

func TestIsKanaOrKanji_Romaji(t *testing.T) {
	got := IsKanaOrKanji(generateKanji() + generateHiragana() + generateKatakana() + generateRomaji())
	assert.False(t, got)
}

// IsRomaji

func TestIsRomaji_Empty(t *testing.T) {
	got := IsRomaji("")
	assert.False(t, got)
}

func TestIsRomaji_Hiragana(t *testing.T) {
	got := IsRomaji(generateKanji() + generateHiragana())
	assert.False(t, got)
}

func TestIsRomaji_Katakana(t *testing.T) {
	got := IsRomaji(generateKanji() + generateKatakana())
	assert.False(t, got)
}

func TestIsRomaji_Kanji(t *testing.T) {
	got := IsRomaji(generateKanji())
	assert.False(t, got)
}

func TestIsRomaji_Romaji(t *testing.T) {
	got := IsRomaji(generateRomaji())
	assert.True(t, got)
}

// HasHiragana

func TestHasHiragana_Empty(t *testing.T) {
	got := HasHiragana("")
	assert.False(t, got)
}

func TestHasHiragana_Hiragana(t *testing.T) {
	got := HasHiragana("test" + generateHiragana())
	assert.True(t, got)
}

func TestHasHiragana_Katakana(t *testing.T) {
	got := HasHiragana("test" + generateKatakana())
	assert.False(t, got)
}

func TestHasHiragana_Kanji(t *testing.T) {
	got := HasHiragana("test" + generateKanji())
	assert.False(t, got)
}

func TestHasHiragana_Romaji(t *testing.T) {
	got := HasHiragana("test" + generateRomaji())
	assert.False(t, got)
}

// HasKatakana

func TestHasKatakana_Empty(t *testing.T) {
	got := HasKatakana("")
	assert.False(t, got)
}

func TestHasKatakana_Hiragana(t *testing.T) {
	got := HasKatakana("test" + generateHiragana())
	assert.False(t, got)
}

func TestHasKatakana_Katakana(t *testing.T) {
	got := HasKatakana("test" + generateKatakana())
	assert.True(t, got)
}

func TestHasKatakana_Kanji(t *testing.T) {
	got := HasKatakana("test" + generateKanji())
	assert.False(t, got)
}

func TestHasKatakana_Romaji(t *testing.T) {
	got := HasKatakana("test" + generateRomaji())
	assert.False(t, got)
}

// HasKana

func TestHasKana_Empty(t *testing.T) {
	got := HasKana("")
	assert.False(t, got)
}

func TestHasKana_Hiragana(t *testing.T) {
	got := HasKana("test" + generateHiragana())
	assert.True(t, got)
}

func TestHasKana_Katakana(t *testing.T) {
	got := HasKana("test" + generateKatakana())
	assert.True(t, got)
}

func TestHasKana_Kanji(t *testing.T) {
	got := HasKana("test" + generateKanji())
	assert.False(t, got)
}

func TestHasKana_Romaji(t *testing.T) {
	got := HasKana("test" + generateRomaji())
	assert.False(t, got)
}

// HasKanji

func TestHasKanji_Empty(t *testing.T) {
	got := HasKanji("")
	assert.False(t, got)
}

func TestHasKanji_Hiragana(t *testing.T) {
	got := HasKanji("test" + generateHiragana())
	assert.False(t, got)
}

func TestHasKanji_Katakana(t *testing.T) {
	got := HasKanji("test" + generateKatakana())
	assert.False(t, got)
}

func TestHasKanji_Kanji(t *testing.T) {
	got := HasKanji("test" + generateKanji())
	assert.True(t, got)
}

func TestHasKanji_Romaji(t *testing.T) {
	got := HasKana("test" + generateRomaji())
	assert.False(t, got)
}

// HasKanaOrKanji

func TestHasKanaOrKanji_Empty(t *testing.T) {
	got := HasKanaOrKanji("")
	assert.False(t, got)
}

func TestHasKanaOrKanji_Hiragana(t *testing.T) {
	got := HasKanaOrKanji("test" + generateHiragana())
	assert.True(t, got)
}

func TestHasKanaOrKanji_Katakana(t *testing.T) {
	got := HasKanaOrKanji("test" + generateKatakana())
	assert.True(t, got)
}

func TestHasKanaOrKanji_Kanji(t *testing.T) {
	got := HasKanaOrKanji("test" + generateKanji())
	assert.True(t, got)
}

func TestHasKanaOrKanji_Romaji(t *testing.T) {
	got := HasKanaOrKanji("test" + generateRomaji())
	assert.False(t, got)
}

// HasRomaji

func TestHasRomaji_Empty(t *testing.T) {
	got := HasRomaji("")
	assert.False(t, got)
}

func TestHasRomaji_Hiragana(t *testing.T) {
	got := HasRomaji(generateHiragana())
	assert.False(t, got)
}

func TestHasRomaji_Katakana(t *testing.T) {
	got := HasRomaji(generateKatakana())
	assert.False(t, got)
}

func TestHasRomaji_Kanji(t *testing.T) {
	got := HasRomaji(generateKanji())
	assert.False(t, got)
}

func TestHasRomaji_Romaji(t *testing.T) {
	got := HasRomaji(generateKanji() + generateRomaji())
	assert.True(t, got)
}
