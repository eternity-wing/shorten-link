package base62

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type symbolTestCase struct {
	character rune
	value     int
}

type stringTestCase struct {
	str   string
	value int
}

var validSymbols = []symbolTestCase{
	{character: '0', value: 0},
	{character: '5', value: 5},
	{character: 'A', value: 10},
	{character: 'Y', value: 34},
	{character: 'a', value: 36},
	{character: 'z', value: 61},
}

var validStringCases = []stringTestCase{
	{str: "100", value: 3844},
	{str: "ABC", value: 39134},
	{str: "00Z", value: 35},
}

func TestConvertor_ValidStringDecode(t *testing.T) {
	convertor := Convertor{}

	for _, tc := range validStringCases {
		assert.Equal(t, convertor.Decode(tc.str), tc.value)
	}

}

func TestConvertor_InValidStringDecode(t *testing.T) {
	convertor := Convertor{}
	tc := stringTestCase{str: "00-", value: -1}
	assert.Panics(t, func() { convertor.Decode(tc.str) }, "Expected panic for invalid string, but nothing happened")
}

func TestConvertor_Encode(t *testing.T) {
	convertor := Convertor{}

	for _, tc := range validStringCases {
		assert.Equal(t, strings.TrimLeft(tc.str, "0"), convertor.Encode(tc.value))
	}

}

func TestConvertor_getSymbolValueForValidSymbols(t *testing.T) {
	convertor := Convertor{}
	for _, tc := range validSymbols {
		assert.Equal(t, convertor.getSymbolValue(tc.character), tc.value)
	}

}

func TestConvertor_getSymbolValueForInvalidSymbol(t *testing.T) {
	convertor := Convertor{}
	tc := symbolTestCase{character: '*', value: -1}
	assert.Panics(t, func() { convertor.getSymbolValue(tc.character) }, "Expected panic for invalid symbol, but nothing happened")
}

func TestConvertor_getSymbolOfInvalidNumber(t *testing.T) {
	convertor := Convertor{}
	tc := symbolTestCase{character: '*', value: -1}
	assert.Panics(t, func() { convertor.getSymbolOfNumber(tc.value) }, "Expected panic for invalid number, but nothing happened")
}
