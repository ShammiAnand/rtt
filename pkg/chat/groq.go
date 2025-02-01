package chat

import (
	"context"
	"fmt"
	"os"

	"github.com/carlmjohnson/requests"
)

type groqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type groqRequest struct {
	Model    string        `json:"model"`
	Messages []groqMessage `json:"messages"`
}

type groqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func ProcessQuery(ctx context.Context, filePath string, question string, apiKey string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	req := groqRequest{
		Model: "mixtral-8x7b-32768", // using Mixtral for high context
		Messages: []groqMessage{
			{
				Role:    "system",
				Content: "You are a helpful assistant. Answer questions about the provided content.",
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Content:\n\n%s\n\nQuestion: %s", string(content), question),
			},
		},
	}

	return respond(ctx, req, apiKey)
}

func respond(ctx context.Context, req groqRequest, apiKey string) error {
	var resp groqResponse
	err := requests.
		URL("https://api.groq.com/openai/v1/chat/completions").
		Header("Authorization", "Bearer "+apiKey).
		BodyJSON(&req).
		ToJSON(&resp).
		Fetch(ctx)

	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	if len(resp.Choices) > 0 {
		fmt.Println()
		fmt.Println(resp.Choices[0].Message.Content)
	}

	return nil
}
