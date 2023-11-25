package chatgpt

import (
	"context"

	"github.com/Str1kez/OCR-GPTBot/pkg/telegram"

	"github.com/sashabaranov/go-openai"
)

func (c *ChatCompletionClient) GetCompletionText(text string, settings telegram.Settings) (string, error) {
	ctx := context.Background()
	completionRequest := generateChatCompletionRequest(text, settings.Context, settings.Temperature, settings.FrequencyPenalty, settings.Stream)
	response, err := c.client.CreateChatCompletion(ctx, completionRequest)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Message.Content, nil
}

func (c *ChatCompletionClient) GetCompletionStream(text string, settings telegram.Settings) (*openai.ChatCompletionStream, error) {
	ctx := context.Background()
	request := generateChatCompletionRequest(text, settings.Context, settings.Temperature, settings.FrequencyPenalty, settings.Stream)
	return c.client.CreateChatCompletionStream(ctx, request)
}

func generateChatCompletionRequest(text, context string, temperature, penalty float32, stream bool) openai.ChatCompletionRequest {
	userMessage := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: text}
	contextMessage := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: context}
	return openai.ChatCompletionRequest{
		Model:            openai.GPT3Dot5Turbo,
		Messages:         []openai.ChatCompletionMessage{contextMessage, userMessage},
		MaxTokens:        2048,
		Temperature:      temperature,
		Stream:           stream,
		TopP:             1.0,
		FrequencyPenalty: penalty,
		PresencePenalty:  0,
	}
}
