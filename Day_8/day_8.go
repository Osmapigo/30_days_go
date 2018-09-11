package main

import (
	"fmt"
)

func main() {
	var entries int
	fmt.Scan(&entries)
	directory := setDirectory(entries)
	searchNames(entries, directory)
}

func setDirectory(entries int) map[string]int {
	var directory = make(map[string]int)
	var name string
	var number int

	for i := 0; i < entries; i++ {
		fmt.Scan(&name, &number)
		directory[name] = number
	}
	return directory
}

func searchNames(entries int, directory map[string]int) {
	var name string
	for i := 0; i < entries; i++ {
		n, _ := fmt.Scanln(&name)
		if n == 0 {
			break
		}
		if directory[name] != 0 {
			fmt.Printf("%v=%v\n", name, directory[name])
		} else {
			fmt.Println("Not found")
		}
	}
}
