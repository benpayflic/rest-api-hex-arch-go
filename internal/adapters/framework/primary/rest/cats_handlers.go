package rest

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/benpayflic/rest-api-hex-arch-go/internal/adapters/framework/primary/rest/dto"
	"github.com/benpayflic/rest-api-hex-arch-go/internal/application/domain/cats"
)

// TODO: Implement other crud operations

func createCatHandler(restAdapter Adapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		b, err := io.ReadAll(r.Body)
		if err != nil {
			httpErrorResponse(400, "Failed to read response body", w, err)
			return
		}
		c, err := cats.UnmarshalCat(b)
		if err != nil {
			httpErrorResponse(400, "Failed to unmarshal response body", w, err)
			return
		}
		err = restAdapter.api.CreateCat(&c)
		if err != nil {
			httpErrorResponse(400, "Failed to create new cat", w, err)
			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(dto.CreateCatDTO{
			Message: "Created a new cat",
			Cat:     c,
		})

	}
}
