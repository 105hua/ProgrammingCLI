# ProgrammingCLI

## Overview
ProgrammingCLI is a command-line interface tool written in Go. Currently, it interacts with the OpenRouter API to retrieve a list of available AI models.

## Project Structure
- **main.go**: The entry point of the application. It initializes the process and calls the utility functions.
- **utils/**: Contains utility packages.
  - **openai.go**: Handles interactions with the OpenRouter API (despite the name, it uses `https://openrouter.ai/api/v1`). It defines the `Model` structure and the `GetModels` function.
- **.gemini/**: Contains configuration settings for the Gemini environment.
  - **settings.json**: Stores general settings like preview features.

## Dependencies
- `github.com/carlmjohnson/requests`: Used for making HTTP requests to the API.

## Usage
To run the project:
```bash
go run main.go
```

## Key Components
### Model Struct
Represents an AI model with an ID and Name.

### GetModels Function
Fetches the list of models from `https://openrouter.ai/api/v1/models` and returns a `GetModelsResponse`.
