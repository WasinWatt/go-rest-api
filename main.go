package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/wasinwatt/go-rest-api/app"
	"github.com/wasinwatt/go-rest-api/config"
)

func main() {
	a := app.App{}
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := config.Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config.SQLURL)

	a.Initialize(config.SQLURL)
	a.Run(":8080")
}
