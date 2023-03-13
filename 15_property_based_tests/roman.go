package roman

import (
	"errors"
	"strings"
)

var ErrNumberTooLarge = errors.New("number is too large")

func ConvertToRoman(arabic uint16) (string, error) {
	var result strings.Builder

	if arabic > 3999 {
		return "", ErrNumberTooLarge
	}

	for _, numeral := range allNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String(), nil
}

func ConvertFromRoman(roman string) uint16 {
	var total uint16 = 0
	for _, symbols := range windowedRoman(roman).Symbols() {
		total += allNumerals.ValueOf(symbols...)
	}
	return total
}

type numeral struct {
	Value  uint16
	Symbol string
}

type numerals []numeral

func (n numerals) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, num := range n {
		if num.Symbol == symbol {
			return num.Value
		}
	}
	return 0
}

func (n numerals) Exists(symbols ...byte) bool {
	return n.ValueOf(symbols...) > 0
}

var allNumerals = numerals{
	{Value: 1000, Symbol: "M"},
	{Value: 900, Symbol: "CM"},
	{Value: 500, Symbol: "D"},
	{Value: 400, Symbol: "CD"},
	{Value: 100, Symbol: "C"},
	{Value: 90, Symbol: "XC"},
	{Value: 50, Symbol: "L"},
	{Value: 40, Symbol: "XL"},
	{Value: 10, Symbol: "X"},
	{Value: 9, Symbol: "IX"},
	{Value: 5, Symbol: "V"},
	{Value: 4, Symbol: "IV"},
	{Value: 1, Symbol: "I"},
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractive(symbol) && allNumerals.Exists(symbol, w[i+1]) {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}
	return
}

func isSubtractive(symbol uint8) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
