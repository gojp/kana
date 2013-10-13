package kana

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var consonants []string = []string{"b", "d", "f", "g", "h", "j", "k", "l", "m", "p", "r", "s", "t", "w", "z"}

var kanaToRomajiTrie *Trie
var romajiToHiraganaTrie *Trie
var romajiToKatakanaTrie *Trie

func Initialize() {
	/*
		Build the Hiragana + Katakana trie.

		Because there is no overlap between the hiragana and katakana sets,
		they both use the same trie without conflict. Nice bonus!
	*/
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

func KanaToRomaji(kana string) (romaji string) {
	romaji = kanaToRomajiTrie.convert(kana)

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

func replace_tsus(romaji string, tsu string) (result string) {
	result = romaji
	for _, consonant := range consonants {
		result = strings.Replace(result, consonant+consonant, tsu+consonant, -1)
	}
	return result
}

func RomajiToHiragana(romaji string) (hiragana string) {
	romaji = strings.Replace(romaji, "-", "ー", -1)
	romaji = replace_tsus(romaji, "っ")
	hiragana = romajiToHiraganaTrie.convert(romaji)
	return hiragana
}

func RomajiToKatakana(romaji string) (katakana string) {
	romaji = strings.Replace(romaji, "-", "ー", -1)
	// convert double consonants to little tsus first
	romaji = replace_tsus(romaji, "ッ")
	katakana = romajiToKatakanaTrie.convert(romaji)
	return katakana
}

func IsLatin(s string) bool {
	isLatin := true
	runeForm := []rune(s)
	for _, r := range runeForm {
		isLatin = isLatin && unicode.IsOneOf([]*unicode.RangeTable{unicode.Latin, unicode.ASCII_Hex_Digit, unicode.White_Space, unicode.Hyphen}, r)
		if !isLatin {
			return isLatin
		}
	}
	return isLatin
}

func IsKana(s string) bool {
	isKana := true
	runeForm := []rune(s)
	for _, r := range runeForm {
		isKana = isKana && unicode.IsOneOf([]*unicode.RangeTable{unicode.Hiragana, unicode.Katakana, unicode.Hyphen, unicode.Diacritic}, r)
		if !isKana {
			return isKana
		}
	}
	return isKana
}

func IsKanji(s string) bool {
	isKanji := true
	runeForm := []rune(s)
	for _, r := range runeForm {
		isKanji = isKanji && unicode.IsOneOf([]*unicode.RangeTable{unicode.Ideographic}, r)
		if !isKanji {
			return isKanji
		}
	}
	return isKanji
}

func replaceAll(haystack string, needles []string, replacements []string) (replaced string) {
	replaced = haystack
	for i := range needles {
		replaced = strings.Replace(replaced, needles[i], replacements[i], -1)
	}
	return replaced
}

func NormalizeRomaji(s string) (romaji string) {
	// transform romaji input to one specific standard form,
	// which should be as close as possible to hiragana so that
	// this library gives correct output when transforming to
	// hiragana / katakana

	romaji = s
	romaji = strings.ToLower(romaji)
	romaji = replaceAll(
		romaji,
		[]string{"ā", "ē", "ī", "ō", "ū", "ee", "uu"},
		[]string{"a-", "ei", "ii", "oo", "u-", "ei", "u-"},
	)

	return romaji
}
