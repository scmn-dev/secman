package api

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"

	"github.com/briandowns/spinner"
	httpClient "github.com/abdfnx/resto/client"
)

func GetLatest(product string, isChecker bool) string {
	p := ""

	if product == "secman-cli" {
		p = "latest"
	} else if product == "secman-core" {
		p = "latest-core"
	} else if product == "sc" {
		p = "latest-sc"
	}

	url := "https://api.secman.dev/" + p

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Printf("Error creating request: %s \n", err.Error())
		os.Exit(0)
	}

	suffix := " 🔍 Checking for updates..."

	if isChecker {
		suffix = " 🔍 Requesting..."
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = suffix
	s.Start()

	client := httpClient.HttpClient()
	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error sending request: %s", err.Error())
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Error reading response: %s", err.Error())
	}

	s.Stop()

	return string(b)
}
