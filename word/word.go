package word

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
	Describe(ws WordSource) (WordDescribed, error)
}

func NewWordDescribed() WordDescribed {
	return WordDescribed{}
}

func NewWordSource() WordSource {
	return WordSource{}
}

func (ws WordSource) Describe(describer WordDescriber) (*WordDescribed, error) {
	wd, err := describer.Describe(ws)
	if err != nil {
		return nil, err
	}
	return &wd, nil
}
