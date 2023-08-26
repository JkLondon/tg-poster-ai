package openai

import (
	"context"
	"encoding/base64"
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

func (o *OpenAI) CreateNews(ctx context.Context, town string) (result string, image []byte, err error) {
	println(os.Getenv("CHATPROMPT"))
	resp, err := o.client.CreateChatCompletion(
		ctx,
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
		return result, nil, err
	}

	reqBase64 := openaiApi.ImageRequest{
		Prompt:         "бабушка с советским флагом в руках под украиной",
		Size:           openaiApi.CreateImageSize256x256,
		ResponseFormat: openaiApi.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	respBase64, err := o.client.CreateImage(ctx, reqBase64)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		return
	}

	return resp.Choices[0].Message.Content, imgBytes, nil
}
