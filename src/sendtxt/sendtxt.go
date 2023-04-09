package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <URL> <file>")
		os.Exit(1)
	}

	url := os.Args[1]
	filePath := os.Args[2]

	// テキストファイルを読み込む
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	// HTTP POSTリクエストを作成する
	resp, err := http.Post(url, "text/plain", bytes.NewBuffer(fileContent))
	if err != nil {
		fmt.Printf("Error sending POST request: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// レスポンスを表示する
	fmt.Printf("Status: %s\n", resp.Status)
}
