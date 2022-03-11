package api

import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"

	"github.com/briandowns/spinner"
	httpClient "github.com/abdfnx/resto/client"
)

func GetSCLatest() string {
	url := "https://api.secman.dev/latest-sc"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Errorf("Error creating request: %s", err.Error())
	}

	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Suffix = " 🔍 Getting latest release..."
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
