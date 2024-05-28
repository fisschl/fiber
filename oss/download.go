package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gofiber/fiber/v2"
	"log"
)

var ObjectStore *oss.Client

func init() {
	// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量 OSS_ACCESS_KEY_ID 和 OSS_ACCESS_KEY_SECRET
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		log.Panicln(err)
	}
	// 创建 ClientOSS 实例。
	ObjectStore, err = oss.New(
		"https://oss-cn-shanghai.aliyuncs.com",
		"",
		"",
		oss.SetCredentialsProvider(&provider))
	if err != nil {
		log.Panicln(err)
	}
}

func downloadHandler(ctx *fiber.Ctx) error {
	objectName := ctx.Query("key")
	// 将Object下载到本地文件，并保存到指定的本地路径中。如果指定的本地文件存在会覆盖，不存在则新建。
	bucket, err := ObjectStore.Bucket("fisschl")
	if err != nil {
		return err
	}
	// 生成用于下载的签名URL，并指定签名URL的有效时间为60秒。
	signedURL, err := bucket.SignURL(objectName, oss.HTTPGet, 60)
	if err != nil {
		return err
	}
	return ctx.Redirect(signedURL)
}
