package chatgpt

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

func (c *ChatCompletionClient) CreateChatCompletion(text, userContext string) (openai.ChatCompletionResponse, error) {
	completionRequest := generateChatCompletionRequest(text, userContext, false)
	return c.client.CreateChatCompletion(context.Background(), completionRequest)
}

func (c *ChatCompletionClient) PerformCompletion(text, userContext string) (string, error) {
	response, err := c.CreateChatCompletion(text, userContext)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Message.Content, nil
}

func (c *ChatCompletionClient) GetCompletionStream(text, userContext string) (*openai.ChatCompletionStream, error) {
	ctx := context.Background()
	request := generateChatCompletionRequest(text, userContext, true)
	return c.client.CreateChatCompletionStream(ctx, request)
}

func generateChatCompletionRequest(text, ctx string, stream bool) openai.ChatCompletionRequest {
	userMessage := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: text}
	contextMessage := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: ctx}
	return openai.ChatCompletionRequest{
		Model:            openai.GPT3Dot5Turbo,
		Messages:         []openai.ChatCompletionMessage{contextMessage, userMessage},
		MaxTokens:        2048,
		Temperature:      0.5,
		Stream:           stream,
		TopP:             1.0,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
}
