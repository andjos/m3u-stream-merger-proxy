package m3u

import (
	"bufio"
	"bytes"
	"fmt"
)

func filterM3u(buffer bytes.Buffer) bytes.Buffer {

	scanner := bufio.NewScanner(&buffer)

	// Create a buffer to store the filtered lines
	var filteredBuffer bytes.Buffer

	for scanner.Scan() {

		line := scanner.Text()
		_, err := filteredBuffer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to buffer:", err)

		}

	}
	return filteredBuffer
}
