package services

import (
	"context"
	"fmt"
	"strings"

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
    prompt := constructPrompt(map[string]string{
        "ProductDescription": text,
        "Keywords":           strings.Join(keywords, ", "),
    })

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

	content, err := extractContent(resp.Choices[0].Message.Content, "rewritten_description")
	if err != nil {
		return "", err
	}

    if len(resp.Choices) > 0 {
        return content, nil
    }

    return "", fmt.Errorf("no content in OpenAI response")
}

