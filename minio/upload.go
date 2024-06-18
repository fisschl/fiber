package minio

import (
	"github.com/fisschl/fiber/config"
	"github.com/fisschl/fiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"strings"
	"time"
)

var Minio *minio.Client

func init() {
	option := &minio.Options{
		Creds:  credentials.NewStaticV4(config.GetEnv("MINIO_ACCESS_KEY_ID"), config.GetEnv("MINIO_ACCESS_KEY_SECRET"), ""),
		Secure: true,
	}
	minioClient, err := minio.New("oss.bronya.world", option)
	if err != nil {
		log.Fatalln(err)
	}
	Minio = minioClient
}

func privateObjectName(ctx *fiber.Ctx) (string, error) {
	objectName := ctx.Query("key")
	token := utils.AuthToken(ctx)
	if token == "" {
		return "", ctx.SendStatus(fiber.StatusUnauthorized)
	}
	user, err := utils.Rdb.HGet(ctx.Context(), token, "user").Result()
	if err != nil {
		return "", ctx.SendStatus(fiber.StatusUnauthorized)
	}
	if !strings.HasPrefix(objectName, "/"+user+"/") {
		return "", ctx.SendStatus(fiber.StatusForbidden)
	}
	return objectName, err
}

func HandleUpload(ctx *fiber.Ctx) error {
	objectName, err := privateObjectName(ctx)
	if err != nil {
		return err
	}
	value, err := Minio.PresignedPutObject(ctx.Context(),
		"home",
		objectName,
		time.Second*24*60*60)
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"url": value.String(),
	})
}
