package base62

import (
	"math"
	"strings"
)

type Convertor struct {
}

func (c Convertor) GetBaseNumber() int {
	return 62
}

func (c Convertor) Encode(num int) string {
	str := ""
	for true {
		r := num % c.GetBaseNumber()
		num = num / c.GetBaseNumber()
		symbol := c.getSymbolOfNumber(r)
		str = string(symbol) + str
		if num == 0 {
			break
		}
	}
	return str
}
func (c Convertor) Decode(str string) int {
	decodedValue := 0
	str = strings.TrimLeft(str, "0")
	strLen := len(str)
	for i, char := range str {
		symbolValue := c.getSymbolValue(char)
		pow := strLen - (i + 1)
		decodedValue += symbolValue * int(math.Pow(float64(c.GetBaseNumber()), float64(pow)))
	}
	return decodedValue
}

const (
	ZeroAsciCode         rune = '0'
	NumberOfAlphabets    int  = 26
	CapitalLettersOffset      = int('A'-'9') - 1
	SmallLettersOffset        = int('a'-'9') - NumberOfAlphabets - 1
)

func (c Convertor) getSymbolValue(char rune) int {
	var offset = int(char - ZeroAsciCode)
	switch {
	case char >= '0' && char <= '9':
		return offset
	case char >= 'A' && char <= 'Z':
		return offset - CapitalLettersOffset
	case char >= 'a' && char <= 'z':
		return offset - SmallLettersOffset
	default:
		panic("Unexpected character!\nAcceptable characters are numerical and alphabetical")
	}
}

func (c Convertor) getSymbolOfNumber(num int) rune {
	var offset = rune(num) + ZeroAsciCode
	switch {
	case num >= 0 && num <= 9:
		return offset
	case num > 9 && num <= 9+NumberOfAlphabets:
		return offset + rune(CapitalLettersOffset)
	case num > NumberOfAlphabets && num <= (9+2*NumberOfAlphabets):
		return offset + rune(SmallLettersOffset)
	default:
		panic("Unexpected Number!\nAcceptable numbers are between 0 and 61")
	}
}
