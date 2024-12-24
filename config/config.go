package config

import (
	"github.com/spf13/viper"
	appcontext "kotakemail.id/pkg/context"
)

type Config struct {
	Rest        RestConfig       `mapstructure:"rest"`
	Grpc        GrpcConfig       `mapstructure:"grpc"`
	Environment string           `mapstructure:"environment"`
	AppName     string           `mapstructure:"app_name"`
	Databases   []DatabaseConfig `mapstructure:"databases"`
	Storages    []StorageConfig  `mapstructure:"storages"`
	Logging     LoggingConfig    `mapstructure:"logging"`
}

type GrpcServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type GrpcClientConfig struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type GrpcConfig struct {
	Server  GrpcServerConfig   `mapstructure:"server"`
	Clients []GrpcClientConfig `mapstructure:"clients"`
}

func NewConfig(ctx *appcontext.AppContext, path, name string) (c *Config, err error) {
	c = &Config{}
	vp := viper.New()
	vp.AddConfigPath(path)
	vp.SetConfigName(name)
	vp.SetConfigType("yaml")
	err = vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	vp.Unmarshal(c)

	ctx.Set(appcontext.AppNameKey, c.AppName)
	ctx.Set(appcontext.EnvironmentKey, c.Environment)

	return c, nil

}
