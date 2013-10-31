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
	// some basic checks
	c.Check(KanaToRomaji("ああいうえお"), Equals, "aaiueo")
	c.Check(KanaToRomaji("かんじ"), Equals, "kanji")
	c.Check(KanaToRomaji("ちゃう"), Equals, "chau")
	c.Check(KanaToRomaji("はんのう"), Equals, "hannou")
	c.Check(KanaToRomaji("きょうじゅ"), Equals, "kyouju")
	c.Check(KanaToRomaji("ぜんいん"), Equals, "zennin")
	c.Check(KanaToRomaji("はんのう"), Equals, "hannnou")
	c.Check(KanaToRomaji("はんおう"), Equals, "hannou")

	// check that spacing is preserved
	c.Check(KanaToRomaji("な\nに	ぬ	ね	の"), Equals, "na\nni	nu	ne	no")

	// check that english text is preserved
	c.Check(KanaToRomaji("ばか dog"), Equals, "baka dog")

	// check double-consonants and long vowels
	c.Check(KanaToRomaji("きった"), Equals, "kitta")
}

func (s *KanaSuite) TestKatakanaToRomaji(c *C) {
	// basic tests
	c.Check(KanaToRomaji("バナナ"), Equals, "banana")
	c.Check(KanaToRomaji("カンジ"), Equals, "kanji")

	// check that r is preferred
	c.Check(KanaToRomaji("テレビ"), Equals, "terebi")

	// check english + katakana mix
	c.Check(KanaToRomaji("baking バナナ pancakes"), Equals, "baking banana pancakes")

	// check that double-consonants and long vowels get converted correctly
	c.Check(KanaToRomaji("ベッド"), Equals, "beddo")
	c.Check(KanaToRomaji("モーター"), Equals, "mo-ta-")

	// check random input
	c.Check(KanaToRomaji("ＣＤプレーヤー"), Equals, "ＣＤpure-ya-")
	c.Check(KanaToRomaji("オーバーヘッドキック"), Equals, "o-ba-heddokikku")
}

func (s *KanaSuite) TestRomajiToKatakana(c *C) {
	// basic tests
	c.Check(RomajiToKatakana("banana"), Equals, "バナナ")
	c.Check(RomajiToKatakana("rajio"), Equals, "ラジオ")
	c.Check(RomajiToKatakana("terebi"), Equals, "テレビ")
	c.Check(RomajiToKatakana("furi-ta-"), Equals, "フリーター")
	c.Check(RomajiToKatakana("fa-suto"), Equals, "ファースト")
	c.Check(RomajiToKatakana("fesutibaru"), Equals, "フェスティバル")
	c.Check(RomajiToKatakana("ryukkusakku"), Equals, "リュックサック")
	c.Check(RomajiToKatakana("myu-jikku"), Equals, "ミュージック")
	c.Check(RomajiToKatakana("nyanda"), Equals, "ニャンダ")
	c.Check(RomajiToKatakana("hyakumeootokage"), Equals, "ヒャクメオオトカゲ")

	// shouldn't do anything:
	c.Check(RomajiToKatakana("ＣＤプレーヤー"), Equals, "ＣＤプレーヤー")
}

func (s *KanaSuite) TestRomajiToHiragana(c *C) {
	c.Check(RomajiToHiragana("banana"), Equals, "ばなな")
	c.Check(RomajiToHiragana("hiragana"), Equals, "ひらがな")
	c.Check(RomajiToHiragana("suppai"), Equals, "すっぱい")
	c.Check(RomajiToHiragana("konnichiha"), Equals, "こんにちは")
	c.Check(RomajiToHiragana("zouryou"), Equals, "ぞうりょう")
	c.Check(RomajiToHiragana("myaku"), Equals, "みゃく")
	c.Check(RomajiToHiragana("nyanko"), Equals, "にゃんこ")
	c.Check(RomajiToHiragana("hyaku"), Equals, "ひゃく")
	c.Check(RomajiToHiragana("motoduku"), Equals, "もとづく")
	c.Check(RomajiToHiragana("zennin"), Equals, "ぜんいん")
	c.Check(RomajiToHiragana("hannnou"), Equals, "はんのう")
	c.Check(RomajiToHiragana("hannou"), Equals, "はんおう")

	// shouldn't do anything:
	c.Check(RomajiToHiragana("ＣＤプレーヤー"), Equals, "ＣＤプレーヤー")
}

func (s *KanaSuite) TestIsLatin(c *C) {
	c.Check(IsLatin("banana"), Equals, true)
	c.Check(IsLatin("a sd ds ds"), Equals, true)
	c.Check(IsLatin("ばなな"), Equals, false)
	c.Check(IsLatin("ファースト"), Equals, false)
	c.Check(IsLatin("myu-jikku"), Equals, true)

	c.Check(IsLatin("ＣＤプレーヤー"), Equals, false)
}

func (s *KanaSuite) TestIsKana(c *C) {
	c.Check(IsKana("ばなな"), Equals, true)
	c.Check(IsKana("ファースト"), Equals, true)
	c.Check(IsKana("test"), Equals, false)
}

func (s *KanaSuite) TestIsKanji(c *C) {
	c.Check(IsKanji("ばなな"), Equals, false)
	c.Check(IsKanji("ファースト"), Equals, false)
	c.Check(IsKanji("test"), Equals, false)
	c.Check(IsKanji("路加"), Equals, true)
	c.Check(IsKanji("減少"), Equals, true)
}

func (s *KanaSuite) TestNormalizeRomaji(c *C) {
	c.Check(NormalizeRomaji("myuujikku"), Equals, "myu-jikku")
	c.Check(NormalizeRomaji("Myūjikku"), Equals, "myu-jikku")
	c.Check(NormalizeRomaji("Banana"), Equals, "banana")
	c.Check(NormalizeRomaji("shitsuree"), Equals, "shitsurei")
	c.Check(NormalizeRomaji("減少"), Equals, "減少")
	c.Check(NormalizeRomaji("myuujikku Myūjikku Banana shitsuree"), Equals, "myu-jikku myu-jikku banana shitsurei")
}
