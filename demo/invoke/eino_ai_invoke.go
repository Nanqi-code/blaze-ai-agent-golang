package invoke

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/schema"
	"time"
)

func eion_ai_invoke() {
	ctx := context.Background()

	// 初始化模型
	model, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  API_KEY,
		Timeout: 30 * time.Second,
		Model:   "qwen-plus",
		BaseURL: "https://dashscope.aliyuncs.com/compatible-mode/v1/",
	})
	if err != nil {
		panic(err)
	}

	// 准备消息
	messages := []*schema.Message{
		schema.SystemMessage("你是一个小狗，你只会“汪汪汪”"),
		schema.UserMessage("你好呀"),
	}

	// 同步回复
	response, err := model.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}

	// 处理回复
	println(response.Content)

	//// 获取流式回复
	//reader, err := model.Stream(ctx, messages)
	//if err != nil {
	//	panic(err)
	//}
	//defer reader.Close() // 注意要关闭
	//
	//// 处理流式内容
	//for {
	//	chunk, err := reader.Recv()
	//	if err != nil {
	//		break
	//	}
	//	print(chunk.Content)
	//}
}
