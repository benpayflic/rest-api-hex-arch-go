package api

import (
	"github.com/benpayflic/rest-api-hex-arch-go/internal/ports"
	"github.com/benpayflic/rest-api-hex-arch-go/pkg/config"
)

// Application implements the APIPort interface
type Application struct {
	rest       ports.RestPort
	db         ports.DBPort
	catService CatService
	config     config.Config
}

// Return new instance of Application
func NewApplication(rest ports.RestPort, db ports.DBPort, catService CatService, config config.Config) *Application {
	return &Application{
		rest:       rest,
		db:         db,
		catService: catService,
		config:     config,
	}
}
