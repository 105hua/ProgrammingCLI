package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/carlmjohnson/requests"
)

const BASE_URL string = "https://openrouter.ai/api/v1"
const DEFAULT_MODEL string = "minimax/minimax-m2.1" // Make sure this is always a relatively cheap model at the least, don't want to burn through peoples moneys ~~ William, 2025

type Model struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetModelsResponse struct {
	Data []Model `json:"data"`
}

type Conversation struct {
	Messages []Message
}

type Message struct {
	Role      string `json:"role"`
	Content   string `json:"content"`
	Reasoning string `json:"reasoning,omitempty"`
}

type LogProbs any

type GetCompletionResponse struct {
	Id                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Provider          string   `json:"provider"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Usage             Usage    `json:"usage"`
}

type Choice struct {
	Message            Message  `json:"message"`
	Index              int      `json:"index"`
	Logprobs           LogProbs `json:"logprobs"`
	FinishReason       string   `json:"finish_reason"`
	NativeFinishReason string   `json:"native_finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"error"`
}

func GetModels(apiKey string) GetModelsResponse {
	var models GetModelsResponse
	err := requests.
		URL(BASE_URL+"/models").
		Header("Authorization", "Bearer "+apiKey).
		ToJSON(&models).
		Fetch(context.Background())

	if err != nil {
		panic(err)
	}

	return models
}

func GetCompletion(userMessage string, modelId *string, apiKey string) Message {
	// Handle modelId
	model := DEFAULT_MODEL
	if modelId != nil {
		model = *modelId
	}

	// Add user message to conversation
	CurrentConversation.Messages = append(CurrentConversation.Messages, Message{
		Role:    "user",
		Content: userMessage,
	})

	var rawResponse string
	err := requests.
		URL(BASE_URL+"/chat/completions").
		Header("Authorization", "Bearer "+apiKey).
		BodyJSON(map[string]interface{}{
			"model":    model,
			"messages": CurrentConversation.Messages,
		}).
		ToString(&rawResponse).
		Fetch(context.Background())

	if err != nil {
		panic(err)
	}

	// Check for error response
	var errorResp ErrorResponse
	if json.Unmarshal([]byte(rawResponse), &errorResp) == nil && errorResp.Error.Message != "" {
		panic(fmt.Errorf("API Error: %s (Code: %d)", errorResp.Error.Message, errorResp.Error.Code))
	}

	var completion GetCompletionResponse
	if err := json.Unmarshal([]byte(rawResponse), &completion); err != nil {
		panic(err)
	}

	// Add AI message to conversation
	aiMessage := completion.Choices[0].Message
	CurrentConversation.Messages = append(CurrentConversation.Messages, aiMessage)

	return aiMessage
}
