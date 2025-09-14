package grid

import (
	"errors"
	"strings"
	"unicode/utf8"
)

type GridDiamondBuilder struct{}

func (g GridDiamondBuilder) Print(letter string) (string, error) {
	if utf8.RuneCountInString(letter) != 1 {
		panic("You should pass one letter")
	}
	runeLetter, _ := utf8.DecodeRuneInString(letter[0:])
	if runeLetter <= 'A' || runeLetter > 'Z' {
		return "", errors.New("You should pass a capital letter between A and Z")
	}

	alphabetPosition := int(runeLetter - 'A')
	diamondDimension := alphabetPosition*2 + 1

	// Grid Option
	diamondGrid := make([]string, diamondDimension)

	for i := range alphabetPosition + 1 {
		currentLetter := rune('A' + i)
		currentRow := []rune(strings.Repeat(" ", diamondDimension))
		currentRow[alphabetPosition-i] = currentLetter
		currentRow[alphabetPosition+i] = currentLetter
		diamondGrid[i] = string(currentRow)
		diamondGrid[alphabetPosition*2-i] = string(currentRow)
	}

	return strings.Join(diamondGrid, "\n"), nil
}
