package readwrite

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TargetRead(targets string) ([]string, error) {
	file, err := os.Open(targets)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var url []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			url = append(url, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return url, nil

}
