[![Build Status](https://travis-ci.org/gojp/kana.png?branch=master)](https://travis-ci.org/gojp/kana) [![go report card](https://goreportcard.com/badge/github.com/gojp/kana)](http://goreportcard.com/report/github.com/gojp/kana)

# kana

Golang library for convertiong hiragana to romaji, katakana to romaji, romaji to hiragana, and romaji to katakana. 

## Installation

Simply install with `go get`:

    go get github.com/gojp/kana

## Usage

### Convert hiragana or katakana to romaji:

    s := kana.KanaToRomaji("バナナ") // -> banana
    s = kana.KanaToRomaji("かんじ") // -> kanji

### Convert romaji to hiragana or katakana:

    s := kana.RomajiToHiragana("kanji") // -> かんじ
    s = kana.RomajiToKatakana("banana") // -> バナナ

### Tell whether strings are written with kana, kanji or latin characters:

    kana.IsLatin("banana") // -> true
    kana.IsLatin("バナナ") // -> false

    kana.IsKana("banana") // -> false
    kana.IsKana("バナナ") // -> true

    kana.IsKanji("banana") // -> false
    kana.IsKanji("減少") // -> true

### Normalize a romaji string to a standardized form (from the form given by Google Translate, for example):

    kana.NormalizeRomaji("Myūjikku") // -> myu-jikku
    kana.NormalizeRomaji("shitsuree") // -> shitsurei

Please feel free to use, contribute, and enjoy! You can also see this in action at [nihongo.io](https://nihongo.io).

##

- [Herman Schaaf](http://github.com/hermanschaaf)
- [Shawn Smith](http://github.com/shawnps)
