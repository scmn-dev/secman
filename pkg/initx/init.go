package initx

import (
	"os"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/abdfnx/tran/dfs"
)

func Init() {
	var err error

	homeDir, err := dfs.GetHomeDirectory()

	if err != nil {
		log.Fatal(err)
	}

	err = dfs.CreateDirectory(filepath.Join(homeDir, ".secman"))
	if err != nil {
		log.Fatal(err)
	}

	viper.AddConfigPath("$HOME/.secman")
	viper.SetConfigName("secman")
	viper.SetConfigType("json")

	// Setup config defaults.
	viper.SetDefault("config.name", "")
	viper.SetDefault("config.secret", "")
	viper.SetDefault("config.user", "")
	viper.SetDefault("data.access_token", "")
	viper.SetDefault("data.master_password_hash", "")
	viper.SetDefault("data.refresh_token", "")
	viper.SetDefault("data.transmission_key", "")

	if err := viper.SafeWriteConfig(); err != nil {
		if os.IsNotExist(err) {
			err = viper.WriteConfig()

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal(err)
		}
	}
}
