package kana

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type KanaSuite struct{}

var _ = Suite(&KanaSuite{})

func (s *KanaSuite) TestHiraganaToRomaji(c *C) {
	k := NewKana()

	// some basic checks
	c.Check(k.Kana_to_romaji("ああいうえお"), Equals, "aaiueo")
	c.Check(k.Kana_to_romaji("かんじ"), Equals, "kanji")
	c.Check(k.Kana_to_romaji("ちゃう"), Equals, "chau")
	c.Check(k.Kana_to_romaji("はんのう"), Equals, "hannou")
	c.Check(k.Kana_to_romaji("きょうじゅ"), Equals, "kyouju")

	// check that spacing is preserved
	c.Check(k.Kana_to_romaji("な\nに	ぬ	ね	の"), Equals, "na\nni	nu	ne	no")

	// check that english text is preserved
	c.Check(k.Kana_to_romaji("ばか dog"), Equals, "baka dog")

	// check double-consonants and long vowels
	c.Check(k.Kana_to_romaji("きった"), Equals, "kitta")
}

func (s *KanaSuite) TestKatakanaToRomaji(c *C) {
	k := NewKana()

	// basic tests
	c.Check(k.Kana_to_romaji("バナナ"), Equals, "banana")
	c.Check(k.Kana_to_romaji("カンジ"), Equals, "kanji")

	// check that r is preferred
	c.Check(k.Kana_to_romaji("テレビ"), Equals, "terebi")

	// check english + katakana mix
	c.Check(k.Kana_to_romaji("baking バナナ pancakes"), Equals, "baking banana pancakes")

	// check that double-consonants and long vowels get converted correctly
	c.Check(k.Kana_to_romaji("ベッド"), Equals, "beddo")
	c.Check(k.Kana_to_romaji("モーター"), Equals, "mootaa")
}

func (s *KanaSuite) TestRomajiToKatakana(c *C) {
	k := NewKana()

	// basic tests
	c.Check(k.Romaji_to_katakana("banana"), Equals, "バナナ")
	c.Check(k.Romaji_to_katakana("rajio"), Equals, "ラジオ")
	c.Check(k.Romaji_to_katakana("terebi"), Equals, "テレビ")
	c.Check(k.Romaji_to_katakana("furi-ta-"), Equals, "フリーター")
	c.Check(k.Romaji_to_katakana("fa-suto"), Equals, "ファースト")
	c.Check(k.Romaji_to_katakana("fesutibaru"), Equals, "フェスティバル")
}

func (s *KanaSuite) TestRomajiToHiragana(c *C) {
	k := NewKana()

	c.Check(k.Romaji_to_hiragana("banana"), Equals, "ばなな")
	c.Check(k.Romaji_to_hiragana("hiragana"), Equals, "ひらがな")
	c.Check(k.Romaji_to_hiragana("suppai"), Equals, "すっぱい")
	c.Check(k.Romaji_to_hiragana("konnichiha"), Equals, "こんにちは")
}
