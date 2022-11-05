package source

import "github.com/DeaLoic/rwd/word"

type WordSource interface {
	GetWords() []word.WordSource
}
