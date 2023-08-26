package openai

import (
	"context"
	"fmt"
	"os"

	openaiApi "github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	client *openaiApi.Client
}

func NewOpenAIClient(token string) *OpenAI {
	client := openaiApi.NewClient(token)
	return &OpenAI{client: client}
}

func (o *OpenAI) CreateNews(town string) (result string, err error) {
	println(os.Getenv("CHATPROMPT"))
	resp, err := o.client.CreateChatCompletion(
		context.Background(),
		openaiApi.ChatCompletionRequest{
			Model: openaiApi.GPT3Dot5Turbo,
			Messages: []openaiApi.ChatCompletionMessage{
				{
					Role:    openaiApi.ChatMessageRoleUser,
					Content: fmt.Sprintf(os.Getenv("CHATPROMPT"), town),
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return result, err
	}

	return resp.Choices[0].Message.Content, nil
}
