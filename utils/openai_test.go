package utils

import (
	"encoding/json"
	"testing"
)

func TestMessageJSONMarshaling(t *testing.T) {
	testCases := []struct {
		name    string
		message Message
	}{
		{
			name: "Basic message",
			message: Message{
				Role:    "user",
				Content: "Hello, world!",
			},
		},
		{
			name: "Message with reasoning",
			message: Message{
				Role:      "assistant",
				Content:   "The answer is 42",
				Reasoning: "Based on deep thought",
			},
		},
		{
			name: "Empty content",
			message: Message{
				Role:    "system",
				Content: "",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Marshal to JSON
			data, err := json.Marshal(tc.message)
			if err != nil {
				t.Fatalf("Failed to marshal message: %v", err)
			}

			// Unmarshal back
			var decoded Message
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("Failed to unmarshal message: %v", err)
			}

			// Verify values match
			if decoded.Role != tc.message.Role {
				t.Errorf("Role = %s; want %s", decoded.Role, tc.message.Role)
			}

			if decoded.Content != tc.message.Content {
				t.Errorf("Content = %s; want %s", decoded.Content, tc.message.Content)
			}

			if decoded.Reasoning != tc.message.Reasoning {
				t.Errorf("Reasoning = %s; want %s", decoded.Reasoning, tc.message.Reasoning)
			}
		})
	}
}

func TestConversationManagement(t *testing.T) {
	// Create a new conversation
	conv := Conversation{
		Messages: []Message{},
	}

	// Add messages
	conv.Messages = append(conv.Messages, Message{
		Role:    "system",
		Content: "You are a helpful assistant.",
	})

	conv.Messages = append(conv.Messages, Message{
		Role:    "user",
		Content: "Hello",
	})

	conv.Messages = append(conv.Messages, Message{
		Role:    "assistant",
		Content: "Hi there!",
	})

	// Verify message count
	if len(conv.Messages) != 3 {
		t.Errorf("Conversation length = %d; want 3", len(conv.Messages))
	}

	// Verify message order and content
	if conv.Messages[0].Role != "system" {
		t.Errorf("First message role = %s; want system", conv.Messages[0].Role)
	}

	if conv.Messages[1].Role != "user" {
		t.Errorf("Second message role = %s; want user", conv.Messages[1].Role)
	}

	if conv.Messages[2].Role != "assistant" {
		t.Errorf("Third message role = %s; want assistant", conv.Messages[2].Role)
	}
}

func TestModelJSONMarshaling(t *testing.T) {
	testCases := []struct {
		name  string
		model Model
	}{
		{
			name: "Standard model",
			model: Model{
				ID:   "gpt-4",
				Name: "GPT-4",
			},
		},
		{
			name: "Model with slashes",
			model: Model{
				ID:   "minimax/minimax-m2.1",
				Name: "Minimax M2.1",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Marshal to JSON
			data, err := json.Marshal(tc.model)
			if err != nil {
				t.Fatalf("Failed to marshal model: %v", err)
			}

			// Unmarshal back
			var decoded Model
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("Failed to unmarshal model: %v", err)
			}

			// Verify values match
			if decoded.ID != tc.model.ID {
				t.Errorf("ID = %s; want %s", decoded.ID, tc.model.ID)
			}

			if decoded.Name != tc.model.Name {
				t.Errorf("Name = %s; want %s", decoded.Name, tc.model.Name)
			}
		})
	}
}

func TestGetModelsResponseJSONMarshaling(t *testing.T) {
	response := GetModelsResponse{
		Data: []Model{
			{ID: "model-1", Name: "Model 1"},
			{ID: "model-2", Name: "Model 2"},
		},
	}

	// Marshal to JSON
	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal GetModelsResponse: %v", err)
	}

	// Unmarshal back
	var decoded GetModelsResponse
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal GetModelsResponse: %v", err)
	}

	// Verify values match
	if len(decoded.Data) != len(response.Data) {
		t.Errorf("Data length = %d; want %d", len(decoded.Data), len(response.Data))
	}

	for i := range response.Data {
		if decoded.Data[i].ID != response.Data[i].ID {
			t.Errorf("Data[%d].ID = %s; want %s", i, decoded.Data[i].ID, response.Data[i].ID)
		}
		if decoded.Data[i].Name != response.Data[i].Name {
			t.Errorf("Data[%d].Name = %s; want %s", i, decoded.Data[i].Name, response.Data[i].Name)
		}
	}
}

