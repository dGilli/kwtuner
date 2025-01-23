package services

import (
	"fmt"
	"strings"
	"sync"
)

var (
    registry = make(map[string]AIService)
    mu       sync.RWMutex
)

type AIService interface {
    GenerateText(keywords []string, text string) (string, error)
}

func RegisterService(name string, service AIService) {
    mu.Lock()
    defer mu.Unlock()
    registry[name] = service
}

func GetService(name string) (AIService, bool) {
    mu.RLock()
    defer mu.RUnlock()
    s, ok := registry[name]
    return s, ok
}

var (
    Prompt string
)

func SetPrompt(prompt string) {
    Prompt = prompt
}

func constructPrompt(d map[string]string) string {
    var p = Prompt
	p = strings.ReplaceAll(p, "{ProductDescription}", d["ProductDescription"])
	p = strings.ReplaceAll(p, "{Keywords}", d["Keywords"])

    return p
}

func extractContent(htmlContent, tagName string) (string, error) {
	openTag := fmt.Sprintf("<%s>", tagName)
	closeTag := fmt.Sprintf("</%s>", tagName)

	// Find the opening and closing tag positions
	startIdx := strings.Index(htmlContent, openTag)
	if startIdx == -1 {
		return "", fmt.Errorf("opening tag <%s> not found", tagName)
	}
	endIdx := strings.Index(htmlContent, closeTag)
	if endIdx == -1 {
		return "", fmt.Errorf("closing tag </%s> not found", tagName)
	}

	startIdx += len(openTag) // Move index to the end of the opening tag
	if startIdx >= endIdx {
		return "", fmt.Errorf("invalid content for <%s>", tagName)
	}

	return htmlContent[startIdx:endIdx], nil
}

