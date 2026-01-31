# ProgrammingCLI

## Overview
ProgrammingCLI is a command-line interface tool written in Go that enables interaction with AI agents through the OpenRouter API. The application provides a framework for managing conversations with various AI models and accessing their capabilities programmatically.

## Project Structure
- **main.go**: The entry point of the application. Initializes the process, loads configuration, and sets up the CLI environment.
- **utils/**: Contains utility packages for core functionality.
  - **openai.go**: Handles interactions with the OpenRouter API (`https://openrouter.ai/api/v1`). Defines data structures for models, messages, and conversations, and provides functions for fetching available models and getting completions from AI agents.
  - **config.go**: Manages application configuration including API keys and default model selection.
  - **menu.go**: Provides display utilities for the CLI interface.
  - **state.go**: Manages application state and conversation history.
  - **fun.go**: Additional utility functions.
- **config.json**: Stores user configuration including API key and preferred model.

## Dependencies
- `github.com/carlmjohnson/requests`: Used for making HTTP requests to the OpenRouter API.

## Usage
To run the project:
```bash
go run main.go
```

To build the executable:
```bash
go build -o ProgrammingCLI.exe
```

## Key Components

### Agent Configuration
The application uses a configuration system to store API credentials and model preferences:
- **Default Model**: `minimax/minimax-m2.1` (selected for cost-effectiveness)
- **API Key**: Stored in `config.json` and loaded at startup

### Data Structures

#### Model Struct
Represents an AI model available through OpenRouter with an ID and Name.

#### Message Struct
Represents a single message in a conversation with:
- `Role`: The sender's role (user, assistant, system)
- `Content`: The message content
- `Reasoning`: Optional reasoning information (for models that support it)

#### Conversation Struct
Manages the conversation history with a list of messages, enabling multi-turn interactions with AI agents.

#### GetCompletionResponse
Contains the full API response including:
- Model information
- Generated choices
- Token usage statistics
- Provider details

### Core Functions

#### GetModels(apiKey string)
Fetches the list of available AI models from the OpenRouter API. Returns a `GetModelsResponse` containing all accessible models.

#### GetCompletion(userMessage string, modelId *string, apiKey string)
Sends a user message to the specified AI agent and returns the response. Features:
- Automatically maintains conversation history
- Supports custom model selection or uses default
- Handles error responses gracefully
- Appends both user and AI messages to the conversation

## API Integration
The application integrates with OpenRouter API at `https://openrouter.ai/api/v1`, which provides access to multiple AI model providers through a unified interface. All API calls require authentication via Bearer token.

## Error Handling
The application includes comprehensive error handling for:
- API request failures
- JSON parsing errors
- API error responses with detailed error codes and messages
