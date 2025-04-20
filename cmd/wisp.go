package main

import (
	"fmt"
	"os"
	"wisp/internal/config"
	"wisp/internal/wispweb"
)

// global variables
var wispServerConfig *config.ServerConfig

func main() {
	// service greeting, always to STDOUT
	fmt.Print(wispweb.WispLogo)

	// initialize main error object for handling
	var err error

	// intitialize the main server object
	var srv wispweb.WispServer

	// configuration file location
	configFilePath := "./config.yaml"

	// load configuration
	wispServerConfig, err = config.LoadConfigFromYAML(configFilePath)

	if err != nil { // we had a problem loading and/or parsing configuration
		os.Exit(1) // logger will handle messages
	}

	// assign loaded configs
	srv.Config = wispServerConfig

	// run server listener
	err = wispweb.Run(srv)
	if err != nil { // we catch server startup issues
		os.Exit(1)
	}
}
