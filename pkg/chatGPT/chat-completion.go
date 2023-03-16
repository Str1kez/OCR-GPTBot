package chatgpt

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func (c *ChatCompletionClient) CreateChatCompletion(text string) (openai.ChatCompletionResponse, error) {
	completionRequest := generateChatCompletionRequest(text)
	return c.client.CreateChatCompletion(context.Background(), completionRequest)
}

func (c *ChatCompletionClient) PerformCompletion(text string) (string, error) {
	response, err := c.CreateChatCompletion(text)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Message.Content, nil
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
