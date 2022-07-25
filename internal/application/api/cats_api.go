package api

import "github.com/benpayflic/rest-api-hex-arch-go/internal/application/domain/cats"

func (api Application) CreateCat(c *cats.Cat) error {
	// calculate cat age in human years
	api.catService.CalculateHumanYears(c)

	// get cat funny fact, third part rest call.

	// push cat to database.

	return nil
}
