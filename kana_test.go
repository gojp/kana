package kana

import "testing"

type kanaTest struct {
	orig, want string
}

var hiraganaToRomajiTests = []kanaTest{
	{"ああいうえお", "aaiueo"},
	{"かんじ", "kanji"},
	{"ちゃう", "chau"},
	{"きょうじゅ", "kyouju"},
	{"な\nに	ぬ	ね	の", "na\nni	nu	ne	no"},
	{"ばか dog", "baka dog"},
	{"きった", "kitta"},
	{"はんのう", "hannnou"},
	{"ぜんいん", "zennin"},
	{"んい", "nni"},
	{"はんのう", "hannnou"},
	{"はんおう", "hannou"},
	{"あうでぃ", "audexi"},
}

func TestHiraganaToRomaji(t *testing.T) {
	for _, tt := range hiraganaToRomajiTests {
		if got := KanaToRomaji(tt.orig); got != tt.want {
			t.Errorf("KanaToRomaji(%q) = %q, want %q", tt.orig, got, tt.want)
		}
	}
}

var katakanaToRomajiTests = []kanaTest{
	{"バナナ", "banana"},
	{"カンジ", "kanji"},
	{"テレビ", "terebi"},
	{"baking バナナ pancakes", "baking banana pancakes"},
	{"ベッド", "beddo"},
	{"モーター", "mo-ta-"},
	{"ＣＤプレーヤー", "ＣＤpure-ya-"},
	{"オーバーヘッドキック", "o-ba-heddokikku"},
	{"ハンノウ", "hannnou"},
	{"アウディ", "audexi"},
}

func TestKatakanaToRomaji(t *testing.T) {
	for _, tt := range katakanaToRomajiTests {
		if got := KanaToRomaji(tt.orig); got != tt.want {
			t.Errorf("KanaToRomaji(%q) = %q, want %q", tt.orig, got, tt.want)
		}
	}
}

var romajiToKatakanaTests = []kanaTest{
	{"banana", "バナナ"},
	{"rajio", "ラジオ"},
	{"terebi", "テレビ"},
	{"furi-ta-", "フリーター"},
	{"fa-suto", "ファースト"},
	{"fesutibaru", "フェスティバル"},
	{"ryukkusakku", "リュックサック"},
	{"myu-jikku", "ミュージック"},
	{"nyanda", "ニャンダ"},
	{"hyakumeootokage", "ヒャクメオオトカゲ"},
	{"ＣＤプレーヤー", "ＣＤプレーヤー"},
	{"cheri-", "チェリー"},
}

func TestRomajiToKatakana(t *testing.T) {
	for _, tt := range romajiToKatakanaTests {
		if got := RomajiToKatakana(tt.orig); got != tt.want {
			t.Errorf("RomajiToKatakana(%q) = %q, want %q", tt.orig, got, tt.want)
		}
	}
}

var romajiToHiraganaTests = []kanaTest{
	{"banana", "ばなな"},
	{"hiragana", "ひらがな"},
	{"suppai", "すっぱい"},
	{"konnnichiha", "こんにちは"},
	{"zouryou", "ぞうりょう"},
	{"myaku", "みゃく"},
	{"nyanko", "にゃんこ"},
	{"hyaku", "ひゃく"},
	{"motoduku", "もとづく"},
	{"zenin", "ぜにん"},
	{"zennin", "ぜんいん"},
	{"hannnou", "はんのう"},
	{"hannou", "はんおう"},
	{"chuutohanpa", "ちゅうとはんぱ"},
	{"ＣＤプレーヤー", "ＣＤプレーヤー"},
	// {"meccha", "めっちゃ"},
	{"che", "ちぇ"},
}

func TestRomajiToHiragana(t *testing.T) {
	for _, tt := range romajiToHiraganaTests {
		if got := RomajiToHiragana(tt.orig); got != tt.want {
			t.Errorf("RomajiToHiragana(%q) = %q, want %q", tt.orig, got, tt.want)
		}
	}
}

type typeTest struct {
	text  string
	valid bool
}

var isLatinTests = []typeTest{
	{"banana", true},
	{"a sd ds ds", true},
	{"ばなな", false},
	{"ファースト", false},
	{"myu-jikku", true},
	{"ＣＤプレーヤー", false},
}

func TestIsLatin(t *testing.T) {
	for _, tt := range isLatinTests {
		if got := IsLatin(tt.text); got != tt.valid {
			t.Errorf("IsLatin(%q) = %t, want %t", tt.text, got, tt.valid)
		}
	}
}

var isKanaTests = []typeTest{
	{"ばなな", true},
	{"ファースト", true},
	{"test", false},
}

func TestIsKana(t *testing.T) {
	for _, tt := range isKanaTests {
		if got := IsKana(tt.text); got != tt.valid {
			t.Errorf("IsKana(%q) = %t, want %t", tt.text, got, tt.valid)
		}
	}
}

var isKanjiTests = []typeTest{
	{"ばなな", false},
	{"ファースト", false},
	{"test", false},
	{"路加", true},
	{"減少", true},
}

func TestIsKanji(t *testing.T) {
	for _, tt := range isKanjiTests {
		if got := IsKanji(tt.text); got != tt.valid {
			t.Errorf("IsKanji(%q) = %t, want %t", tt.text, got, tt.valid)
		}
	}
}

var normalizeRomajiTests = []kanaTest{
	{"myuujikku", "myu-jikku"},
	{"Myūjikku", "myu-jikku"},
	{"Banana", "banana"},
	{"shitsuree", "shitsurei"},
	{"減少", "減少"},
	{"myuujikku Myūjikku Banana shitsuree", "myu-jikku myu-jikku banana shitsurei"},
}

func TestNormalizeRomaji(t *testing.T) {
	for _, tt := range normalizeRomajiTests {
		if got := NormalizeRomaji(tt.orig); got != tt.want {
			t.Errorf("NormalizeRomaji(%q) = %q, want %q", tt.orig, got, tt.want)
		}
	}
}
