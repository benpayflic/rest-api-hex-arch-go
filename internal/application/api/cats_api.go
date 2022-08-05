package api

import "github.com/benpayflic/rest-api-hex-arch-go/internal/application/domain/cats"

func (api Application) CreateCat(c *cats.Cat) error {
	// calculate cat age in human years
	api.catService.CalculateHumanYears(c)

	// get cat fact, third party rest call.
	fact, err := api.rest.GetCatFact()
	if err != nil {
		return err
	}
	// add cat fact to cat object
	c.CatFact = fact

	// push cat to database.
	err = api.db.CreateCat(c)
	if err != nil {
		return err
	}

	return nil
}

func (api Application) RetrieveCat(catName string) (*cats.Cat, error) {

	// get cat from database.
	cat, err := api.db.RetrieveCat(catName)
	if err != nil {
		return nil, err
	}
	return cat, nil
}

func (api Application) UpdateCat(c *cats.Cat) error {
	// calculate cat age in human years
	api.catService.CalculateHumanYears(c)

	// get cat fact, third party rest call.
	fact, err := api.rest.GetCatFact()
	if err != nil {
		return err
	}
	// add cat fact to cat object
	c.CatFact = fact

	// push cat to database.
	err = api.db.UpdateCat(c)
	if err != nil {
		return err
	}
	return nil
}

func (api Application) DeleteCat(catName string) error {

	// delete cat from database.
	err := api.db.DeleteCat(catName)
	if err != nil {
		return err
	}
	return nil
}
