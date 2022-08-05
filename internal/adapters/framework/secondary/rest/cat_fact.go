package rest

import (
	"errors"
	"io/ioutil"
	"net/http"

	c "github.com/benpayflic/rest-api-hex-arch-go/internal/application/domain/cats"
	"github.com/benpayflic/rest-api-hex-arch-go/pkg/globals"
)

func (ha Adapter) GetCatFact() (string, error) {

	req, err := http.NewRequest("GET", globals.GetCatRoute, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode > 200 {
		return "", errors.New("Could not get cat fact. Unknwon error: " + resp.Status)
	}

	catFact, err := c.UnmarshalCat(body)
	if err != nil {
		return "", err
	}
	return catFact.CatFact, nil
}
