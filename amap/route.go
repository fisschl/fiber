package amap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gookit/config/v2"
	"net/url"
	"strings"
)

func Register(router fiber.Router) {

	//   location /_AMapService/v4/map/styles {
	//    set $args "$args&jscode=你的安全密钥";
	//    proxy_pass https://webapi.amap.com/v4/map/styles;
	//  }

	router.Get("/_AMapService/v4/map/styles", func(ctx *fiber.Ctx) error {
		path := strings.Replace(ctx.OriginalURL(), "/_AMapService/", "/", 1)
		target, err := url.Parse("https://webapi.amap.com" + path)
		if err != nil {
			return err
		}
		query := target.Query()
		query.Set("jscode", config.String("AMAP_SECRET"))
		target.RawQuery = query.Encode()
		return proxy.Do(ctx, target.String())
	})

	//  # 海外地图服务代理
	//  location /_AMapService/v3/vectormap {
	//    set $args "$args&jscode=你的安全密钥";
	//    proxy_pass https://fmap01.amap.com/v3/vectormap;
	//  }

	router.Get("/_AMapService/v3/vectormap", func(ctx *fiber.Ctx) error {
		path := strings.Replace(ctx.OriginalURL(), "/_AMapService/", "/", 1)
		target, err := url.Parse("https://fmap01.amap.com" + path)
		if err != nil {
			return err
		}
		query := target.Query()
		query.Set("jscode", config.String("AMAP_SECRET"))
		target.RawQuery = query.Encode()
		return proxy.Do(ctx, target.String())
	})

	//  # Web服务API 代理
	//  location /_AMapService/ {
	//    set $args "$args&jscode=你的安全密钥";
	//    proxy_pass https://restapi.amap.com/;
	//  }

	router.Get("/_AMapService/*", func(ctx *fiber.Ctx) error {
		path := strings.Replace(ctx.OriginalURL(), "/_AMapService/", "/", 1)
		target, err := url.Parse("https://restapi.amap.com" + path)
		if err != nil {
			return err
		}
		query := target.Query()
		query.Set("jscode", config.String("AMAP_SECRET"))
		target.RawQuery = query.Encode()
		return proxy.Do(ctx, target.String())
	})

}
