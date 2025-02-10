package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func GetCheatSheet(topic string) <-chan string {
	systemPrompt :=
		"You are a helpful assistant that generates, markdown formatted cheat sheets for a given topic, you must use markdown formatting. The given topic is: "
	requestBody := fmt.Sprintf(`{
		"model": "mistral",
		"prompt": "%s%s",
		"stream": true
	}`, systemPrompt, topic)

	output := make(chan string)

	go func() {
        defer close(output)

		res, err := http.Post(
			"http://localhost:11434/api/generate",
			"application/json",
			bytes.NewBuffer([]byte(requestBody)),
		)

		if err != nil {
			output <- "An error occured while requesting: " + err.Error()
			return
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			output <- "Unexpected status code: " + res.Status
			return
		}

		decoder := json.NewDecoder(res.Body)
		for {
			var chunk OllamaResponse
			if err := decoder.Decode(&chunk); err != nil {
				if err == io.EOF {
					break
				}
				output <- "An error occurred while parsing: " + err.Error()
				return
			}
			output <- chunk.Response
			if chunk.Done {
				break
			}
		}
	}()

	return output
}
