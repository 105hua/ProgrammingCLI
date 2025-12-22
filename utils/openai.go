package utils

import (
	"context"

	"github.com/carlmjohnson/requests"
)

const BASE_URL string = "https://openrouter.ai/api/v1"

type Model struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetModelsResponse struct {
	Data []Model `json:"data"`
}

func GetModels() GetModelsResponse {
	var models GetModelsResponse
	err := requests.
		URL(BASE_URL + "/models").
		ToJSON(&models).
		Fetch(context.Background())

	if err != nil {
		panic(err)
	}

	return models
}
