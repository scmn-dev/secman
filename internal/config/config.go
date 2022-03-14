package config

import (
	"bytes"

	"github.com/spf13/viper"
	"github.com/scmn-dev/secman/constants"
)

func Config(obj string) string {
	viper.SetConfigType("json")

	err := viper.ReadConfig(bytes.NewBuffer(constants.SecmanConfig()))

	if err == nil {
		return viper.Get(obj).(string)
	}

	return ""
}