func TestGetCompletionResponseJSONMarshaling(t *testing.T) {
	response := GetCompletionResponse{
		Id:      "chatcmpl-123",
		Object:  "chat.completion",
		Created: 1234567890,
		Model:   "gpt-4",
		Choices: []Choice{
			{
				Message: Message{
					Role:    "assistant",
					Content: "Hello!",
				},
				Index:        0,
				FinishReason: "stop",
			},
		},
		Usage: Usage{
			PromptTokens:     10,
			CompletionTokens: 5,
			TotalTokens:      15,
		},
	}

	// Marshal to JSON
	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal GetCompletionResponse: %v", err)
	}

	// Unmarshal back
	var decoded GetCompletionResponse
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal GetCompletionResponse: %v", err)
	}

	// Verify core values
	if decoded.Id != response.Id {
		t.Errorf("Id = %s; want %s", decoded.Id, response.Id)
	}

	if decoded.Model != response.Model {
		t.Errorf("Model = %s; want %s", decoded.Model, response.Model)
	}

	if len(decoded.Choices) != len(response.Choices) {
		t.Errorf("Choices length = %d; want %d", len(decoded.Choices), len(response.Choices))
	}

	if decoded.Usage.TotalTokens != response.Usage.TotalTokens {
		t.Errorf("Usage.TotalTokens = %d; want %d", decoded.Usage.TotalTokens, response.Usage.TotalTokens)
	}
}

func TestErrorResponseJSONMarshaling(t *testing.T) {
	errorResp := ErrorResponse{}
	errorResp.Error.Message = "Invalid API key"
	errorResp.Error.Code = 401

	// Marshal to JSON
	data, err := json.Marshal(errorResp)
	if err != nil {
		t.Fatalf("Failed to marshal ErrorResponse: %v", err)
	}

	// Unmarshal back
	var decoded ErrorResponse
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Failed to unmarshal ErrorResponse: %v", err)
	}

	// Verify values match
	if decoded.Error.Message != errorResp.Error.Message {
		t.Errorf("Error.Message = %s; want %s", decoded.Error.Message, errorResp.Error.Message)
	}

	if decoded.Error.Code != errorResp.Error.Code {
		t.Errorf("Error.Code = %d; want %d", decoded.Error.Code, errorResp.Error.Code)
	}
}

func TestUsageStruct(t *testing.T) {
	usage := Usage{
		PromptTokens:     100,
		CompletionTokens: 50,
		TotalTokens:      150,
	}

	// Verify total is sum of prompt and completion
	expectedTotal := usage.PromptTokens + usage.CompletionTokens
	if usage.TotalTokens != expectedTotal {
		t.Errorf("TotalTokens = %d; want %d (sum of prompt and completion)", usage.TotalTokens, expectedTotal)
	}
}

func TestDefaultModel(t *testing.T) {
	// Verify the default model constant
	if DEFAULT_MODEL == "" {
		t.Error("DEFAULT_MODEL is empty")
	}

	expected := "minimax/minimax-m2.1"
	if DEFAULT_MODEL != expected {
		t.Errorf("DEFAULT_MODEL = %s; want %s", DEFAULT_MODEL, expected)
	}
}

func TestBaseURL(t *testing.T) {
	// Verify the base URL constant
	if BASE_URL == "" {
		t.Error("BASE_URL is empty")
	}

	expected := "https://openrouter.ai/api/v1"
	if BASE_URL != expected {
		t.Errorf("BASE_URL = %s; want %s", BASE_URL, expected)
	}
}

func TestCurrentConversationInitialization(t *testing.T) {
	// Verify that CurrentConversation is initialized with a system message
	if len(CurrentConversation.Messages) == 0 {
		t.Error("CurrentConversation.Messages is empty; expected initial system message")
	}

	if CurrentConversation.Messages[0].Role != "system" {
		t.Errorf("CurrentConversation.Messages[0].Role = %s; want system", CurrentConversation.Messages[0].Role)
	}

	if CurrentConversation.Messages[0].Content == "" {
		t.Error("CurrentConversation.Messages[0].Content is empty")
	}
}
