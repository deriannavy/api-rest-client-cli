package application

import (
	"encoding/json"
	"io"
	"os"

	"github.com/deriannavy/api-rest-client-cli/item"
)

// https://schema.postman.com/json/collection/v2.1.0/collection.json
type Configuration struct {
	Information Information `json:"information"`
	Items       []item.Item `json:"item"`
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

}
