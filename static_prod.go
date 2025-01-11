//go:build !dev
// +build !dev

package main

import (
	"embed"
	"net/http"
)

//go:embed public
var publicFS embed.FS

//go:embed prompt.txt
var promptStr string

func public() http.Handler {
	return http.FileServerFS(publicFS)
}

func prompt() string {
    return promptStr
}
