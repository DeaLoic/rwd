package destination

import (
	"encoding/csv"
	"os"

	"github.com/DeaLoic/rwd/word"
)

type CsvDestination struct {
	fileName string
}

func NewCsvDestination(fileName string) *CsvDestination {
	dest := CsvDestination{}
	if len(fileName) <= 0 {
		return nil
	}

	dest.fileName = fileName
	return &dest
}

func (dest CsvDestination) WriteWords(words []word.WordDescribed) error {
	f, err := os.Create(dest.fileName)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()
	err = writer.Write(word.WordDescribed{}.GetHeaders())
	if err != nil {
		return err
	}

	for _, wordSrc := range words {
		writer.Write(wordSrc.GetValues())
	}

	return nil
}
