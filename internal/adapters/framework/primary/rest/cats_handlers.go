package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/benpayflic/rest-api-hex-arch-go/internal/adapters/framework/primary/rest/dto"
	"github.com/benpayflic/rest-api-hex-arch-go/internal/application/domain/cats"
	"github.com/gorilla/mux"
)

func createCatHandler(restAdapter Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		b, err := io.ReadAll(r.Body)
		if err != nil {
			httpErrorResponse(422, "Failed to read response body", w, err)
			return
		}
		c, err := cats.UnmarshalCat(b)
		if err != nil {
			httpErrorResponse(422, "Failed to unmarshal response body", w, err)
			return
		}
		err = restAdapter.api.CreateCat(&c)
		if err != nil {
			httpErrorResponse(400, "Failed to create new cat", w, err)
			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(dto.CreateCatDTO{
			Message: fmt.Sprintf("Created a new cat with name: %v ", c.Name),
		})

	}
}

func getCatHandler(restAdapter Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		if _, ok := vars["catName"]; !ok {
			httpErrorResponse(422, "Invalid request", w, errors.New("please provide cat name as a query string parameter"))
			return
		}

		cat, err := restAdapter.api.RetrieveCat(vars["catName"])
		if err != nil {
			httpErrorResponse(400, "No cat found with provided name", w, err)
			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(dto.GetCatDTO{
			Message: "Found cat",
			Cat:     *cat,
		})

	}
}

func updateCatHandler(restAdapter Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		b, err := io.ReadAll(r.Body)
		if err != nil {
			httpErrorResponse(422, "Failed to read response body", w, err)
			return
		}
		c, err := cats.UnmarshalCat(b)
		if err != nil {
			httpErrorResponse(422, "Failed to unmarshal response body", w, err)
			return
		}
		err = restAdapter.api.UpdateCat(&c)
		if err != nil {
			httpErrorResponse(400, "Failed to update cat", w, err)
			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(dto.UpdateCatDTO{
			Message: fmt.Sprintf("Updated cat with name: %v", c.Name),
		})

	}
}

func deleteCatHandler(restAdapter Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		if _, ok := vars["catName"]; !ok {
			httpErrorResponse(400, "Invalid request", w, errors.New("please provide cat name as a query string parameter"))
			return
		}

		err := restAdapter.api.DeleteCat(vars["catName"])
		if err != nil {
			httpErrorResponse(400, "Failed to delete cat", w, err)
			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(dto.DeleteCatDTO{
			Message: fmt.Sprintf("Deleted cat with name: %v", vars["catName"]),
		})

	}
}
