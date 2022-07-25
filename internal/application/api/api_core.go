package api

import (
	"github.com/benpayflic/rest-api-hex-arch-go/pkg/config"
)

// Application implements the APIPort interface
type Application struct {
	catService CatService
	config     config.Config
}

// Return new instance of Application
func NewApplication(catService CatService, config config.Config) *Application {
	return &Application{
		catService: catService,
		config:     config,
	}
}
