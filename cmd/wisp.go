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
	// initialize main error object for handling
	var err error

	// intitialize the main server object
	var srv wispweb.WispServer

	// configuration file location
	configFilePath := "./config.yaml"

	// load configuration
	wispServerConfig, err = config.LoadConfigFromYAML(configFilePath)

	if err != nil {
		fmt.Println("! Could not load configuration file:", err)
		fmt.Println("  ... please ensure you have a config.yaml file in the " +
			"server executable directory.")
		os.Exit(1)
	}
	fmt.Println("* Loaded configuration from", configFilePath)

	srv.Config = wispServerConfig

	err = wispweb.Run(srv)
	if err != nil {
		fmt.Println("! Couldn't launch web service:", err)
		os.Exit(1)
	}

}
