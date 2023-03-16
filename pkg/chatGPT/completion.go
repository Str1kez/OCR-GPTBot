package chatgpt

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func (c *ChatGPT) CreateCompletion(text string) (openai.CompletionResponse, error) {
	completionRequest := generateCompletionRequest(text)
	return c.client.CreateCompletion(context.Background(), completionRequest)
}

func (c *ChatGPT) CreateChatCompletion(text string) (openai.ChatCompletionResponse, error) {
	completionRequest := generateChatCompletionRequest(text)
	return c.client.CreateChatCompletion(context.Background(), completionRequest)
}

func generateChatCompletionRequest(text string) openai.ChatCompletionRequest {
	message := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: text}
	return openai.ChatCompletionRequest{
		Model:            openai.GPT3Dot5Turbo,
		Messages:         []openai.ChatCompletionMessage{message},
		MaxTokens:        2048,
		Temperature:      0.5,
		TopP:             1.0,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
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
