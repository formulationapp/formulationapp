package client

import "github.com/formulationapp/formulationapp/internal/dto"

type Clients interface {
	AI() AIClient
}

type clients struct {
	aiClient AIClient
}

func NewClients(config dto.Config) Clients {
	return &clients{
		aiClient: newAIClient(config.DefaultOpenAIToken),
	}
}

func (c clients) AI() AIClient {
	return c.aiClient
}
