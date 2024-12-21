package config

import "github.com/spf13/viper"

type BaseConfig struct {
	Environment string           `mapstructure:"environment"`
	AppName     string           `mapstructure:"app_name"`
	Databases   []DatabaseConfig `mapstructure:"databases"`
	Storages    []StorageConfig  `mapstructure:"storages"`

	Logging LoggingConfig `mapstructure:"logging"`
}

type RestConfig struct {
	BaseConfig
	Server struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		BasePath string `mapstructure:"base_path"`
	} `mapstructure:"server"`
}

type GrpcConfig struct {
	BaseConfig
	Server struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
}

func NewConfig[C any](path, name string) (c *C, err error) {
	c = new(C)
	vp := viper.New()
	vp.AddConfigPath(path)
	vp.SetConfigName(name)
	vp.SetConfigType("yaml")
	err = vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	vp.Unmarshal(c)

	return c, nil

}
