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
	k := new(Kana)
	k.initialize()

	// some basic checks
	c.Check(k.kana_to_romaji("ああいうえお"), Equals, "aaiueo")
	c.Check(k.kana_to_romaji("かんじ"), Equals, "kanji")
	c.Check(k.kana_to_romaji("ちゃう"), Equals, "chau")
	c.Check(k.kana_to_romaji("はんのう"), Equals, "hannou")

	// check that spacing is preserved
	c.Check(k.kana_to_romaji("な	に	ぬ	ね	の"), Equals, "na	ni	nu	ne	no")

	// check that english text is preserved
	c.Check(k.kana_to_romaji("ばか dog"), Equals, "baka dog")

	// check that double-consonants and long vowels get converted correctly
	// c.Check(k.kana_to_romaji(""), Equals, "beddo")
}

// func (s *KanaSuite) TestKatakanaToRomaji(c *C) {
// 	k := new(Kana)
// 	k.initialize()
// 	c.Check(k.kana_to_romaji("バナナ"), Equals, "banana")
// 	c.Check(k.kana_to_romaji("カンジ"), Equals, "kanji")
// }
