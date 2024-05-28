package oss

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts "github.com/alibabacloud-go/sts-20150401/v2/client"
	teaUtils "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/fisschl/fiber/utils"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

var stsClient *sts.Client

func init() {
	config := &openapi.Config{
		AccessKeyId:     tea.String(os.Getenv("OSS_ACCESS_KEY_ID")),
		AccessKeySecret: tea.String(os.Getenv("OSS_ACCESS_KEY_SECRET")),
	}
	config.Endpoint = tea.String("sts.cn-shanghai.aliyuncs.com")
	var err error
	stsClient, err = sts.NewClient(config)
	if err != nil {
		log.Panicln(err)
	}
}

func stsHandler(ctx *fiber.Ctx) error {
	user := utils.UserId(ctx)
	if user == "" {
		return ctx.SendStatus(401)
	}
	Action := []string{"oss:GetObject", "oss:PutObject"}
	Resource := []string{
		fmt.Sprintf("acs:oss:*:*:fisschl/home/%s/*", user),
	}
	Policy, err := json.Marshal(fiber.Map{
		"Version": "1",
		"Statement": []fiber.Map{
			{
				"Effect":   "Allow",
				"Action":   Action,
				"Resource": Resource,
			},
		},
	})
	if err != nil {
		return err
	}
	assumeRoleRequest := &sts.AssumeRoleRequest{
		RoleArn:         tea.String(os.Getenv("OSS_STS_ROLE_ARN")),
		RoleSessionName: tea.String("upload"),
		DurationSeconds: tea.Int64(3000),
		Policy:          tea.String(string(Policy)),
	}
	runtime := &teaUtils.RuntimeOptions{}
	response, err := stsClient.AssumeRoleWithOptions(assumeRoleRequest, runtime)
	if err != nil {
		return err
	}
	return ctx.JSON(response.Body.Credentials)
}
