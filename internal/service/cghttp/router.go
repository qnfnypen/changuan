package cghttp

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	// 初始化错误日志引起和配置文件
	_ "github.com/qnfnypen/changuan/internal/pkg"
	"github.com/qnfnypen/changuan/internal/service/cghttp/middleware"
	"github.com/qnfnypen/changuan/pkg/opfile"
)

// GenerateEngine 生成http端路由引擎
func GenerateEngine() *gin.Engine {
	// 设置gin日志和运行模式
	fp := viper.GetString("Logger.GinLogFilePath")
	f, err := opfile.CreateFileWithTimeStamp(fp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create gin log file error: %v", err)
		os.Exit(1)
	}
	gin.DefaultWriter = io.MultiWriter(f)
	gin.SetMode(viper.GetString("HTTP.Mode"))

	// 生成http路由引擎
	r := gin.Default()

	// 判断是否开启跨域
	if viper.GetBool("HTTP.CORS.Enable") {
		r.Use(middleware.CORS())
	}
	// 配置Swagger接口文档

	return r
}
