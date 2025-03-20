package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/deriannavy/api-rest-client-cli/handler"
	"github.com/deriannavy/api-rest-client-cli/styles"
)

type Item struct {
	// > Json data
	Index   int
	Name    string  `json:"name"`
	Request Request `json:"request"`
}

type KeyValueObject struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	ValueType   string `json:"type"`
	Description string `json:"description"`
}

type Body struct {
	Mode string `json:"mode"`
	Raw  string `json:"raw"`
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
	Styles styles.ItemStyle
}

// NewDefaultDelegate creates a new delegate with default styles.
func NewComplement(width, height int) ItemComplement {
	return ItemComplement{
		Size:   handler.NewSizeSpec(width, height),
		Styles: styles.DefaultItemStyle(),
	}
}

func (r Request) MethodAbreviated(align string) string {
	switch r.Method {
	case "GET", "POST", "PUT", "DELETE", "HEAD":
		return handler.Truncate(r.Method, "", align, 4)
	case "PATCH":
		return "PTCH"
	case "OPTIONS":
		return "OPTS"
	default:
		return "UKNW"
	}
}

func (i Item) UrlFormat() string {
	var (
		protocol = i.Request.Url.Protocol
		host     = strings.Join(i.Request.Url.Host, ".")
		path     = strings.Join(i.Request.Url.Path, "/")
		query    = ""
	)

	urlData := []string{protocol, "://", host, "/", path, query}
	return strings.Join(urlData, "")
}

func (i Item) UrlFormatStyle(ic ItemComplement) string {
	return ic.Styles.UrlStyle.Render(i.UrlFormat())
}

func (i Item) MethodFormat(align string, abreviated bool) string {
	if abreviated {
		return i.Request.MethodAbreviated(align)
	} else {
		return i.Request.Method
	}
}

func (i Item) MethodFormatStyle(ic ItemComplement, align string, abreviated bool) string {
	var (
		method = i.MethodFormat(align, abreviated)
		style  lipgloss.Style
	)
	switch method {
	case "GET", " GET":
		style = ic.Styles.GetMethod
	case "POST":
		style = ic.Styles.PostMethod
	case "PUT", " PUT":
		style = ic.Styles.PutMethod
	case "PTCH", "PATCH":
		style = ic.Styles.PatchMethod
	case "DELE", "DELETE", " DEL":
		style = ic.Styles.DeleteMethod
	case "OPTS", "OPTIONS":
		style = ic.Styles.OptionsMethod
	case "HEAD":
		style = ic.Styles.HeadMethod
	default:
		style = ic.Styles.UnknowMethod
	}
	return style.Render(method)
}

func (i Item) TitleFormat(ic ItemComplement, isSelected bool) string {
	var (
		textwidth = ic.Size.Width() - ic.Styles.NormalTitle.GetPaddingLeft() - ic.Styles.NormalTitle.GetPaddingRight()
		name      = handler.Truncate(i.Name, styles.Ellipsis, "left", textwidth)
		title     = ic.Styles.NormalTitle.Render(name)
	)
	if isSelected {
		title = ic.Styles.SelectedTitle.Render(name)
	}
	return title
}

func (i Item) View(ic ItemComplement, isSelected bool) string {
	var (
		cursor = " "
		method = i.MethodFormatStyle(ic, "right", true)
		title  = i.TitleFormat(ic, isSelected)
	)

	if isSelected {
		cursor = styles.Cursor
	}
	return ic.Styles.SelectedCursor.Render(cursor) + method + title
}
