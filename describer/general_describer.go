package describer

import (
	"fmt"

	"github.com/DeaLoic/rwd/describer/meaninger"
	"github.com/DeaLoic/rwd/describer/translator"
	"github.com/DeaLoic/rwd/word"
)

type GeneralWordDescriber struct {
	translator translator.Translator
	meaning    meaninger.Meaninger
}

func NewGeneralWordDescriber() (GeneralWordDescriber, error) {
	return GeneralWordDescriber{}, nil
}

func NewGeneralWordDescriberBuild(translator translator.Translator, meaninger meaninger.Meaninger) GeneralWordDescriber {
	return GeneralWordDescriber{translator, meaninger}
}

func (descr GeneralWordDescriber) Describe(words []word.WordSource) ([]word.WordDescribed, error) {
	described := make([]word.WordDescribed, len(words))
	simpleWords := make([]string, len(words))
	for i, wordSrc := range words {
		simpleWords[i] = wordSrc.Word
		described[i] = word.WordDescribed{Page: wordSrc.Page, Word: wordSrc.Word, Translation: "N/A", Meaning: "N/A"}
	}
	translations, err := descr.translator.Translate(simpleWords)
	if err != nil {
		fmt.Printf("Error in Translation: %s", err.Error())
	}

	meanings, err := descr.meaning.Meaning(simpleWords)
	if err != nil {
		fmt.Printf("Error in Meaning: %s", err.Error())
	}

	for i, _ := range described {
		described[i].Meaning = meanings[i]
		described[i].Translation = translations[i]
	}

	return described, nil
}
