package pbt

import (
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic		int
		Want		string
	}{
		{"1 gets converted to I", 1, "I"},
		{"2 gets converted to II", 2, "II"},
		{"3 gets converted to III", 3, "III"},
		{"4 gets converted to IV(can't repeat more than 3 times)", 4, "IV"},
		{"5 gets converted to V", 5, "V"},
		{"6 gets converted to VI", 6, "VI"},
		{"7 gets converted to VII", 7, "VII"},
		{"8 gets converted to VIII", 8, "VIII"},
		{"9 gets converted to IX", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
		{"100 gets converted to C", 100, "C"},
		{"150 gets converted to CL", 150, "CL"},
		{"300 gets converted to CCC", 300, "CCC"},
		{"400 gets converted to CD", 400, "CD"},
		{"469 gets converted to CDLXIX", 469, "CDLXIX"},
		{"500 gets converted to D", 500, "D"},
		{"1984 gets converted to MCMLXXXIV", 1984, "MCMLXXXIV"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T){
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}

}