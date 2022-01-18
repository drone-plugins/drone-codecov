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

	resp, err := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}.Get(downloadURL)
	checkErr(err)
	defer resp.Body.Close()

	file, err := os.Create(fileName)
	checkErr(err)
	defer file.Close()
	checkErr(file.Chmod(0o444))

	io.Copy(file, resp.Body)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}
