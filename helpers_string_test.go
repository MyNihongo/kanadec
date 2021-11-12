package kanadec

import (
	"math/rand"
	"strings"
	"time"

	res "github.com/MyNihongo/kanadec/internal/resources"
)

func generateHiragana() string {
	return randomWord(res.Hiragana_Start, res.Hiragana_End)
}

func generateKatakana() string {
	return randomWord(res.Katakana_Start, res.Katakana_End)
}

func generateKanji() string {
	return randomWord(res.Kanji_Start, res.Kanji_End, res.KanjiRare_Start, res.KanjiRare_End)
}

func generateRomaji() string {
	return randomWord(res.English_Start, res.English_End)
}

func randomWord(pairs ...rune) string {
	if len(pairs) == 0 || len(pairs)%2 != 0 {
		panic("provide an even number of indices")
	}

	rand.Seed(time.Now().UnixNano())
	pairLen := len(pairs) / 2

	var sb strings.Builder
	strLen := getRandomInt(5, 10)

	for i := 0; i < strLen; i++ {
		pairIndex := i % pairLen * 2
		runeVal := getRandomRune(int(pairs[pairIndex]), int(pairs[pairIndex+1]))
		sb.WriteRune(runeVal)
	}

	return sb.String()
}

func getRandomRune(from int, to int) rune {
	intVal := getRandomInt(from, to)
	return rune(intVal)
}

func getRandomInt(from int, to int) int {
	return rand.Intn(to-from+1) + from
}
