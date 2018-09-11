package readerpkg

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// ReadInt get the integer from the commad line
func ReadInt() int64 {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	Number, _ := strconv.ParseInt(readLine(reader), 10, 32)
	return Number
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
