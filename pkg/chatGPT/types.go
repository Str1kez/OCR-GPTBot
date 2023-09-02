package chatgpt

import "github.com/sashabaranov/go-openai"

type ChatCompletionClient struct {
	client *openai.Client
}

type CompletionClient struct {
	client *openai.Client
}
