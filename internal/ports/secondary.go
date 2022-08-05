package ports

import (
	c "github.com/benpayflic/rest-api-hex-arch-go/internal/application/domain/cats"
)

// Driven ports

type RestPort interface {
	GetCatFact() (string, error)
}

type DBPort interface {
	Connect()
	CreateCat(*c.Cat) error
	RetrieveCat(string) (*c.Cat, error)
	UpdateCat(*c.Cat) error
	DeleteCat(string) error
}
