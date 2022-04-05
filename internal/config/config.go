package config

import (
	"bytes"

	"github.com/scmn-dev/secman/v6/constants"
	"github.com/spf13/viper"
)

func Config(obj string) string {
	viper.SetConfigType("json")

	err := viper.ReadConfig(bytes.NewBuffer(constants.SecmanConfig()))

	if err == nil {
		return viper.Get(obj).(string)
	}

	return ""
}
