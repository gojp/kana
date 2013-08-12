kana
=======

A golang library to convert hiragana to romaji, katakana to romaji, romaji to hiragana and romaji to katakana. 

Usage
-------

To use *kana*, you'll have to first initialize it:

    import "github.com/hermanschaaf/kana"
    ...
    k := kana.NewKana()

Here are a couple of examples of how you could use *kana*:

    s := k.kana_to_romaji("バナナ") // -> banana
    s = k.kana_to_romaji("かんじ") // -> kanji
