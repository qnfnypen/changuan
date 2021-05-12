package main

import (
	"crypto/tls"
	"crypto/x509"
	"net"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// generateRPCServer 生成rpc服务端
func generateRPCServer() {
	// 加载证书
	cert, err := tls.LoadX509KeyPair(viper.GetString("RPC.CertFile"), viper.GetString("RPC.KeyFile"))
	if err != nil {
		log.Fatal().Str("tls.LoadX509KeyPair error", err.Error()).Msg("RPC服务端加载证书失败")
	}

	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(viper.GetString("RPC.CAFile"))
	if err != nil {
		log.Fatal().Str("os.ReadFile error", err.Error()).Msg("读取ca证书文件失败")
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal().Str("certPool.AppendCertsFromPEM error", "append cert to pool false").Msg("添加证书文件到池子错误")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	// 创建rpc服务端
	server := grpc.NewServer(grpc.Creds(c))
	// 注册rpc服务

	// 开启rpc服务端
	lis, err := net.Listen(viper.GetString("RPC.Network"), viper.GetString("RPC.Address"))
	if err != nil {
		log.Fatal().Str("net.Listen error",err.Error()).Msg("net监听失败")
	}
	if err = server.Serve(lis);err != nil {
		log.Fatal().Str("server.Serve error",err.Error()).Msg("开启rpc服务端失败")
	}
}
