package ports

import "github.com/benpayflic/rest-api-hex-arch-go/internal/application/domain/cats"

// Driving ports

// APIPort is the technology neutral
// port for driving adapters
type APIPort interface {
	CreateCat(c *cats.Cat) error
	RetrieveCat(string) (*cats.Cat, error)
	UpdateCat(*cats.Cat) error
	DeleteCat(string) error
}
