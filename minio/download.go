package minio

import (
	"github.com/gofiber/fiber/v2"
	"net/url"
	"time"
)

func HandleDownload(ctx *fiber.Ctx) error {
	objectName, err := privateObjectName(ctx)
	if err != nil {
		return err
	}
	value, err := Minio.PresignedGetObject(ctx.Context(),
		"home",
		objectName,
		time.Second*24*60*60,
		url.Values{},
	)
	if err != nil {
		return err
	}
	return ctx.Redirect(value.String())
}
