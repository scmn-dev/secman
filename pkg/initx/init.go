package initx

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/gosh"
	"github.com/abdfnx/tran/dfs"
	gapi "github.com/scmn-dev/get-latest/api"
	"github.com/scmn-dev/secman/constants"
	"github.com/spf13/viper"
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

	secmanDirPath := ""

	if runtime.GOOS == "windows" {
		secmanDirPath = `$HOME\\.secman`
	} else {
		secmanDirPath = `$HOME/.secman`
	}

	viper.AddConfigPath(secmanDirPath)
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

	// Get SMUI
	smuiLatest := gapi.LatestWithArgs("david-tomson/smui", "")
	url := "https://github.com/david-tomson/smui/releases/download/" + smuiLatest + "/smui.zip"

	uCmd := fmt.Sprintf(`
		if ! [ -d %s/ui ]; then
			wget %s
			sudo chmod 755 smui.zip
			unzip -qq smui.zip
			mv ui %s/ui
			rm smui.zip
		fi
	`, constants.DotSecmanPath, url, constants.DotSecmanPath)

	wCmd := fmt.Sprintf(`
		if (-not (Test-Path -path %s/ui)) {
			Invoke-WebRequest %s -outfile smui.zip
			Expand-Archive smui.zip
			Move-Item -Path smui/ui -Destination .
			Move-Item -Path ui -Destination %s
			Remove-Item smui* -Recurse -Force
		}
	`, constants.DotSecmanPath, fmt.Sprintf("\"%s\"", url), constants.DotSecmanPath)

	gosh.RunMulti(uCmd, wCmd)
}
