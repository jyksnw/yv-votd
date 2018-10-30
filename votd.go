package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	apiURL         = "https://developers.youversionapi.com/1.0/verse_of_the_day/"
	cacheDir       = ".votd"
	defaultTimeout = 5 * time.Second
)

var (
	apiToken   = os.Getenv("YOUVERSION_VOTD_TOKEN")
	versionID  = os.Getenv("YOUVERSION_VOTD_VERSION")
	fileName   = time.Now().Local().Format("20060102")
	today      = time.Now().Local().YearDay()
	votdReqURL = apiURL + strconv.Itoa(today) + "?version_id=" + versionID
)

type votd struct {
	Day   int `json:"day"`
	Image struct {
		Attribution string `json:"attribution"`
		URL         string `json:"url"`
	} `json:"image"`
	Verse struct {
		HTML           string   `json:"html"`
		HumanReference string   `json:"human_reference"`
		Text           string   `json:"text"`
		URL            string   `json:"url"`
		Usfms          []string `json:"usfms"`
	} `json:"verse"`
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

func checkError(err error) {
	if err != nil {
		// Print to STDERR the error
		println(err.Error())

		// Print a user friendly message to STDOUT
		fmt.Println("Could not load verse of the day")
		os.Exit(0)
	}
}

func main() {
	if apiToken == "" {
		checkError(errors.New("Missing API Token"))
	}

	if versionID == "" {
		versionID = "1"
	}

	ex, err := os.Executable()
	if err != nil {
		checkError(err)
	}

	exPath := filepath.Dir(ex)
	cache := exPath + string(os.PathSeparator) + cacheDir
	exist, err := exists(cache)
	if err != nil {
		checkError(err)
	}

	if !exist {
		err = os.Mkdir(cache, os.ModePerm)
		if err != nil {
			checkError(err)
		}
	}

	cacheFile := cache + string(os.PathSeparator) + fileName
	exist, err = exists(cacheFile)
	if err != nil {
		checkError(err)
	}

	var votd votd

	if exist {
		data, err := ioutil.ReadFile(cacheFile)
		if err != nil {
			checkError(err)
		}

		err = json.Unmarshal(data, &votd)
		if err != nil {
			checkError(err)
		}
	} else {
		var netTransport = &http.Transport{
			Dial: (&net.Dialer{
				Timeout: defaultTimeout,
			}).Dial,
			TLSHandshakeTimeout: defaultTimeout,
		}

		var netClient = &http.Client{
			Timeout:   defaultTimeout,
			Transport: netTransport,
		}

		request, err := http.NewRequest("GET", votdReqURL, nil)
		if err != nil {
			checkError(err)
		}

		request.Header.Add("accept", "application/json")
		request.Header.Add("x-youversion-developer-token", apiToken)

		response, err := netClient.Do(request)
		if err != nil {
			checkError(err)
		}

		if response.StatusCode < 200 || response.StatusCode >= 400 {
			checkError(errors.New(response.Status))
		}

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			checkError(err)
		}

		err = json.Unmarshal(data, &votd)
		if err != nil {
			checkError(err)
		}

		// Ignore any errors here for now since this is only caching the response
		_ = ioutil.WriteFile(cacheFile, data, os.ModePerm)
	}

	fmt.Println(votd.Verse.HumanReference)
	fmt.Println()
	fmt.Println(votd.Verse.Text)
}
