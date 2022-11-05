package destination

import "github.com/DeaLoic/rwd/word"

type WordDestinations interface {
	WriteWords(word.WordDescribed)
}
