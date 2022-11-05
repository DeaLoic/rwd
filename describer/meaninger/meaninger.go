package meaninger

type Meaninger interface {
	Meaning([]string) ([]string, error)
}
