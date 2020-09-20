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

func (c Convertor) Encode(num int) (string, error) {
	str := ""
	for true {
		r := num % c.GetBaseNumber()
		num = num / c.GetBaseNumber()
		symbol, err := c.getSymbolOfNumber(r)
		if err != nil {
			return "", err
		}
		str = string(symbol) + str
		if num == 0 {
			break
		}
	}
	return str, nil
}
func (c Convertor) Decode(str string) (int, error) {
	decodedValue := 0
	str = strings.TrimLeft(str, "0")
	strLen := len(str)
	for i, char := range str {
		symbolValue, err := c.getSymbolValue(char)
		if err != nil {
			return -1, err
		}
		pow := strLen - (i + 1)
		decodedValue += symbolValue * int(math.Pow(float64(c.GetBaseNumber()), float64(pow)))
	}
	return decodedValue, nil
}

const (
	ZeroAsciCode         rune = '0'
	NumberOfAlphabets    int  = 26
	CapitalLettersOffset      = int('A'-'9') - 1
	SmallLettersOffset        = int('a'-'9') - NumberOfAlphabets - 1
)

func (c Convertor) getSymbolValue(char rune) (int, error) {
	var offset = int(char - ZeroAsciCode)
	switch {
	case char >= '0' && char <= '9':
		return offset, nil
	case char >= 'A' && char <= 'Z':
		return offset - CapitalLettersOffset, nil
	case char >= 'a' && char <= 'z':
		return offset - SmallLettersOffset, nil
	default:
		return -1, &UnexpectedCharacterError{}
	}
}

func (c Convertor) getSymbolOfNumber(num int) (rune, error) {
	var offset = rune(num) + ZeroAsciCode
	switch {
	case num >= 0 && num <= 9:
		return offset, nil
	case num > 9 && num <= 9+NumberOfAlphabets:
		return offset + rune(CapitalLettersOffset), nil
	case num > NumberOfAlphabets && num <= (9+2*NumberOfAlphabets):
		return offset + rune(SmallLettersOffset), nil
	default:
		return -1, &UnexpectedNumberError{}
	}
}
