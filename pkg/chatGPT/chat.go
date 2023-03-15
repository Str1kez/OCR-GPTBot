package chatgpt

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type ChatGPT struct {
	client *openai.Client
}

func NewChatGPT(token string) *ChatGPT {
	return &ChatGPT{client: openai.NewClient(token)}
}

func (c *ChatGPT) CreateCompletion(text string) (openai.CompletionResponse, error) {
	completionRequest := generateCompletionRequest(text)
	return c.client.CreateCompletion(context.Background(), completionRequest)
}

func generateCompletionRequest(text string) openai.CompletionRequest {
	return openai.CompletionRequest{
		Model:            openai.GPT3Dot5Turbo,
		Prompt:           text,
		MaxTokens:        2048,
		Temperature:      0.5,
		TopP:             1.0,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}

}
