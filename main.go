package main

import (
	"laji/v1/config"
	"laji/v1/service"
)

var (
	cfgfile = "C:/app/www/packages/regenerate-management/regenerate-go/laji.yml"
)

func main() {
	config.Initialize(cfgfile)
	service.Start()
}
