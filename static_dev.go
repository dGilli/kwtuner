//+build dev
//go:build dev
// +build dev

package main

import (
	"fmt"
	"net/http"
	"os"
)

func public() http.Handler {
	fmt.Println("building static files for development")
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}

func prompt() string {
    filePath := "prompt.txt"

    content, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Convert byte slice to string
    fileAsString := string(content)

    return fileAsString
}

func logger() *services.LoggerDevConfig {
    return *services.LoggerDevConfig
}
