package services

import (
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

