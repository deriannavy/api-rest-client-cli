package application

import (
	"fmt"
)

type Server struct {
	Protocol string `yaml:"protocol"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}

type Request struct {
	Method string `yaml:"method"`
	Path   string `yaml:"path"`
}

type Config struct {
	ID      string
	Name    string `yaml:"name"`
	Server  Server
	Request Request
}

func (c Config) toItem() Item {
	return NewItem(c.ID, c.Name, c.getUri())
}

func (c Config) getUri() string {

	p := ""
	if c.Server.Port != "" {
		p = fmt.Sprintf(":%s", c.Server.Port)
	}

	uri := fmt.Sprintf("%s%s%s%s",
		c.Server.Protocol,
		c.Server.Host,
		c.Request.Path,
		p,
	)

	return uri
}
