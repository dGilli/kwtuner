package services

import (
	"context"
	"fmt"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type AnthropicService struct {
    apiKey string
    client *anthropic.Client
}

func NewAnthropicService(apiKey string) *AnthropicService {
    return &AnthropicService{
        apiKey: apiKey,
        client: anthropic.NewClient(option.WithAPIKey(apiKey)),
    }
}

func (a *AnthropicService) GenerateText(keywords []string, text string) (string, error) {
    // Create a simple prompt
    prompt := fmt.Sprintf(
        "Please integrate these keywords %v into the following text: %s",
        keywords, text,
    )

    message, err := a.client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(anthropic.ModelClaude3_5SonnetLatest),
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		}),
	})
	if err != nil {
        return "", err
	}

    return fmt.Sprintf("%+v\n", message.Content), nil
}

