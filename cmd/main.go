package main

import (
	"log"

	"github.com/benpayflic/rest-api-hex-arch-go/internal/adapters/framework/primary/rest"
	"github.com/benpayflic/rest-api-hex-arch-go/internal/application/api"
	"github.com/benpayflic/rest-api-hex-arch-go/internal/application/domain/cats"
	c "github.com/benpayflic/rest-api-hex-arch-go/pkg/config"
)

func main() {
	config, err := c.LoadConfig("./pkg/config/")
	if err != nil {
		log.Fatalln("Failed to load application configurations : ", err)
	}

	// init core domain logic services
	catService := cats.NewCatService()

	applicationAPI := api.NewApplication(catService, config)

	// init left side (primary) adapters
	restAdapter := rest.NewAdapter(applicationAPI, &config)
	// start rest server
	restAdapter.Start()
}
