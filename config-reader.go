package mucupa

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Url  string
	Port string
}

func ReadConfig(fileName string) Configuration {

	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)

	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		// return an error
	}

	return configuration
}
