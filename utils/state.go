package utils

var CurrentConversation = Conversation{
	Messages: []Message{
		{
			Role:    "system",
			Content: "You are a helpful assistant.",
		},
	},
}
