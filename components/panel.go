package components

import (
	"github.com/deriannavy/api-rest-client-cli/application"
	"github.com/deriannavy/api-rest-client-cli/panel"
)

func NewPanel(Config application.Config) panel.Model {

	p := panel.New(Config)

	return p
}
