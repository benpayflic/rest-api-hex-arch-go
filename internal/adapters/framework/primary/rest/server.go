package rest

import (
	"log"
	"net/http"

	"github.com/benpayflic/rest-api-hex-arch-go/internal/ports"
	"github.com/benpayflic/rest-api-hex-arch-go/pkg/config"
	"github.com/gorilla/mux"
)

// Rest port adpater
type Adapter struct {
	api    ports.APIPort
	config *config.Config
}

func NewAdapter(api ports.APIPort, config *config.Config) *Adapter {
	return &Adapter{api: api, config: config}
}

func registerHandlers(restAdapter Adapter, r *mux.Router) {

	r.HandleFunc("/api/v1/cats", createCatHandler(restAdapter)).Methods("POST")
}

func (restAdapter Adapter) Start() {
	router := mux.NewRouter().StrictSlash(true)
	// middlewares
	router.Use(responseManagerMiddleware)
	router.Use(loggingMiddleware)
	// handlers
	registerHandlers(restAdapter, router)

	log.Printf("Server running on host %v and listenting on port %v . ctrl + c to quit!",
		restAdapter.config.AppHost, restAdapter.config.AppPort)

	if err := http.ListenAndServe(restAdapter.config.AppPort, router); err != nil {
		log.Fatalf("failed to serve rest server over port %v: %v", restAdapter.config.AppPort, err)
	}
}
