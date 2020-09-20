package base62

import (
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
		if val := convertor.Decode(tc.str); val != tc.value {
			t.Errorf("Expect return %d for %s, but return %d", tc.value, tc.str, val)
		}
	}

}

func TestConvertor_InValidStringDecode(t *testing.T) {
	convertor := Convertor{}
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected panic for invalid string, but nothing happened")
		}
	}()
	tc := stringTestCase{str: "00-", value: -1}
	convertor.Decode(tc.str)

}

func TestConvertor_Encode(t *testing.T) {
	convertor := Convertor{}

	for _, tc := range validStringCases {
		val := convertor.Encode(tc.value)
		expectedStr := strings.TrimLeft(tc.str, "0")
		if val != expectedStr {
			t.Errorf("Expect return %s for %d, but return %s", expectedStr, tc.value, val)
		}
	}

}

func TestConvertor_getSymbolValueForValidSymbols(t *testing.T) {
	convertor := Convertor{}
	for _, tc := range validSymbols {
		val := convertor.getSymbolValue(tc.character)

		if val != tc.value {
			t.Errorf("Expect return %d for %c, but return %d", tc.value, tc.character, val)
		}
	}

}

func TestConvertor_getSymbolValueForInvalidSymbol(t *testing.T) {

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected panic for invalid symbol, but nothing happened")
		}
	}()

	convertor := Convertor{}
	tc := symbolTestCase{character: '*', value: -1}
	convertor.getSymbolValue(tc.character)
}

func TestConvertor_getSymbolOfInvalidNumber(t *testing.T) {

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Expected panic for invalid number, but nothing happened")
		}
	}()

	convertor := Convertor{}
	tc := symbolTestCase{character: '*', value: -1}
	convertor.getSymbolOfNumber(tc.value)

}
