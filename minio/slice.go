package minio

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"time"
)

// HandleSlice 上传一个切片
func HandleSlice(ctx *fiber.Ctx) error {
	tag := ctx.Query("slice")
	info, err := Minio.StatObject(ctx.Context(),
		"temp",
		tag,
		minio.StatObjectOptions{})
	standard := time.Now().AddDate(0, 0, 7)
	if err == nil && !info.Expiration.Before(standard) {
		// 分片在7天内不会过期，返回分片信息
		return ctx.JSON(info)
	}
	// 该分片不存在，获取上传地址
	value, err := Minio.PresignedPutObject(ctx.Context(),
		"temp",
		tag,
		time.Hour*24)
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"url":    value.String(),
		"status": "not_exist",
	})
}

type composeRequest struct {
	List []string `json:"list"`
	Type string   `json:"type"`
}

func HandleCompose(ctx *fiber.Ctx) error {
	objectName, err := privateObjectName(ctx)
	if err != nil {
		return err
	}
	request := new(composeRequest)
	err = ctx.BodyParser(request)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	srcs := make([]minio.CopySrcOptions, 0)
	for _, tag := range request.List {
		srcs = append(srcs, minio.CopySrcOptions{
			Bucket: "temp",
			Object: tag,
		})
	}
	option := minio.CopyDestOptions{
		Bucket: "home",
		Object: objectName,
		UserMetadata: map[string]string{
			"Content-Type": request.Type,
		},
		ReplaceMetadata: true,
	}
	uploadInfo, err := Minio.ComposeObject(ctx.Context(), option, srcs...)
	if err != nil {
		return err
	}
	return ctx.JSON(uploadInfo)
}
