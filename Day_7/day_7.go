/*
Objective
Today, we're learning about the Array data structure. Check out the Tutorial tab for learning materials and an instructional video!

Task
Given an array, , of  integers, print 's elements in reverse order as a single line of space-separated numbers.

Input Format

The first line contains an integer,  (the size of our array).
The second line contains  space-separated integers describing array 's elements.

Constraints

, where  is the  integer in the array.
Output Format

Print the elements of array  in reverse order as a single line of space-separated numbers.

Sample Input

4
1 4 3 2
Sample Output

2 3 4 1
*/
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

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	for i := 1; i <= n; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[n-i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		fmt.Print(arrItem, " ")
	}
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
