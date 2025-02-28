package ui

import (
	"fmt"

	"github.com/deriannavy/api-rest-client-cli/handler"
)

type Item struct {
	// > Json data
	Index   int
	Name    string  `json:"name"`
	Request Request `json:"request"`
}

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

type ItemComplement struct {
	Size   handler.SizeSpec
	Styles ItemStyle
}

// NewDefaultDelegate creates a new delegate with default styles.
func NewComplement(width, height int) ItemComplement {
	return ItemComplement{
		Size:   handler.NewSizeSpec(width, height),
		Styles: DefaultItemStyle(),
	}
}

func (i Item) MethodFormat(ic ItemComplement) string {
	var (
		method        string
		requestMethod = i.Request.Method
	)
	switch requestMethod {
	case "DELETE":
		method = "DEL"
	case "OPTIONS":
		method = "OPT"
	default:
		method = requestMethod
	}

	method = fmt.Sprintf("%5s", method)

	return method
}

func (i Item) View(ic ItemComplement, isSelected bool) string {
	if isSelected {
		return ic.Styles.SelectedTitle.Render(i.MethodFormat(ic) + " " + i.Name)
	}
	return ic.Styles.NormalTitle.Render(i.MethodFormat(ic) + " " + i.Name)
}

// func (ic ItemComplement) Render(item Item, index int) string {

// 	var (
// 		title = item.TitleFormat()
// 		s     = &ic.Styles
// 	)

// 	// textwidth := ic.Size.Width() - s.NormalTitle.GetPaddingLeft() - s.NormalTitle.GetPaddingRight()
// 	// title = ansi.Truncate(title, textwidth, ellipsis)
// 	if item.Index == index {
// 		return s.SelectedTitle.Render(title)
// 	}

// 	return s.NormalTitle.Render(title) //nolint: errcheck
// }
