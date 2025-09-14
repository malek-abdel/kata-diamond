package main

import (
	"bufio"
	"fmt"
	"kata/diamond/builder"

	// To switch between the two implementations, uncomment the desired import and comment the other one.
	// "kata/diamond/builder/grid"
	"kata/diamond/builder/string"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a capital letter (from A to Z):")
	letter, _ := reader.ReadString('\n')

	var printer builder.DiamondBuilder
	printer = string.StringDiamondBuilder{}
	// To use the grid implementation, comment the above line and uncomment the line below.
	// printer = grid.GridDiamondBuilder{}

	diamond, err := printer.Print(letter[0 : len(letter)-1])
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(diamond)
}
