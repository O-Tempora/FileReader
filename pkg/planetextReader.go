package readers

import (
	"bufio"
	"fmt"
	"os"
)

func ReadPlane(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed reading file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Println("An error occured while reading file: ", err)
		return
	}
}
