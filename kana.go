package kana

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	consonants = []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "p", "r", "s", "t", "w", "z"}

	hiraganaRe = regexp.MustCompile(`ん([あいうえおなにぬねの])`)
	katakanaRe = regexp.MustCompile(`ン([アイウエオナニヌネノ])`)

	kanaToRomajiTrie     *Trie
	romajiToHiraganaTrie *Trie
	romajiToKatakanaTrie *Trie
)

// Initialize builds the Hiragana + Katakana trie.
// Because there is no overlap between the hiragana and katakana sets,
// they both use the same trie without conflict. Nice bonus!
func Initialize() {
	kanaToRomajiTrie = newTrie()
	romajiToHiraganaTrie = newTrie()
	romajiToKatakanaTrie = newTrie()

	tables := []string{HiraganaTable, KatakanaTable}
	for t, table := range tables {
		rows := strings.Split(table, "\n")
		colNames := strings.Split(string(rows[0]), "\t")[1:]
		for _, row := range rows[1:] {
			cols := strings.Split(string(row), "\t")
			rowName := cols[0]
			for i, kana := range cols[1:] {
				value := rowName + colNames[i]
				kanas := strings.Split(kana, "/")
				for _, singleKana := range kanas {
					if singleKana != "" {
						// add to tries
						kanaToRomajiTrie.insert(singleKana, value)
						if t == 0 {
							romajiToHiraganaTrie.insert(value, singleKana)
						} else if t == 1 {
							romajiToKatakanaTrie.insert(value, singleKana)
						}
					}
				}
			}
		}
	}
}

// KanaToRomaji converts a kana string to its romaji form
func KanaToRomaji(kana string) (romaji string) {
	// unfortunate hack to deal with double n's
	romaji = hiraganaRe.ReplaceAllString(kana, "nn$1")
	romaji = katakanaRe.ReplaceAllString(romaji, "nn$1")

	romaji = kanaToRomajiTrie.convert(romaji)

	// do some post-processing for the tsu and stripe characters
	// maybe a bit of a hacky solution - how can we improve?
	// (they act more like punctuation)
	tsus := []string{"っ", "ッ"}
	for _, tsu := range tsus {
		if strings.Index(romaji, tsu) > -1 {
			for _, c := range romaji {
				ch := string(c)
				if ch == tsu {
					i := strings.Index(romaji, ch)
					runeSize := len(ch)
					followingLetter, _ := utf8.DecodeRuneInString(romaji[i+runeSize:])
					followingLetterStr := string(followingLetter)
					if followingLetterStr != tsu {
						romaji = strings.Replace(romaji, tsu, followingLetterStr, 1)
					} else {
						romaji = strings.Replace(romaji, tsu, "", 1)
					}
				}
			}
		}
	}

	line := "ー"
	for i := strings.Index(romaji, line); i > -1; i = strings.Index(romaji, line) {
		if i > 0 {
			romaji = strings.Replace(romaji, line, "-", 1)
		} else {
			romaji = strings.Replace(romaji, line, "", 1)
		}
	}
	return romaji
}

func replaceTsus(romaji string, tsu string) (result string) {
	result = romaji
	for _, consonant := range consonants {
		result = strings.Replace(result, consonant+consonant, tsu+consonant, -1)
	}
	return result
}

func replaceNs(romaji string, n string) (result string) {
	return strings.Replace(romaji, "nn", n, -1)
}

// RomajiToHiragana converts a romaji string to its hiragana form
func RomajiToHiragana(romaji string) (hiragana string) {
	romaji = strings.Replace(romaji, "-", "ー", -1)
	romaji = replaceTsus(romaji, "っ")
	romaji = replaceNs(romaji, "ん")
	hiragana = romajiToHiraganaTrie.convert(romaji)
	return hiragana
}

// RomajiToKatakana converts a romaji string to its katakana form
func RomajiToKatakana(romaji string) (katakana string) {
	romaji = strings.Replace(romaji, "-", "ー", -1)
	// convert double consonants to little tsus first
	romaji = replaceTsus(romaji, "ッ")
	romaji = replaceNs(romaji, "ン")
	katakana = romajiToKatakanaTrie.convert(romaji)
	return katakana
}

func isChar(s string, rangeTable []*unicode.RangeTable) bool {
	runeForm := []rune(s)
	for _, r := range runeForm {
		if !unicode.IsOneOf(rangeTable, r) {
			return false
		}
	}
	return true
}

// IsLatin returns true if the string contains only Latin characters
func IsLatin(s string) bool {
	return isChar(s, []*unicode.RangeTable{unicode.Latin, unicode.ASCII_Hex_Digit, unicode.White_Space, unicode.Hyphen})
}

// IsKana returns true if the string contains only kana
func IsKana(s string) bool {
	return isChar(s, []*unicode.RangeTable{unicode.Hiragana, unicode.Katakana, unicode.Hyphen, unicode.Diacritic})
}

// IsHiragana returns true if the string contains only hiragana
func IsHiragana(s string) bool {
	return isChar(s, []*unicode.RangeTable{unicode.Hiragana, unicode.Hyphen, unicode.Diacritic})
}

// IsKatakana returns true if the string contains only katakana
func IsKatakana(s string) bool {
	return isChar(s, []*unicode.RangeTable{unicode.Katakana, unicode.Hyphen, unicode.Diacritic})
}

// IsKanji return strue if the string contains only kanji
func IsKanji(s string) bool {
	return isChar(s, []*unicode.RangeTable{unicode.Ideographic})
}

func replaceAll(haystack string, needles []string, replacements []string) (replaced string) {
	replaced = haystack
	for i := range needles {
		replaced = strings.Replace(replaced, needles[i], replacements[i], -1)
	}
	return replaced
}

// NormalizeRomaji transforms romaji input to one specific standard form,
// which should be as close as possible to hiragana so that
// this library gives correct output when transforming to
// hiragana/katakana
func NormalizeRomaji(s string) (romaji string) {
	romaji = s
	romaji = strings.ToLower(romaji)
	romaji = replaceAll(
		romaji,
		[]string{"ā", "ē", "ī", "ō", "ū", "ee", "uu"},
		[]string{"a-", "ei", "ii", "oo", "u-", "ei", "u-"},
	)

	return romaji
}
