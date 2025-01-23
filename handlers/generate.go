package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgilli/kwtuner/services"
	"github.com/dgilli/kwtuner/views/home"
)

type GeneratePayload struct {
    APIKey   string   `json:"apikey"`
    Service  string   `json:"service"`
    Keywords []string `json:"keywords"`
    Text     string   `json:"text"`
}

func HandleGenerate(w http.ResponseWriter, r *http.Request) error {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
        return fmt.Errorf("invalid method: %s", r.Method)
    }

    var p GeneratePayload

    contentType := r.Header.Get("Content-Type")
    if strings.Contains(contentType, "application/json") {
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return err
		}
    } else {
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Unable to parse form", http.StatusBadRequest)
            return err
        }
        p.APIKey = r.FormValue("apikey")
        p.Service = r.FormValue("service")
        p.Keywords = strings.Split(r.FormValue("keywords"), ", ")
        p.Text = r.FormValue("text")
    }

    service, ok := services.GetService(p.Service)
    if !ok {
        err := errors.New(fmt.Sprintf("service not found: %s", p.Service))
        http.Error(w, fmt.Sprintf("AI error: %v", err), http.StatusInternalServerError)
        return err
    }

    if p.APIKey != "" && p.Service == "openai" {
        service = services.NewOpenAIService(p.APIKey)
    } else if p.APIKey != "" && p.Service == "claude" {
        service = services.NewAnthropicService(p.APIKey)
    } else {
        err := errors.New("api key not set")
        http.Error(w, fmt.Sprintf("AI error: %v", err), http.StatusInternalServerError)
        return err
    }

    generated, err := service.GenerateText(p.Keywords, p.Text)
    if err != nil {
        http.Error(w, fmt.Sprintf("AI error: %v", err), http.StatusInternalServerError)
        return err
    }

    accept := r.Header.Get("Accept")
    if strings.Contains(accept, "application/json") {
        w.Header().Set("Content-Type", "application/json")
        return json.NewEncoder(w).Encode(map[string]string{"response": strings.TrimSpace(generated)})
    } else {
	    return Render(w, r, home.Result(generated))
    }
}
