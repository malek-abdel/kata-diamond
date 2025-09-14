package string

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type StringDiamondBuilder struct{}

func (s StringDiamondBuilder) Print(letter string) (string, error) {
	if utf8.RuneCountInString(letter) != 1 {
		panic("You should pass one letter")
	}
	runeLetter, _ := utf8.DecodeRuneInString(letter[0:])
	if runeLetter <= 'A' || runeLetter > 'Z' {
		return "", errors.New("You should pass a capital letter between A and Z")
	}

	alphabetPosition := int(runeLetter - 'A')

	return PrintLetter('A', alphabetPosition), nil
}

func PrintLetter(letter rune, leftSpaces int) string {
	row := []rune(strings.Repeat(" ", leftSpaces*2+(int(letter-'A'))*2+1))
	row[leftSpaces] = letter
	row[len(row)-leftSpaces-1] = letter

	if leftSpaces == 0 {
		return string(row)
	}

	return strings.Join([]string{string(row), PrintLetter(letter+1, leftSpaces-1), string(row)}, "\n")
}

func Print(letter string) (string, error) {
	if utf8.RuneCountInString(letter) != 1 {
		panic("You should pass one letter")
	}
	runeLetter, _ := utf8.DecodeRuneInString(letter[0:])
	if runeLetter <= 'A' || runeLetter > 'Z' {
		return "", errors.New("You should pass a capital letter between A and Z")
	}

	alphabetPosition := int(runeLetter - 'A')
	// diamondDimension := alphabetPosition*2 + 1

	return PrintLetter('A', alphabetPosition), nil
	// Grid Option
	// diamondGrid := make([]string, diamondDimension)

	// for i := range alphabetPosition + 1 {
	// 	currentLetter := rune('A' + i)
	// 	currentRow := []rune(strings.Repeat(" ", diamondDimension))
	// 	currentRow[alphabetPosition-i] = currentLetter
	// 	currentRow[alphabetPosition+i] = currentLetter
	// 	diamondGrid[i] = string(currentRow)
	// 	diamondGrid[alphabetPosition*2-i] = string(currentRow)
	// }

	// return strings.Join(diamondGrid, "\n")
}
