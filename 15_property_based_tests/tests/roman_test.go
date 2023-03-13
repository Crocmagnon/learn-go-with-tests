package tests

import (
	"fmt"
	roman "learn-go-with-tests/15_property_based_tests"
	"testing"
	"testing/quick"
)

type testCase struct {
	num   uint16
	roman string
}

var tests = []testCase{
	{num: 1, roman: "I"},
	{num: 2, roman: "II"},
	{num: 3, roman: "III"},
	{num: 4, roman: "IV"},
	{num: 5, roman: "V"},
	{num: 6, roman: "VI"},
	{num: 7, roman: "VII"},
	{num: 8, roman: "VIII"},
	{num: 9, roman: "IX"},
	{num: 10, roman: "X"},
	{num: 14, roman: "XIV"},
	{num: 18, roman: "XVIII"},
	{num: 20, roman: "XX"},
	{num: 39, roman: "XXXIX"},
	{num: 40, roman: "XL"},
	{num: 47, roman: "XLVII"},
	{num: 49, roman: "XLIX"},
	{num: 50, roman: "L"},
	{num: 100, roman: "C"},
	{num: 90, roman: "XC"},
	{num: 400, roman: "CD"},
	{num: 500, roman: "D"},
	{num: 900, roman: "CM"},
	{num: 1000, roman: "M"},
	{num: 1984, roman: "MCMLXXXIV"},
	{num: 3999, roman: "MMMCMXCIX"},
	{num: 2014, roman: "MMXIV"},
	{num: 1006, roman: "MVI"},
	{num: 798, roman: "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v gets converted to %v", tt.num, tt.roman), func(t *testing.T) {
			if got := roman.ConvertToRoman(tt.num); got != tt.roman {
				t.Errorf("ConvertToRoman() = %v, want %v", got, tt.roman)
			}
		})
	}
}

func TestConvertFromRoman(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v gets converted to %v", tt.roman, tt.num), func(t *testing.T) {
			if got := roman.ConvertFromRoman(tt.roman); got != tt.num {
				t.Errorf("ConvertFromRoman() = %v, want %v", got, tt.num)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		romanResult := roman.ConvertToRoman(arabic)
		fromRoman := roman.ConvertFromRoman(romanResult)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error(err)
	}
}
