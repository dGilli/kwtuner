package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgilli/kwtuner/services"
)

type AIRequest struct {
    APIKey   string   `json:"apikey"`
    Service  string   `json:"service"`
    Keywords []string `json:"keywords"`
    Text     string   `json:"text"`
}

func HandleGenerateIndex(w http.ResponseWriter, r *http.Request) error {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return fmt.Errorf("invalid method: %s", r.Method)
    }

    var req AIRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return err
    }

    service, ok := services.GetService(req.Service)
    if !ok {
        err := errors.New("AI error, service not found")
        http.Error(w, fmt.Sprintf("AI error: %v", err), http.StatusInternalServerError)
        return err
    }

    if req.APIKey != "" && req.Service == "openai" {
        service = services.NewOpenAIService(req.APIKey)
    }
    if req.APIKey != "" && req.Service == "anthropic" {
        service = services.NewAnthropicService(req.APIKey)
    }

    generatedText, err := service.GenerateText(req.Keywords, req.Text)
    if err != nil {
        http.Error(w, fmt.Sprintf("AI error: %v", err), http.StatusInternalServerError)
        return err
    }

    w.Header().Set("Content-Type", "application/json")
    return json.NewEncoder(w).Encode(map[string]string{"response": generatedText})
}
