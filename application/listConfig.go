package application

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/list"
	"gopkg.in/yaml.v2"
	// "gopkg.in/yaml.v2"
)

type ListConfiguration struct {
	Configurations []Config
}

func (lc ListConfiguration) GetConfigByIndex(index int) Config {
	return lc.Configurations[index]
}

func (lc ListConfiguration) GetItemList() []list.Item {

	var items []list.Item
	for _, lci := range lc.Configurations {
		items = append(items, lci.toItem())
	}

	return items
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadConfig(lcfg *ListConfiguration) {

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
