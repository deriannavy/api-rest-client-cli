package item

type Header struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ValueType string `json:"type"`
}

type Body struct {
	Mode string `json:"raw"`
}

type Request struct {
	Header []Header `json:"header"`
	Method string   `json:"method"`
	Body   Body     `json:"body"`
	// Url    string `json:"uri"`
}
