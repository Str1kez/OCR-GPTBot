package chatgpt

import (
	"github.com/sashabaranov/go-openai"
)

type ChatCompletionClient struct {
	client *openai.Client
}

type CompletionClient struct {
	client *openai.Client
}

func NewChatCompletionClient(token string) *ChatCompletionClient {
	return &ChatCompletionClient{client: openai.NewClient(token)}
}

func NewCompletionClient(token string) *CompletionClient {
	return &CompletionClient{client: openai.NewClient(token)}
}
