package translator

type Translator interface {
	Translate([]string) ([]string, error)
}
