package invoke

import (
	"context"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func sdk_ai_invoke() {
	client := openai.NewClient(
		option.WithAPIKey(API_KEY),
		option.WithBaseURL(
			"https://dashscope.aliyuncs.com/compatible-mode/v1/",
		),
	)
	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(), openai.ChatCompletionNewParams{
			Messages: openai.F(
				[]openai.ChatCompletionMessageParamUnion{
					openai.UserMessage(
						"你是谁",
					),
				},
			),
			Model: openai.F(
				"qwen-plus",
			),
		},
	)

	if err != nil {
		panic(err.Error())
	}

	println(chatCompletion.Choices[0].Message.Content)

}
