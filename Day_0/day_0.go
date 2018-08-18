package day0

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello, World.")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
