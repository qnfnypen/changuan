package pkg

import (
	"fmt"
	"os"

	"github.com/qnfnypen/changuan/internal/domain/param"
	"github.com/qnfnypen/changuan/pkg/opfile"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	initialConf()
	initialLogger()
}

// initialConf 初始化程序配置信息
func initialConf() {
	confName := "conf_"
	switch param.ConfType {
	case "test":
		confName += "test"
	case "debug":
		confName += "debug"
	case "release":
		confName += "release"
	default:
		fmt.Fprintf(os.Stderr, "tag mode error: %s", "只支持: test debug release三种模式")
		os.Exit(1)
	}

	viper.SetConfigName(confName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("config/")
	viper.AddConfigPath(param.LocalConfFilePath)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "config file path or name error: %v", err)
		os.Exit(1)
	}
}

// initialLogger 初始化程序日志系统
func initialLogger() {
	fp := viper.GetString("Logger.ErrorLog.FilePath")
	f, err := opfile.CreateFileWithTimeStamp(fp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create error log file error: %v", err)
		os.Exit(1)
	}

	zerolog.TimeFieldFormat = "2006/01/02 15:04:05"
	logLevel := map[string]zerolog.Level{
		"trace": -1,
		"debug": 0,
		"info":  1,
		"warn":  2,
		"error": 3,
		"fatal": 4,
		"panic": 5,
	}
	zerolog.SetGlobalLevel(logLevel[viper.GetString("Logger.ErrorLog.Level")])

	log.Logger = zerolog.New(f).With().Timestamp().Caller().Logger()
}
