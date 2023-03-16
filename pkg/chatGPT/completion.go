package chatgpt

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func (c *CompletionClient) CreateCompletion(text string) (openai.CompletionResponse, error) {
	completionRequest := generateCompletionRequest(text)
	return c.client.CreateCompletion(context.Background(), completionRequest)
}

func (c *CompletionClient) PerformCompletion(text string) (string, error) {
	response, err := c.CreateCompletion(text)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Text, nil
}

func generateCompletionRequest(text string) openai.CompletionRequest {
	return openai.CompletionRequest{
		Model:            openai.GPT3TextDavinci003,
		Prompt:           text,
		MaxTokens:        2048,
		Temperature:      0.5,
		TopP:             1.0,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
}
