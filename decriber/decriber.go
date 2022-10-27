package decriber

import (
	"fmt"

	"github.com/DeaLoic/rwd/word"
)

type WordDescriber struct {
	GetTranslation func(word string) (string, error)
	GetMeaning     func(word string) (string, error)
}

func (descr *WordDescriber) Describe(wordSource word.WordSource) word.WordDescribed {
	described := word.WordDescribed{wordSource.Page, wordSource.Word, "N/A", "N/A"}

	description, err := descr.GetTranslation(wordSource.Word)
	if err != nil {
		fmt.Printf("Error in Translation of \"%s\": %s", described.Word, err.Error())
	} else {
		described.Translation = description
	}

	meaning, err := descr.GetMeaning(wordSource.Word)
	if err != nil {
		fmt.Printf("Error in Meaning of \"%s\": %s", described.Word, err.Error())
	} else {
		described.Meaning = meaning
	}

	return described
}
