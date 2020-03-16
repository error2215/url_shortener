package main

import (
	"github.com/error2215/url_shortener/server"
	_ "github.com/error2215/url_shortener/server/config"
)

func main() {
	server.Start()
}
