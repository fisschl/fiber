package minio

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"time"
)

func HandleSlice(ctx *fiber.Ctx) error {
	tag := ctx.Query("slice")
	info, err := Minio.StatObject(ctx.Context(),
		"temp",
		tag,
		minio.StatObjectOptions{})
	if err == nil {
		return ctx.JSON(info)
	}
	// 该分片不存在，获取上传地址
	value, err := Minio.PresignedPutObject(ctx.Context(),
		"temp",
		tag,
		time.Second*24*60*60)
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"url":    value.String(),
		"status": "not_exist",
	})
}
