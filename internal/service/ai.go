package service

import (
	"bytes"
	"context"
	_ "embed"
	"github.com/formulationapp/formulationapp/internal/client"
	"github.com/sirupsen/logrus"
	"html/template"
)

//go:embed assets/openai-generate-form.txt
var generateFormTemplate string

type AIService interface {
	GenerateForm(ctx context.Context, prompt string) (string, error)
}

type aiService struct {
	aiClient client.AIClient
}

func newAIService(aiClient client.AIClient) AIService {
	return &aiService{aiClient: aiClient}
}

func (a aiService) GenerateForm(ctx context.Context, prompt string) (string, error) {
	tpl, err := template.New("openai").Parse(generateFormTemplate)
	if err != nil {
		return "", err
	}
	var rendered bytes.Buffer
	if err := tpl.Execute(&rendered, struct {
		Prompt string
	}{
		Prompt: prompt,
	}); err != nil {
		return "", err
	}

	output, usedTokens, err := a.aiClient.Complete(ctx, rendered.String())
	if err != nil {
		return "", err
	}
	logrus.Infof("used %d tokens", usedTokens)
	return output, nil
}
