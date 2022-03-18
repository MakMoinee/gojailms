package config

import (
	"github.com/MakMoinee/gojailms/internal/common"
	"github.com/spf13/viper"
)

var Registry *viper.Viper

// Set reading the configurations from settings.yaml
func Set() {
	var err error
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Registry = viper.GetViper()

	common.SERVER_PORT = Registry.GetString("SERVER_PORT")
	common.SERVER_ENABLE_PROFILING = Registry.GetBool("SERVER_PROFILING")
	common.DB_NAME = Registry.GetString("DB_NAME")
	common.DB_DRIVER = Registry.GetString("DB_DRIVER")
	common.MYSQL_USERNAME = Registry.GetString("MYSQL_USERNAME")
	common.MYSQL_PASSWORD = Registry.GetString("MYSQL_PASSWORD")
	common.CONNECTION_STRING = Registry.GetString("CONNECTION_STRING")
}
