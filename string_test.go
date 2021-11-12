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
