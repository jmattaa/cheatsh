package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// my bigbrain does curl into | jq and we fire
type OllamaResponse struct {
	Model           string `json:"model"`
	CreatedAt       string `json:"created_at"`
	Response        string `json:"response"`
	Done            bool   `json:"done"`
	DoneReason      string `json:"done_reason"`
	Context         []int  `json:"context"`
	TotalDuration   int    `json:"total_duration"`
	LoadDuration    int    `json:"load_duration"`
	PromptEvalCount int    `json:"prompt_eval_count"`
	EvalCount       int    `json:"eval_count"`
	EvalDuration    int    `json:"eval_duration"`
}

// TODO: make this stream the repsonse instead of just returning a string
// cuz it be nicer when you see what's happening
func GetCheatSheet(topic string) string {
	systemPrompt :=
		"You are a helpful assistant that generates, markdown formatted cheat sheets for a given topic, you must use markdown formatting. The given topic is: "

	requestBody := fmt.Sprintf(`{
		"model": "mistral",
		"prompt": "%s%s",
		"stream": false
	}`, systemPrompt, topic)

	res, err := http.Post(
		"http://localhost:11434/api/generate",
		"application/json",
		bytes.NewBuffer([]byte(requestBody)),
	)

	if err != nil {
		return "An error occured while requesting: " + err.Error()
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "Unexpected status code: " + res.Status
	}

	OllamaResponse := OllamaResponse{}
	derr := json.NewDecoder(res.Body).Decode(&OllamaResponse)
	if derr != nil {
		return "An error occured while decoding " + derr.Error()
	}

	if OllamaResponse.Response == "" {
		return "Got an empty response wut :("
	}

	return OllamaResponse.Response
}
