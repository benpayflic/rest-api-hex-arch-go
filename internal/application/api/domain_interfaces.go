package api

import "github.com/benpayflic/rest-api-hex-arch-go/internal/application/domain/cats"

type CatService interface {
	CalculateHumanYears(c *cats.Cat)
}
