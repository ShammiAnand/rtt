package chat

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/carlmjohnson/requests"
	"github.com/shammianand/rtt/utils/logger"
)

type groqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type groqRequest struct {
	Model    string        `json:"model"`
	Messages []groqMessage `json:"messages"`
	Stream   bool          `json:"stream"`
}

type groqResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func ProcessQuery(ctx context.Context, filePath string, question string, apiKey string, stream bool) error {
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
		Stream: stream,
	}

	if stream {
		return streamResponse(ctx, req, apiKey)
	}

	return normalResponse(ctx, req, apiKey)
}

func normalResponse(ctx context.Context, req groqRequest, apiKey string) error {
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

func streamResponse(ctx context.Context, req groqRequest, apiKey string) error {
	var buf bytes.Buffer
	err := requests.
		URL("https://api.groq.com/openai/v1/chat/completions").
		Header("Authorization", "Bearer "+apiKey).
		BodyJSON(&req).
		ToWriter(&buf).
		Fetch(ctx)

	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	fmt.Println()

	decoder := json.NewDecoder(bytes.NewReader(buf.Bytes()))
	for {
		var chunk struct {
			Choices []struct {
				Delta struct {
					Content string `json:"content"`
				} `json:"delta"`
			} `json:"choices"`
		}

		if err := decoder.Decode(&chunk); err != nil {
			if err == io.EOF {
				break
			}
			logger.Log.Error("Error decoding chunk:", err)
			continue
		}

		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			fmt.Print(chunk.Choices[0].Delta.Content)
		}
	}

	return nil
}
