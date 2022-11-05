package meaninger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const wordApi string = "https://api.dictionaryapi.dev/api/v2/entries/en/"

type FreeDictionaryApiMeaningerConfig struct {
}

type definition struct {
	Definition string `json:"definition"`
}

type meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []definition `json:"definitions"`
}

type definitionedWord struct {
	Word       string    `json:"word"`
	SourceUrls []string  `json:"sourceUrls"`
	Meanings   []meaning `json:"meanings"`
}

type responseBody struct {
	words []definitionedWord
}

type FreeDictionaryApiMeaninger struct {
	config *FreeDictionaryApiMeaningerConfig
}

func NewFreeDictionaryApiMeaninger(config *FreeDictionaryApiMeaningerConfig) FreeDictionaryApiMeaninger {
	meaninger := FreeDictionaryApiMeaninger{config}

	return meaninger
}

func (meaninger FreeDictionaryApiMeaninger) Meaning(words []string) ([]string, error) {
	hc := http.Client{}
	definitioned := make([]string, len(words))
	for i, word := range words {
		definitioned[i] = "N/A"
		request := meaninger.formRequest(word)
		resp, err := hc.Do(request)
		if err != nil {
			fmt.Printf("Error in ma resp %e", err)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		data := responseBody{}
		err = json.Unmarshal(body, &data.words)
		if err != nil {
			fmt.Printf("\nError in me Unmarshal %s\n\n", err)
			continue
		}

		if len(data.words) > 0 && len(data.words[0].Meanings) > 0 && len(data.words[0].Meanings[0].Definitions) > 0 {
			definitioned[i] = data.words[0].Meanings[0].Definitions[0].Definition
		}
	}

	return definitioned, nil
}

func (meaninger *FreeDictionaryApiMeaninger) formRequest(word string) *http.Request {
	req, _ := http.NewRequest("GET", wordApi+word, nil)
	req.Header.Add("Content-Type", "application/json")
	return req
}
