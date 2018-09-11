package main

import (
	"fmt"
)

// func main() {
// 	numberOfCases := readerpkg.ReadInt()
// 	for index := 0; index < int(numberOfCases); index++ {
// 		inputString := readerpkg.ReadStr()
// 		printLetters(inputString)
// 	}
// }

func main() {
	var numberOfCases int
	var inputString string
	fmt.Scan(&numberOfCases)
	for index := 0; index < int(numberOfCases); index++ {
		fmt.Scan(&inputString)
		printLetters(inputString)
	}
}

func printLetters(inputString string) {
	oddLetters := ""
	evenLetters := ""
	for index := 0; index < len(inputString); index++ {
		if index%2 == 0 {
			evenLetters = evenLetters + string(inputString[index])
		} else {
			oddLetters = oddLetters + string(inputString[index])
		}
	}
	fmt.Println(evenLetters + " " + oddLetters)
}
