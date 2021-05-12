package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/qnfnypen/changuan/internal/pkg"
	"github.com/qnfnypen/changuan/internal/service/cghttp"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// generateHTTPServer 生成http服务端
func generateHTTPServer() {
	server := &http.Server{
		Addr:           viper.GetString("HTTP.Address"),
		Handler:        cghttp.GenerateEngine(),
		ReadTimeout:    viper.GetDuration("HTTP.ReadTimeout") * time.Second,
		WriteTimeout:   viper.GetDuration("HTTP.WriteTimeout") * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 捕获用户行为，关闭http服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Kill, os.Interrupt)
	go func() {
		<-quit

		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatal().Str("server.Shutdown error", err.Error()).Msg("HTTP服务端退出失败")
		}
	}()

	// 开启http服务端
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Debug().Str("server.ListenAndServe error", err.Error()).Msg("HTTP服务端由于请求而关闭")
		} else {
			log.Warn().Str("server.ListenAndServe error", err.Error()).Msg("HTTP服务端由于意外而关闭")
		}
	}
}
