package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/deriannavy/api-rest-client-cli/handler"
)

type Item struct {
	// > Json data
	Index   int
	Name    string  `json:"name"`
	Request Request `json:"request"`
}

type KeyValueObject struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ValueType string `json:"type"`
}

type Body struct {
	Mode string `json:"raw"`
}

type Url struct {
	Protocol string           `json:"protocol"`
	Host     []string         `json:"host"`
	Path     []string         `json:"path"`
	Query    []KeyValueObject `json:"query"`
}

type Request struct {
	Header []KeyValueObject `json:"header"`
	Method string           `json:"method"`
	Body   Body             `json:"body"`
	Url    Url              `json:"url"`
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
func (i Item) UrlFormat(ic ItemComplement) string {
	var (
		protocol = i.Request.Url.Protocol
		host     = strings.Join(i.Request.Url.Host, ".")
		path     = strings.Join(i.Request.Url.Path, "/")
		query    = ""
	)

	urlData := []string{protocol, "://", host, "/", path, query}
	url := strings.Join(urlData, "")

	return ic.Styles.UrlStyle.Render(url)
}

func (i Item) MethodFormat(ic ItemComplement, align string) string {
	var (
		method        string
		requestMethod = i.Request.Method
		strFormat     = "%*s"
		wlarge        = 4
	)
	if align == "right" {
		strFormat = "%-*s"
	}
	switch requestMethod {
	case "GET":
		method = ic.Styles.GetMethod.Render(fmt.Sprintf(strFormat, wlarge, "GET"))
	case "POST":
		method = ic.Styles.PostMethod.Render(fmt.Sprintf(strFormat, wlarge, "POST"))
	case "PUT":
		method = ic.Styles.PutMethod.Render(fmt.Sprintf(strFormat, wlarge, "PUT"))
	case "PATCH":
		method = ic.Styles.PatchMethod.Render(fmt.Sprintf(strFormat, wlarge, "PTCH"))
	case "DELETE":
		method = ic.Styles.DeleteMethod.Render(fmt.Sprintf(strFormat, wlarge, "DEL"))
	case "OPTIONS":
		method = ic.Styles.OptionsMethod.Render(fmt.Sprintf(strFormat, wlarge, "OPT"))
	case "HEAD":
		method = ic.Styles.HeadMethod.Render(fmt.Sprintf(strFormat, wlarge, "HEAD"))
	default:
		method = ic.Styles.UnknowMethod.Render(fmt.Sprintf(strFormat, wlarge, "UKNW"))
	}

	return method
}

func (i Item) TitleFormat(ic ItemComplement, isSelected bool) string {
	var (
		textwidth = ic.Size.Width() - ic.Styles.NormalTitle.GetPaddingLeft() - ic.Styles.NormalTitle.GetPaddingRight()
		name      = ansi.Truncate(i.Name, textwidth, ellipsis)
		title     = ic.Styles.NormalTitle.Render(name)
	)
	if isSelected {
		title = ic.Styles.SelectedTitle.Render(name)
	}
	return title
}

func (i Item) View(ic ItemComplement, isSelected bool) string {
	var (
		ccursor = " "
		method  = i.MethodFormat(ic, "left")
		title   = i.TitleFormat(ic, isSelected)
	)

	if isSelected {
		ccursor = cursor
	}
	return ic.Styles.SelectedCursor.Render(ccursor) + method + title
}
