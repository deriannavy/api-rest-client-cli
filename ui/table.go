package ui

import "github.com/deriannavy/api-rest-client-cli/handler"

type Table struct {
	// Headers & Rows
	headers []string
	rows    []string
	// Window Size
	Size handler.SizeSpec
	// Cells size
	MaxCellHeight int
	MaxCellWidth  int
}
