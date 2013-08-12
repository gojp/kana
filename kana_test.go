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
	k := newKana()

	// some basic checks
	c.Check(k.kana_to_romaji("ああいうえお"), Equals, "aaiueo")
	c.Check(k.kana_to_romaji("かんじ"), Equals, "kanji")
	c.Check(k.kana_to_romaji("ちゃう"), Equals, "chau")
	c.Check(k.kana_to_romaji("はんのう"), Equals, "hannou")
	c.Check(k.kana_to_romaji("きょうじゅ"), Equals, "kyouju")

	// check that spacing is preserved
	c.Check(k.kana_to_romaji("な\nに	ぬ	ね	の"), Equals, "na\nni	nu	ne	no")

	// check that english text is preserved
	c.Check(k.kana_to_romaji("ばか dog"), Equals, "baka dog")

	// check double-consonants and long vowels
	c.Check(k.kana_to_romaji("きった"), Equals, "kitta")
}

func (s *KanaSuite) TestKatakanaToRomaji(c *C) {
	k := newKana()

	// basic tests
	c.Check(k.kana_to_romaji("バナナ"), Equals, "banana")
	c.Check(k.kana_to_romaji("カンジ"), Equals, "kanji")

	// check that r is preferred
	c.Check(k.kana_to_romaji("テレビ"), Equals, "terebi")

	// check english + katakana mix
	c.Check(k.kana_to_romaji("baking バナナ pancakes"), Equals, "baking banana pancakes")

	// check that double-consonants and long vowels get converted correctly
	c.Check(k.kana_to_romaji("ベッド"), Equals, "beddo")
	c.Check(k.kana_to_romaji("モーター"), Equals, "mootaa")
}

func (s *KanaSuite) TestRomajiToKatakana(c *C) {
	k := newKana()

	// basic tests
	c.Check(k.romaji_to_katakana("banana"), Equals, "バナナ")
	c.Check(k.romaji_to_katakana("rajio"), Equals, "ラジオ")

	// test r/l equality
	c.Check(k.romaji_to_katakana("terebi"), Equals, "テレビ")
	c.Check(k.romaji_to_katakana("telebi"), Equals, "テレビ")
	c.Check(k.romaji_to_katakana("furi-ta-"), Equals, "フリーター")

}
