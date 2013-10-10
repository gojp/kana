package kana

import (
	"strings"
	"unicode"
)

type Kana struct {
	kanaToRomajiTrie     *Trie
	romajiToHiraganaTrie *Trie
	romajiToKatakanaTrie *Trie
}

var consonants []string = []string{"b", "d", "f", "g", "h", "j", "k", "l", "m", "p", "r", "s", "t", "w", "z"}

func NewKana() *Kana {
	/*
	   Build a trie for efficient retrieval of entries
	*/
	kana := &Kana{newTrie(), newTrie(), newTrie()}
	kana.initialize()
	return kana
}

func (k *Kana) initialize() {
	/*
		Build the Hiragana + Katakana trie.

		Because there is no overlap between the hiragana and katakana sets,
		they both use the same trie without conflict. Nice bonus!
	*/
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
						k.kanaToRomajiTrie.insert(singleKana, value)
						if t == 0 {
							k.romajiToHiraganaTrie.insert(value, singleKana)
						} else if t == 1 {
							k.romajiToKatakanaTrie.insert(value, singleKana)
						}
					}
				}
			}
		}
	}
}

func (k Kana) KanaToRomaji(kana string) (romaji string) {
	romaji = k.kanaToRomajiTrie.convert(kana)

	// do some post-processing for the tsu and stripe characters
	// maybe a bit of a hacky solution - how can we improve?
	// (they act more like punctuation)
	tsus := []string{"っ", "ッ"}
	for _, tsu := range tsus {
		for i := strings.Index(romaji, tsu); i > -1; i = strings.Index(romaji, tsu) {
			rune_romaji := []rune(romaji)
			if len(rune_romaji) > i+2 {
				// TODO: should check if following letter is consonant
				followingLetter := string(rune_romaji[i+1 : i+2])
				romaji = strings.Replace(romaji, tsu, followingLetter, 1)
			} else {
				romaji = strings.Replace(romaji, tsu, "", 1)
			}
		}
	}

	line := "ー"
	for i := strings.Index(romaji, line); i > -1; i = strings.Index(romaji, line) {
		rune_romaji := []rune(romaji)
		if i > 0 {
			// TODO: should check if following letter is consonant
			previousLetter := string(rune_romaji[i-1 : i])
			romaji = strings.Replace(romaji, line, previousLetter, 1)
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

func (k Kana) RomajiToHiragana(romaji string) (hiragana string) {
	romaji = strings.Replace(romaji, "-", "ー", -1)
	romaji = replace_tsus(romaji, "っ")
	hiragana = k.romajiToHiraganaTrie.convert(romaji)
	return hiragana
}

func (k Kana) RomajiToKatakana(romaji string) (katakana string) {
	romaji = strings.Replace(romaji, "-", "ー", -1)
	// convert double consonants to little tsus first
	romaji = replace_tsus(romaji, "ッ")
	katakana = k.romajiToKatakanaTrie.convert(romaji)
	return katakana
}

func (k Kana) IsLatin(s string) bool {
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

func (k Kana) IsKana(s string) bool {
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

func (k Kana) IsKanji(s string) bool {
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
