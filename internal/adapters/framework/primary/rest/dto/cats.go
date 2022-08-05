package dto

import "github.com/benpayflic/rest-api-hex-arch-go/internal/application/domain/cats"

type CreateCatDTO struct {
	Message string `json:"message"`
}

type GetCatDTO struct {
	Message string   `json:"message"`
	Cat     cats.Cat `json:"cat"`
}

type UpdateCatDTO struct {
	Message string `json:"message"`
}

type DeleteCatDTO struct {
	Message string `json:"message"`
}
