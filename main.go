package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgilli/kwtuner/handlers"
	"github.com/dgilli/kwtuner/services"
)

func init() {
    services.SetPrompt(prompt())
    services.RegisterService("openai", services.NewOpenAIService(""))
    services.RegisterService("anthropic", services.NewAnthropicService(""))
}

func main() {
    http.HandleFunc("/generate", handlers.New(handlers.HandleGenerateIndex))
    fmt.Println("Starting server on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

