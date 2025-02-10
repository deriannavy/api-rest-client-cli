package components

import (
	"github.com/deriannavy/api-rest-client-cli/application"
	"github.com/deriannavy/api-rest-client-cli/panel"
)

func NewPanel(itemsConfig []application.Config) panel.Model {

	p := panel.New(itemsConfig)

	return p
}
