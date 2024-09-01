package reader

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type AlbumQuery struct {
	Title  string
	Artist string
}

const (
	BUFFER_CAP = 1000
)

func GetRecords(inputPath string) ([]AlbumQuery, error) {

	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	buf := make([]byte, BUFFER_CAP)

	scanner.Buffer(buf, BUFFER_CAP)

	records := make([]AlbumQuery, 0) // 0 initial size, resizes

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "-")

		if len(parts) != 2 {
			return nil, errors.New("Invalid album line in file: " + line)
		}

		title := parts[0]
		artist := parts[1]

		records = append(records, AlbumQuery{
			Title:  title,
			Artist: artist,
		})
	}

	if len(records) == 0 {
		fmt.Println("No records found in file, exiting")
		os.Exit(1)
	}

	return records, nil

}
