package ui

import (
	"fmt"
	"strings"

	"github.com/deriannavy/api-rest-client-cli/styles"
)

type Tab struct {
	Name  string
	Badge string
}

func (t *Tab) SetBadge(badge string) {
	t.Badge = badge
}

func (t Tab) RenderTable(data []KeyValueObject) string {

	var b strings.Builder

	fmt.Fprintf(&b, "\n")

	for _, r := range data {
		fmt.Fprintf(&b, " %s %s %s\n", r.Key, styles.Cursor, r.Value)
	}

	return b.String()
}
func (t Tab) RenderBody(body string) string {
	var (
		b  strings.Builder
		bs = strings.Split(body, "\r\n")
	)

	fmt.Fprintf(&b, "\n")

	for _, l := range bs {
		fmt.Fprintf(&b, " %s\n", l)
	}

	return b.String()
}

func (t Tab) Render(item Item) string {
	var st string
	switch t.Name {
	case "Headers":
		st = t.RenderTable(item.Request.Header)
	case "Parameters":
		st = t.RenderTable(item.Request.Url.Query)
	case "Body":
		st = t.RenderBody(item.Request.Body.Raw)
	}

	return st
}
