package ui

import (
	"fmt"
	"strings"

	"github.com/deriannavy/api-rest-client-cli/handler"
	"github.com/deriannavy/api-rest-client-cli/styles"
)

type Table struct {
	// Styles
	Styles styles.TableStyle
	// Headers & Rows
	headers []string
	rows    [][]string
	// Window Size
	Size handler.SizeSpec
	// Cells size
	MaxCellHeight []int
	MaxCellWidth  []int
}

func NewTable() Table {

	t := Table{
		// Styles
		Styles: styles.DefaultTableStyle(),
		// Window Size
		Size: handler.NewSizeSpec(0, 0),
		// Cells size
		MaxCellHeight: []int{},
		MaxCellWidth:  []int{},
	}

	return t
}

func (t *Table) AddHeaders(headers ...string) {
	for _, header := range headers {
		t.MaxCellWidth = append(t.MaxCellWidth, len(header))
		t.headers = append(t.headers, header)
	}
}

func (t *Table) AddRow(row []string) {
	for i, cell := range row {
		t.MaxCellWidth[i] = max(len(cell), t.MaxCellWidth[i])
	}
	t.rows = append(t.rows, row)
}

func (t Table) RenderHeaders() string {
	headers := ""

	for i, header := range t.headers {
		size := t.MaxCellWidth[i]
		headers += handler.FillCenter(header, size)
	}

	return t.Styles.HeaderStyle.Render(headers)
}

func (t Table) RenderRows() string {
	rowsString := ""
	for i, row := range t.rows {
		rString := ""

		for ir, r := range row {
			size := t.MaxCellWidth[ir]
			rString += handler.FillCenter(r, size)
		}

		if i%2 == 0 {
			rowsString += t.Styles.RowOddStyle.Render(rString)
		} else {
			rowsString += t.Styles.RowEvenStyle.Render(rString)
		}
		rowsString += "\n"
	}
	return rowsString
}

func (t Table) View() string {

	var b strings.Builder

	fmt.Fprintf(&b, "%s\n", t.RenderHeaders())

	fmt.Fprintf(&b, "%s", t.RenderRows())

	return b.String()
}
