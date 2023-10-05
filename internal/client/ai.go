package client

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type AIClient interface {
	Complete(ctx context.Context, prompt string) (string, int, error)
}

type aiClient struct {
	ai *openai.Client
}

func newAIClient(token string) AIClient {
	return &aiClient{ai: openai.NewClient(token)}
}

func (a aiClient) Complete(ctx context.Context, prompt string) (string, int, error) {
	response, err := a.ai.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", 0, err
	}
	return response.Choices[0].Message.Content, response.Usage.TotalTokens, nil
}
