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
	c.Check(k.KanaToRomaji("ああいうえお"), Equals, "aaiueo")
	c.Check(k.KanaToRomaji("かんじ"), Equals, "kanji")
	c.Check(k.KanaToRomaji("ちゃう"), Equals, "chau")
	c.Check(k.KanaToRomaji("はんのう"), Equals, "hannou")
	c.Check(k.KanaToRomaji("きょうじゅ"), Equals, "kyouju")

	// check that spacing is preserved
	c.Check(k.KanaToRomaji("な\nに	ぬ	ね	の"), Equals, "na\nni	nu	ne	no")

	// check that english text is preserved
	c.Check(k.KanaToRomaji("ばか dog"), Equals, "baka dog")

	// check double-consonants and long vowels
	c.Check(k.KanaToRomaji("きった"), Equals, "kitta")
}

func (s *KanaSuite) TestKatakanaToRomaji(c *C) {
	k := NewKana()

	// basic tests
	c.Check(k.KanaToRomaji("バナナ"), Equals, "banana")
	c.Check(k.KanaToRomaji("カンジ"), Equals, "kanji")

	// check that r is preferred
	c.Check(k.KanaToRomaji("テレビ"), Equals, "terebi")

	// check english + katakana mix
	c.Check(k.KanaToRomaji("baking バナナ pancakes"), Equals, "baking banana pancakes")

	// check that double-consonants and long vowels get converted correctly
	c.Check(k.KanaToRomaji("ベッド"), Equals, "beddo")
	c.Check(k.KanaToRomaji("モーター"), Equals, "mootaa")
}

func (s *KanaSuite) TestRomajiToKatakana(c *C) {
	k := NewKana()

	// basic tests
	c.Check(k.RomajiToKatakana("banana"), Equals, "バナナ")
	c.Check(k.RomajiToKatakana("rajio"), Equals, "ラジオ")
	c.Check(k.RomajiToKatakana("terebi"), Equals, "テレビ")
	c.Check(k.RomajiToKatakana("furi-ta-"), Equals, "フリーター")
	c.Check(k.RomajiToKatakana("fa-suto"), Equals, "ファースト")
	c.Check(k.RomajiToKatakana("fesutibaru"), Equals, "フェスティバル")
	c.Check(k.RomajiToKatakana("ryukkusakku"), Equals, "リュックサック")
	c.Check(k.RomajiToKatakana("myu-jikku"), Equals, "ミュージック")
	c.Check(k.RomajiToKatakana("nyanda"), Equals, "ニャンダ")
	c.Check(k.RomajiToKatakana("hyakumeootokage"), Equals, "ヒャクメオオトカゲ")
}

func (s *KanaSuite) TestRomajiToHiragana(c *C) {
	k := NewKana()

	c.Check(k.RomajiToHiragana("banana"), Equals, "ばなな")
	c.Check(k.RomajiToHiragana("hiragana"), Equals, "ひらがな")
	c.Check(k.RomajiToHiragana("suppai"), Equals, "すっぱい")
	c.Check(k.RomajiToHiragana("konnichiha"), Equals, "こんにちは")
	c.Check(k.RomajiToHiragana("zouryou"), Equals, "ぞうりょう")
	c.Check(k.RomajiToHiragana("myaku"), Equals, "みゃく")
	c.Check(k.RomajiToHiragana("nyanko"), Equals, "にゃんこ")
	c.Check(k.RomajiToHiragana("hyaku"), Equals, "ひゃく")
}

func (s *KanaSuite) TestIsLatin(c *C) {
	k := NewKana()

	c.Check(k.IsLatin("banana"), Equals, true)
	c.Check(k.IsLatin("a sd ds ds"), Equals, true)
	c.Check(k.IsLatin("ばなな"), Equals, false)
	c.Check(k.IsLatin("ファースト"), Equals, false)
}

func (s *KanaSuite) TestIsKana(c *C) {
	k := NewKana()

	c.Check(k.IsKana("ばなな"), Equals, true)
	c.Check(k.IsKana("ファースト"), Equals, true)
	c.Check(k.IsKana("test"), Equals, false)
}

func (s *KanaSuite) TestIsKanji(c *C) {
	k := NewKana()

	c.Check(k.IsKanji("ばなな"), Equals, false)
	c.Check(k.IsKanji("ファースト"), Equals, false)
	c.Check(k.IsKanji("test"), Equals, false)
	c.Check(k.IsKanji("路加"), Equals, true)
	c.Check(k.IsKanji("減少"), Equals, true)
}
