package word

import "strconv"

type WordSource struct {
	Page uint32
	Word string
}

type WordDescribed struct {
	Page        uint32
	Word        string
	Translation string
	Meaning     string
}

type WordDescriber interface {
	Describe(ws []WordSource) ([]WordDescribed, error)
}

func NewWordDescribed() WordDescribed {
	return WordDescribed{}
}

func NewWordSource() WordSource {
	return WordSource{}
}

func (ws WordSource) Describe(describer WordDescriber) (*WordDescribed, error) {
	arr := make([]WordSource, 1)
	arr[0] = ws
	wd, err := describer.Describe(arr)
	if err != nil {
		return nil, err
	}
	return &wd[0], nil
}

func (wd WordDescribed) GetHeaders() []string {
	return []string{"Page", "Word", "Translation", "Meaning"}
}

func (wd WordDescribed) GetValues() []string {
	return []string{strconv.Itoa(int(wd.Page)), wd.Word, wd.Translation, wd.Meaning}
}
