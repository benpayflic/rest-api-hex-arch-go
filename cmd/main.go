package main

import (
	"log"

	c "github.com/benpayflic/rest-api-hex-arch-go/pkg/config"
)

func main() {
	config, err := c.LoadConfig("./pkg/config/")
	if err != nil {
		log.Fatalln("Failed to load application configurations : ", err)
	}
	println(config.AppHost, config.AppPort)

}
