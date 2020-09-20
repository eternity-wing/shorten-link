package base62

import (
	"reflect"
	"strings"
	"testing"
)

type symbolTestCase struct {
	character rune
	value     int
	err       error
}

type stringTestCase struct {
	str   string
	value int
	err   error
}

var validSymbols = []symbolTestCase{
	{character: '0', value: 0, err: nil},
	{character: '5', value: 5, err: nil},
	{character: 'A', value: 10, err: nil},
	{character: 'Y', value: 34, err: nil},
	{character: 'a', value: 36, err: nil},
	{character: 'z', value: 61, err: nil},
}

var validStringCases = []stringTestCase{
	{str: "100", value: 3844, err: nil},
	{str: "ABC", value: 39134, err: nil},
	{str: "00Z", value: 35, err: nil},
}

func TestConvertor_Decode(t *testing.T) {
	convertor := Convertor{}

	testCases := []stringTestCase{
		{str: "00-", value: -1, err: &UnexpectedCharacterError{}},
	}

	copy(testCases, validStringCases)

	for _, testCase := range testCases {
		val, err := convertor.Decode(testCase.str)

		if testCase.err == nil && val != testCase.value {
			t.Errorf("Expect return %d for %s, but return %d", testCase.value, testCase.str, val)
		} else if reflect.TypeOf(err) != reflect.TypeOf(testCase.err) {
			t.Errorf("Expected error of type %T but %T", testCase.err, err)
		}

	}

}

func TestConvertor_Encode(t *testing.T) {
	convertor := Convertor{}

	for _, testCase := range validStringCases {
		val, err := convertor.Encode(testCase.value)
		expectedStr := strings.TrimLeft(testCase.str, "0")
		if testCase.err == nil && val != expectedStr {
			t.Errorf("Expect return %s for %d, but return %s", expectedStr, testCase.value, val)
		} else if reflect.TypeOf(err) != reflect.TypeOf(testCase.err) {
			t.Errorf("Expected error of type %T but %T", testCase.err, err)
		}

	}

}

func TestConvertor_getSymbolValue(t *testing.T) {
	convertor := Convertor{}
	testCases := []symbolTestCase{
		{character: '*', value: -1, err: &UnexpectedCharacterError{}},
	}
	copy(testCases, validSymbols)
	for _, testCase := range testCases {
		val, err := convertor.getSymbolValue(testCase.character)

		if testCase.err != nil && val != testCase.value {
			t.Errorf("Expect return %d for %c, but return %d", testCase.value, testCase.character, val)
		} else if reflect.TypeOf(err) != reflect.TypeOf(testCase.err) {
			t.Errorf("Expected error of type %T but %T", testCase.err, err)
		}
	}

}

func TestConvertor_getSymbolOfNumber(t *testing.T) {
	convertor := Convertor{}
	testCases := []symbolTestCase{
		{character: '*', value: -1, err: &UnexpectedNumberError{}},
	}
	for _, testCase := range testCases {
		char, err := convertor.getSymbolOfNumber(testCase.value)

		if testCase.err == nil && char != testCase.character {
			t.Errorf("Expect return %c for %d, but return %c", testCase.character, testCase.value, char)
		} else if reflect.TypeOf(err) != reflect.TypeOf(testCase.err) {
			t.Errorf("Expected error of type %T but %T", testCase.err, err)
		}
	}

}
