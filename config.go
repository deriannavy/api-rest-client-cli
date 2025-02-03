package main

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
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

type ListConfiguration struct {
	Configurations []Config
}

func (c Config) toItem() item {

	name := fmt.Sprintf("%s %s", c.Request.Method, c.Name)
	// method := c.Request.Method

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

	// c.ID, name, method, uri
	return item{
		title: name,
		description: uri,
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadConfig(lcfg *ListConfiguration) {

	folderName := "config"

	c, err := os.ReadDir(folderName)
	check(err)

	for index, entry := range c {

		fileName := entry.Name()

		if entry.IsDir() {
			continue
		}

		filePath := fmt.Sprintf("%s/%s", folderName, fileName)

		f, err := os.Open(filePath)
		check(err)

		cfg := Config{ID: strconv.Itoa(index)}

		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&cfg)
		check(err)

		lcfg.Configurations = append(lcfg.Configurations, cfg)

		f.Close()
	}
}
