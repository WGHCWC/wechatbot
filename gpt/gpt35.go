// @Author Bing
// @Date 2023/3/6 20:31:00
// @Desc
package gpt

import (
	"context"
	"errors"
	"github.com/otiai10/openaigo"
	"github.com/qingconglaixueit/abing_logger"
	"github.com/qingconglaixueit/wechatbot/config"
)

var MyGptClient = MyGpt{
	C: NewGpr35(),
}

type MyGpt struct {
	C *openaigo.Client
}

func NewGpr35() *openaigo.Client {
	cfg := config.LoadConfig()
	return openaigo.NewClient(cfg.ApiKey)
}

func (c *MyGpt) Gpt3P5(req string) (string, error) {
	request := openaigo.ChatCompletionRequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []openaigo.ChatMessage{
			{Role: "user", Content: req},
		},
	}
	ctx := context.Background()
	rsp, err := c.C.Chat(ctx, request)
	if err != nil {
		abing_logger.SugarLogger.Errorf("gpt client chat erorr:%+v", err)
		return "", errors.New("请求GTP ERROR")
	}

	if len(rsp.Choices) == 0 || rsp.Choices[0].Message.Content == "" {
		return "", errors.New("请求GTP ERROR")
	}

	return rsp.Choices[0].Message.Content, nil
}
