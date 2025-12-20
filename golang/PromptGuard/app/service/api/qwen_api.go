package main

import (
	"context"
	"os"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("DASHSCOPE_API_KEY")),
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1"),
	)
	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("你是谁"),
			},
			Model: "qwen3-8b",
		},
	)

	if err != nil {
		panic(err.Error())
	}

	println(chatCompletion.Choices[0].Message.Content)
}