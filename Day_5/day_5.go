package main

import (
	"fmt"
)

// func main() {

// 	n := readerpkg.ReadInt()
// 	for index := 1; index <= 10; index++ {
// 		fmt.Printf("%v x %v = %v\n", n, index, int64(index))
// 	}
// }

func main() {

	var n int
	fmt.Scan(&n)
	for index := 1; index <= 10; index++ {
		fmt.Printf("%v x %v = %v\n", n, index, int64(index))
	}
}
