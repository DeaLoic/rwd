package describer

import (
	"errors"

	"github.com/DeaLoic/rwd/word"
)

type WordDescriberType uint

const (
	General WordDescriberType = iota
)

func NewWordDescriber(describerType WordDescriberType) (word.WordDescriber, error) {
	switch describerType {
	case General:
		return NewGeneralWordDescriber()
	default:
		return nil, errors.New("InvalidType")
	}
}
