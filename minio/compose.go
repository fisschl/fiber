package minio

import (
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

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
