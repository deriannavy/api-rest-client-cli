package item

type Item struct {
	Name    string  `json:"name"`
	Request Request `json:"request"`
}
