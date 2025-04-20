package wispweb

import (
	"wisp/internal/config"
)

type WispServer struct {
	Config *config.ServerConfig
}

type RequestHeader struct {
	Name   string
	Values []string
}

type EnvironmentVariable struct {
	Name  string
	Value string
}

type ReportData struct {
	HostName string
	Headers  []RequestHeader
	EnvVars  []EnvironmentVariable
}

const (
	WispLogo = `

	░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓███████▓▒░▒▓███████▓▒░  
	░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
	░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░░▒▓█▓▒░ 
	░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░░▒▓██████▓▒░░▒▓███████▓▒░  
	░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░▒▓█▓▒░        
	░▒▓█▓▒░░▒▓█▓▒░░▒▓█▓▒░▒▓█▓▒░      ░▒▓█▓▒░▒▓█▓▒░        
	 ░▒▓█████████████▓▒░░▒▓█▓▒░▒▓███████▓▒░░▒▓█▓▒░        
															  
`
)
