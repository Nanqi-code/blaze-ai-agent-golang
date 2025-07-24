package invoke

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func langchaingo_ai_invoke() {
	// 初始化大模型
	llm, err := openai.New(
		openai.WithModel("qwen-plus"),
		openai.WithToken(API_KEY),
		openai.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1/"),
	)
	if err != nil {
		panic(err.Error())
	}
	ctx := context.Background()

	prompt, err := llms.GenerateFromSinglePrompt(ctx, llm, "你好吗")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(prompt)
}
