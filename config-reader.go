package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadConfig(fileName string) {

	type Configuration struct {
		Users  []string
		Groups []string
	}

	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(configuration.Users)
	fmt.Println("Hello world!")
}
