package ui

import (
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
	MaxCellHeight int
	MaxCellWidth  int
}

func NewTable(headers []string, rows [][]string) Table {

	return Table{
		// Styles
		Styles: styles.DefaultTableStyle(),
		// Headers & Rows
		headers: headers,
		rows:    rows,
		// Window Size
		Size: handler.NewSizeSpec(0, 0),
		// Cells size
		MaxCellHeight: 0,
		MaxCellWidth:  0,
	}
}

func (t Table) View() string {
	return ""
}
