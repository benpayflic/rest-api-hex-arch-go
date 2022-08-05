package main

import (
	"log"

	leftRest "github.com/benpayflic/rest-api-hex-arch-go/internal/adapters/framework/primary/rest"
	rightDB "github.com/benpayflic/rest-api-hex-arch-go/internal/adapters/framework/secondary/database"
	rightRest "github.com/benpayflic/rest-api-hex-arch-go/internal/adapters/framework/secondary/rest"
	"github.com/benpayflic/rest-api-hex-arch-go/internal/application/api"
	"github.com/benpayflic/rest-api-hex-arch-go/internal/application/domain/cats"
	c "github.com/benpayflic/rest-api-hex-arch-go/pkg/config"
)

func main() {
	// config, err := c.LoadConfig("./pkg/config/")
	config, err := c.LoadConfig("/Users/bengoodenough/Documents/Portfolio-Projects/rest-api-hex-arch-go/pkg/config/")
	if err != nil {
		log.Fatalln("Failed to load application configurations : ", err)
	}

	restDrivenAdapter, err := rightRest.NewAdapter()
	if err != nil {
		log.Fatalf("failed to initialize rest driven adapter: %v", err)
	}
	dbDrivenAdapter, err := rightDB.NewAdapter(&config)
	if err != nil {
		log.Fatalf("failed to initialize db driven adapter: %v", err)
	}
	dbDrivenAdapter.Connect()

	// init core domain logic services
	catService := cats.NewCatService()

	applicationAPI := api.NewApplication(restDrivenAdapter, dbDrivenAdapter, catService, config)

	// init left side (primary) adapters
	restAdapter := leftRest.NewAdapter(applicationAPI, &config)
	// start rest server
	restAdapter.Start()
}
