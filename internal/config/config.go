package config

import (
	"bytes"
	"io/ioutil"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/abdfnx/tran/dfs"
)

var (
	homeDir, _ = dfs.GetHomeDirectory()
	secmanConfigPath = filepath.Join(homeDir, ".secman", "secman.json")
	secmanConfig, _ = ioutil.ReadFile(secmanConfigPath)
)


func Config(obj string) string {
	viper.SetConfigType("json")

	viper.ReadConfig(bytes.NewBuffer(secmanConfig))

	return viper.Get(obj).(string)
}
