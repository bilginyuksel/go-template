package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Read reads the configuration file and returns a viper instance
func Read(filepath, prefix string, conf interface{}) error {
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return viper.Unmarshal(conf)
}
