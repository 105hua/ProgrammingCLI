package utils

// Default system prompt for the AI assistant.
const DefaultSystemPrompt = "You are a helpful assistant."

// CurrentConversation holds the active conversation state.
var CurrentConversation = Conversation{
	Messages: []Message{
		{
			Role:    "system",
			Content: DefaultSystemPrompt,
		},
	},
}

// ResetConversation clears the conversation history and starts fresh.
func ResetConversation() {
	CurrentConversation = Conversation{
		Messages: []Message{
			{
				Role:    "system",
				Content: DefaultSystemPrompt,
			},
		},
	}
}
