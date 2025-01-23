package services

import (
	"context"
	"strings"

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
    prompt := constructPrompt(map[string]string{
        "ProductDescription": text,
        "Keywords":           strings.Join(keywords, ", "),
    })

    message, err := a.client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.F(anthropic.ModelClaude3_5SonnetLatest),
		MaxTokens: anthropic.F(int64(1024)),
		Messages:  anthropic.F([]anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		}),
	})
	if err != nil {
        return "", err
	}

	content, err := extractContent(message.Content[0].Text, "rewritten_description")
	if err != nil {
		return "", err
	}

    return content, nil
}

