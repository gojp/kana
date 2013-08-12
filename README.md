kana
=======

Hiragana, Katakana to romaji and Romaji to Hiragana, Katakana converter library for Go 

Usage
-------

To use go-kana, you'll have to first initialize it:

    import (
        . "github.com/hermanschaaf/go-kana"
    )
    
    ...
    
    k := NewKana()

Here are a couple of examples of how you could use go-kana:

    s := k.kana_to_romaji("バナナ") // -> banana
    s = k.kana_to_romaji("かんじ") // -> kanji
