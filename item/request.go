package item

type Header struct {
	key       string `json:"key"`
	value     string `json:"value"`
	valueType string `json:"type"`
}

type Body struct {
	mode string `json:"raw"`
}

type Request struct {
	Header []Header `json:"header"`
	Method string   `json:"method"`
	Body   Body     `json:"body"`
	// Url    string `json:"uri"`
}
