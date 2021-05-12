package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	// 初始化错误日志引起和配置文件
	_ "github.com/qnfnypen/changuan/internal/pkg"
)

// CORS 跨域中间件
func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     viper.GetStringSlice("HTTP.CORS.AllowOrigins"),
		AllowMethods:     viper.GetStringSlice("HTTP.CORS.AllowMethods"),
		AllowHeaders:     viper.GetStringSlice("HTTP.CORS.AllowHeaders"),
		AllowCredentials: viper.GetBool("HTTP.CORS.AllowCredentials"),
		MaxAge:           viper.GetDuration("HTTP.CORS.MaxAge") * time.Second,
	})
}
