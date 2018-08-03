package configs

import (
	"log"

	"github.com/adharmonics/llog"
	"github.com/spf13/viper"
)

type appConfig struct {
	AppVersion string
	GRPCPort   int
	RestPost   int
	LogEnable  bool
	LogLevel   string
}

var Config appConfig

func init() {
	if err := initDefault(); err != nil {
		log.Printf("error initializing default config: %s", err)
	} else {
		log.Print("configs successfully initialized")
	}
}

func initDefault() error {
	viper.SetEnvPrefix("FT")
	viper.AutomaticEnv()
	viper.SetDefault("app_version", "undefined")
	viper.SetDefault("rest_port", 5500)
	viper.SetDefault("grpc_port", 5501)
	viper.SetDefault("log_enable", true)
	viper.SetDefault("log_level", "DebugLevel")

	Config = appConfig{
		AppVersion: viper.GetString("app_version"),
		GRPCPort:   viper.GetInt("grpc_port"),
		RestPost:   viper.GetInt("rest_port"),
		LogEnable:  viper.GetBool("log_enable"),
		LogLevel:   viper.GetString("log_level"),
	}

	llog.ConfigLlog(llog.Options{
		Enabled:    Config.LogEnable,
		AppVersion: Config.AppVersion,
		LogLevel:   Config.LogLevel,
	})

	return nil
}
