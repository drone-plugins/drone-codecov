//go:build download
// +build download

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

const baseUrl = "https://uploader.codecov.io/latest/%s/codecov"

func main() {
	if runtime.GOARCH != "amd64" {
		fmt.Println("Only 'amd64' is supported!")
		os.Exit(1)
	}

	var (
		downloadURL string
		fileName    string
	)
	switch runtime.GOOS {
	case "darwin":
		downloadURL = fmt.Sprintf(baseUrl, "macos")
		fileName = "codecov"
	case "linux":
		downloadURL = fmt.Sprintf(baseUrl, "linux")
		fileName = "codecov"
	case "windows":
		downloadURL = fmt.Sprintf(baseUrl+".exe", "windows")
		fileName = "codecov.exe"
	default:
		fmt.Println("Only \"darwin\", \"linux\" & \"windows\" are supported!")
		os.Exit(1)
	}

	fmt.Printf("Download '%s' ...\n", downloadURL)

	client := &http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(downloadURL)
	checkErr(err)
	defer resp.Body.Close()

	file, err := os.Create(fileName)
	checkErr(err)
	defer file.Close()

	io.Copy(file, resp.Body)
	checkErr(file.Chmod(0o555))
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}
