package chatgpt

import (
	"github.com/sashabaranov/go-openai"
)

type ChatGPT struct {
	client *openai.Client
}

func NewChatGPT(token string) *ChatGPT {
	return &ChatGPT{client: openai.NewClient(token)}
}
