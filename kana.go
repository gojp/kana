package kana

import (
	"io/ioutil"
	"strings"
)

type Kana struct {
	kanaToRomajiTrie     *Trie
	romajiToHiraganaTrie *Trie
	romajiToKatakanaTrie *Trie
}

func newKana() *Kana {
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
	filenames := []string{"hiragana.in", "katakana.in"}
	for _, filename := range filenames {
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			//Do something
		}
		rows := strings.Split(string(content), "\n")
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
						if filename == "hiragana.in" {
							k.romajiToHiraganaTrie.insert(value, singleKana)
						} else if filename == "katakana.in" {
							k.romajiToKatakanaTrie.insert(value, singleKana)
						}
					}
				}
			}
		}
	}
}

func (k Kana) kana_to_romaji(kana string) (romaji string) {
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

func (k Kana) romaji_to_hiragana(romaji string) (hiragana string) {
	romaji = strings.Replace(romaji, "-", "ー", -1)
	hiragana = k.romajiToHiraganaTrie.convert(romaji)
	return hiragana
}

func (k Kana) romaji_to_katakana(romaji string) (katakana string) {
	romaji = strings.Replace(romaji, "-", "ー", -1)
	katakana = k.romajiToKatakanaTrie.convert(romaji)
	return katakana
}
