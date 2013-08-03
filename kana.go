package kana

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Kana struct {
	trie *Trie
}

func (k *Kana) initialize() {
	k.trie = newTrie()
	content, err := ioutil.ReadFile("hiragana.in")
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
					// add to trie
					k.trie.insert(singleKana, value)
					fmt.Println(singleKana, value)
				}
			}
		}
	}
}

func (k Kana) kana_to_romaji(kana string) (romaji string) {
	fmt.Println(k.trie)
	// kana_rune := []rune(kana)
	romaji_rune := []rune{}
	for i := 0; i < len(kana_rune); i++ {
		for lookAhead := 2; lookAhead >= 1; lookAhead-- {
			if len(kana_rune) >= i+lookAhead {
				letters := []rune(k.hiraganaTable[string(kana_rune[i:i+lookAhead])])
				if len(letters) > 0 {
					// found in map
					for _, l := range letters {
						romaji_rune = append(romaji_rune, l)
					}
					if lookAhead > 1 {
						i += 1
					}
					break
				} else if lookAhead == 1 {
					// last step and not found
					letters = kana_rune[i : i+1]
					for _, l := range letters {
						romaji_rune = append(romaji_rune, l)
					}
				}
			}
		}
	}
	romaji = string(romaji_rune)
	return romaji
}

func (k Kana) romaji_to_katakana() {

}

func (k Kana) romaji_to_hiragana() {

}
