package dashscope

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"os"
)

func HandleCompletions(ctx *fiber.Ctx) error {
	req, err := http.NewRequest("POST", "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions", bytes.NewReader(ctx.Body()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("DASHSCOPE_API_KEY")))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	ctx.Set("Content-Type", resp.Header.Get("Content-Type"))
	ctx.Set("Cache-Control", "no-cache")
	ctx.Set("Connection", "keep-alive")
	ctx.Set("Transfer-Encoding", resp.Header.Get("Transfer-Encoding"))
	return ctx.Status(resp.StatusCode).SendStream(resp.Body)
}
