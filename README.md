Extension methods to detect Japanese characters in `string` and `rune` variables.

```go
import "github.com/MyNihongo/kanadec"
```

### Rune methods
```go
result := kanadec.IsRuneHiragana('ひ')
result = kanadec.IsRuneKatakana('ナ')
result = kanadec.IsRuneKana('ら')

result = kanadec.IsRuneKanji('日')
result = kanadec.IsRuneKanaOrKanji('本')

result = kanadec.IsRuneRomaji('r')
```
### Kana methods
```go
result := kanadec.IsHiragana("ひらがな")
result = kanadec.IsKatakana("カタカナ")
result = kanadec.IsKana("ひらがなカタカナ")

result = kanadec.HasHiragana("歌う")
result = kanadec.HasKatakana("消しゴム")
result = kanadec.HasKana("ケイタイがある")
```
### Kanji methods
```go
result := kanadec.IsKanji("教育制度")
result = kanadec.IsKanaOrKanji("東京に死んでる")

result = kanadec.HasKanji("漢字を覚える")
result = kanadec.HasKanaOrKanji("日本語の勉強")
```
### Romaji methods
```go
result := kanadec.IsRomaji("romaji")
result = kanadec.HasRomaji("romaji")
```