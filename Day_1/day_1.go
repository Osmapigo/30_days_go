package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	scanner := bufio.NewScanner(os.Stdin)
	scanner = scanner
	scanner2 := bufio.NewReader(os.Stdin)

	var i uint64 = 4
	var d = 4.0
	var s = "HackerRank "

	// Declare second integer, double, and String variables.
	var j uint64
	var f float64
	// Read and save an integer, double, and String to your variables.
	// Note: If you have trouble reading the entire string, please go back and review the Tutorial closely.
	fmt.Scan(&j)
	fmt.Scan(&f)
	stringValue, _ := scanner2.ReadString('\n')
	// Print the sum of both integer variables on a new line.
	fmt.Println(i + j)
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f", d+f)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println("\n" + s + stringValue)
}
