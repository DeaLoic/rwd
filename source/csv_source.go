package source

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/DeaLoic/rwd/word"
)

type CsvSource struct {
	fileName string
}

func NewCsvSource(fileName string) *CsvSource {
	dest := CsvSource{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer file.Close()

	dest.fileName = fileName
	return &dest
}

func (src CsvSource) GetWords() []word.WordSource {
	f, err := os.Open(src.fileName)
	if err != nil {
		return nil
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.Comma = ','
	records, err := reader.ReadAll()

	if err != nil {
		return nil
	}

	words := make([]word.WordSource, len(records))
	for i, record := range records {

		pageStr, _ := strconv.ParseUint(record[0], 10, 32)
		wordSource := word.WordSource{
			Page: uint32(pageStr),
			Word: record[1],
		}

		fmt.Printf("%d %s\n", wordSource.Page, wordSource.Word)
		words[i] = wordSource
	}

	return words
}
