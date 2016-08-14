[![Build Status](https://travis-ci.org/gojp/kana.png?branch=master)](https://travis-ci.org/gojp/kana) [![go report card](https://goreportcard.com/badge/github.com/gojp/kana)](http://goreportcard.com/report/github.com/gojp/kana)

kana
=======

A golang library to convert hiragana to romaji, katakana to romaji, romaji to hiragana and romaji to katakana. 

Installation
-------

Simply install with `go get`:

    go get github.com/gojp/kana

Usage
-------

To use *kana*, you'll have to import it:

    import "github.com/gojp/kana"
    ...
    k := kana.NewKana()

*kana* can do many things. It can convert hiragana or katakana to romaji:

    s := kana.KanaToRomaji("バナナ") // -> banana
    s = kana.KanaToRomaji("かんじ") // -> kanji

It can convert romaji to hiragana or katakana:

    s := kana.RomajiToHiragana("kanji") // -> かんじ
    s = kana.RomajiToKatakana("banana") // -> バナナ

It can tell you whether strings are written with kana, kanji or latin characters:

    kana.IsLatin("banana") // -> true
    kana.IsLatin("バナナ") // -> false

    kana.IsKana("banana") // -> false
    kana.IsKana("バナナ") // -> true

    kana.IsKanji("banana") // -> false
    kana.IsKanji("減少") // -> true

It can also normalize a given romaji string to a more standardized form (from the form given by Google Translate, for example):

    kana.NormalizeRomaji("Myūjikku") // -> myu-jikku
    kana.NormalizeRomaji("shitsuree") // -> shitsurei

Please feel free to use, contribute, and enjoy! You can also see this in action at [nihongo.io](http://nihongo.io).

Contributors
-------

- [Herman Schaaf](http://github.com/hermanschaaf) (author)
- [Shawn Smith](http://github.com/shawnps)
