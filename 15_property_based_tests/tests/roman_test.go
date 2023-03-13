package tests

import (
	"fmt"
	roman "learn-go-with-tests/15_property_based_tests"
	"testing"
)

type testCase struct {
	num  int
	want string
}

var tests = []testCase{
	{num: 1, want: "I"},
	{num: 2, want: "II"},
	{num: 3, want: "III"},
	{num: 4, want: "IV"},
	{num: 5, want: "V"},
	{num: 6, want: "VI"},
	{num: 7, want: "VII"},
	{num: 8, want: "VIII"},
	{num: 9, want: "IX"},
	{num: 10, want: "X"},
	{num: 14, want: "XIV"},
	{num: 18, want: "XVIII"},
	{num: 20, want: "XX"},
	{num: 39, want: "XXXIX"},
	{num: 40, want: "XL"},
	{num: 47, want: "XLVII"},
	{num: 49, want: "XLIX"},
	{num: 50, want: "L"},
	{num: 100, want: "C"},
	{num: 90, want: "XC"},
	{num: 400, want: "CD"},
	{num: 500, want: "D"},
	{num: 900, want: "CM"},
	{num: 1000, want: "M"},
	{num: 1984, want: "MCMLXXXIV"},
	{num: 3999, want: "MMMCMXCIX"},
	{num: 2014, want: "MMXIV"},
	{num: 1006, want: "MVI"},
	{num: 798, want: "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d gets converted to %v", tt.num, tt.want), func(t *testing.T) {
			if got := roman.ConvertToRoman(tt.num); got != tt.want {
				t.Errorf("ConvertToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
