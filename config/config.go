package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"os"
	"path/filepath"
)

var Conf *Config

type Config struct {
	Env       string
	Port      int
	PprofHost string `mapstructure:"pprof_host"`
	PprofPort int    `mapstructure:"pprof_port"`
}

func InitConfig() error {
	exePwd, err := os.Executable()
	if err != nil {
		return errors.New("get path failed")
	}
	configPath := filepath.Dir(exePwd)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath("./config")
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	slog.Info(fmt.Sprintf("config file used: %v", viper.ConfigFileUsed()))
	err = viper.Unmarshal(&Conf)
	if err != nil {
		return errors.New("load config failed")
	}
	return nil
}
