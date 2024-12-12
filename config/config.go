package config

import "github.com/spf13/viper"

type Config struct {
	Environment string `mapstructure:"environment"`
	Databases   []struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		Type     string `mapstructure:"type"`
	} `mapstructure:"databases"`
	Logging struct {
		Level  string `mapstructure:"level"`
		Output string `mapstructure:"output"`
	} `mapstructure:"logging"`
}

func NewConfig(path, name string) (c *Config, err error) {
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

	return c, nil

}
