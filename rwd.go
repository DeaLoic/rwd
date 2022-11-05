package main

import (
	"fmt"

	"github.com/DeaLoic/rwd/describer"
	meaninger "github.com/DeaLoic/rwd/describer/meaninger"
	"github.com/DeaLoic/rwd/describer/translator"
	"github.com/DeaLoic/rwd/destination"
	"github.com/DeaLoic/rwd/source"
)

const IamToken = ""
const FolderId = ""

func main() {
	describer := describer.NewGeneralWordDescriberBuild(translator.NewYandexTranslator(&translator.YandexTanslatorConfig{
		IamToken:       IamToken,
		FolderId:       FolderId,
		TargetLanguage: "ru"}),
		meaninger.NewFreeDictionaryApiMeaninger(&meaninger.FreeDictionaryApiMeaningerConfig{}))

	csvSource := source.NewCsvSource("words.csv")
	csvDest := destination.NewCsvDestination("result.csv")

	words := csvSource.GetWords()
	println(words)
	wordsDescribed, err := describer.Describe(words)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	csvDest.WriteWords(wordsDescribed)

	fmt.Println("123")
}
