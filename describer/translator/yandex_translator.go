package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const yandexUrl string = "https://translate.api.cloud.yandex.net/translate/v2/translate"

type YandexTanslatorConfig struct {
	IamToken       string
	FolderId       string
	TargetLanguage string
}

type requestBody struct {
	FolderId           string   `json:"folderId"`
	Texts              []string `json:"texts"`
	TargetLanguageCode string   `json:"targetLanguageCode"`
}

type translatedWord struct {
	Text                 string `json:"text"`
	DetectedLanguageCode string `json:"detectedLanguageCode"`
}
type responseBody struct {
	Translations []translatedWord `json:"translations"`
}

type YandexTanslator struct {
	config  *YandexTanslatorConfig
	request *http.Request
	body    *requestBody
}

func NewYandexTranslator(config *YandexTanslatorConfig) YandexTanslator {
	yt := YandexTanslator{config, nil, nil}
	yt.formRequest()
	yt.fillBody()

	return yt
}

func (yt YandexTanslator) Translate(words []string) ([]string, error) {
	hc := http.Client{}
	yt.fillWords(words)
	marhalizedStruct, _ := json.Marshal(yt.body)
	yt.request.Body = io.NopCloser(bytes.NewBuffer(marhalizedStruct))

	resp, err := hc.Do(yt.request)
	if err != nil {
		fmt.Printf("Error in tr resp %e", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data responseBody
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Error in tr Unmarshal %e", err)
		return nil, err
	}

	translated := make([]string, len(data.Translations))
	for i, tw := range data.Translations {
		translated[i] = tw.Text
	}

	return translated, nil
}

func (yt *YandexTanslator) fillWords(words []string) {
	yt.body.Texts = words
}

func (yt *YandexTanslator) formRequest() {
	req, _ := http.NewRequest("POST", yandexUrl, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+yt.config.IamToken)
	yt.request = req
}

func (yt *YandexTanslator) fillBody() {
	yt.body = &requestBody{FolderId: yt.config.FolderId, Texts: nil, TargetLanguageCode: yt.config.TargetLanguage}
}
