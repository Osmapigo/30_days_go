package main

import (
	"30_days_go/readerpkg"
	"fmt"
)

func main() {
	message := "Not Weird"
	Number := readerpkg.ReadInt()
	if Number%2 != 0 || (Number > 5 && Number <= 20) {
		message = "Weird"
	}
	fmt.Println(message)
}
