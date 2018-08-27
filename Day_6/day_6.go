package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	numberOfCases, _ := strconv.ParseInt(readLine(reader), 10, 0)
	for index := 0; index < int(numberOfCases); index++ {
		inputString := readLine(reader)
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

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
