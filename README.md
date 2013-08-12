kana
=======

Hiragana, Katakana to romaji and Romaji to Hiragana, Katakana converter library for Go 

Usage
-------

To use *kana*, you'll have to first initialize it:

    import (
        "github.com/hermanschaaf/kana"
    )
    
    ...
    
    k := kana.NewKana()

Here are a couple of examples of how you could use *kana*:

    s := k.kana_to_romaji("バナナ") // -> banana
    s = k.kana_to_romaji("かんじ") // -> kanji
