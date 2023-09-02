package chatgpt

import (
	"github.com/sashabaranov/go-openai"
)

func NewChatCompletionClient(token string) *ChatCompletionClient {
	return &ChatCompletionClient{client: openai.NewClient(token)}
}
