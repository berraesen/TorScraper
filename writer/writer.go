package writer

import (
	"os"
	"time"
)

func SaveHTML(data []byte) (string, error) {

	timestamp := time.Now().Format("20060102_150405")

	fileName := "screenshots/" + timestamp + ".html"

	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
