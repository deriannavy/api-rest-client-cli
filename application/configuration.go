package application

import (
	"encoding/json"
	"io"
	"os"

	"github.com/deriannavy/api-rest-client-cli/ui"
)

// https://schema.postman.com/json/collection/v2.1.0/collection.json
type Configuration struct {
	Information Information `json:"information"`
	Items       []ui.Item   `json:"item"`
}

type Information struct {
	Name string `json:"name"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadConfiguration(c *Configuration) {

	// Abrir archivo JSON
	file, err := os.Open("schema.json")
	check(err)

	defer file.Close()

	// Read file content
	data, err := io.ReadAll(file)
	check(err)

	if err := json.Unmarshal(data, &c); err != nil {
		check(err)
	}

	for i := range c.Items {
		c.Items[i].Index = i
	}

}
