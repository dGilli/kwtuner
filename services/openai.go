package services

import (
    "context"
    "fmt"

    openai "github.com/sashabaranov/go-openai"
)

// OpenAIService is a simple struct for storing OpenAI-specific config.
type OpenAIService struct {
    apiKey string
    client *openai.Client
}

// NewOpenAIService creates a new OpenAI service client with the given API key.
func NewOpenAIService(apiKey string) *OpenAIService {
    return &OpenAIService{
        apiKey: apiKey,
        client: openai.NewClient(apiKey),
    }
}

// GenerateText calls the OpenAI API, passing in your text and keywords.
// This is just a placeholder - adapt as needed to your real prompts.
func (o *OpenAIService) GenerateText(keywords []string, text string) (string, error) {
    prompt := fmt.Sprintf("Integrate these keywords %v into the text: %s", keywords, text)

    resp, err := o.client.CreateChatCompletion(
        context.Background(),
        openai.ChatCompletionRequest{
            Model: "gpt-3.5-turbo",
            Messages: []openai.ChatCompletionMessage{
                {
                    Role:    openai.ChatMessageRoleSystem,
                    Content: "You are a helpful assistant.",
                },
                {
                    Role:    openai.ChatMessageRoleUser,
                    Content: prompt,
                },
            },
        },
    )
    if err != nil {
        return "", err
    }

    // Return the model's reply
    if len(resp.Choices) > 0 {
        return resp.Choices[0].Message.Content, nil
    }

    return "", fmt.Errorf("no content in OpenAI response")
}

