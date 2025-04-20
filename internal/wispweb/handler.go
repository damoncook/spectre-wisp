package wispweb

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"text/template"
	"wisp/internal/logging"
)

// global server variable
var srv WispServer

func Run(s WispServer) error {
	// initialize error object for web request error handling
	var err error

	// copy the provided server instance object to our global instance
	srv = s

	// setup logging for service
	var logSrv = logging.InitLogger("wispweb", srv.Config.LogType)

	// set up the primary web handling function for web requests
	http.HandleFunc("/", contextHandler)

	logging.PrintGreen(logSrv, "✓ launching listener on "+srv.Config.WebAddr+":"+srv.Config.WebPort)

	// open the web service
	err = http.ListenAndServe(srv.Config.WebAddr+":"+srv.Config.WebPort, nil)
	if err != nil {
		logging.PrintErr(logSrv, err)
		return err
	}

	return nil
}

func contextHandler(w http.ResponseWriter, req *http.Request) {

	// setup logging to log each web request
	var logWeb = logging.InitLogger("wispweb", srv.Config.LogType)

	//tmpl, err := template.ParseFiles("./static/html/templates/index.html")
	tmpl := template.Must(template.ParseFiles("./static/html/templates/index.html"))

	var reportData ReportData
	reportData.HostName, _ = os.Hostname()
	reportData.Headers = reportHeaders(req, logWeb)
	reportData.EnvVars = reportEnvironmentVariables(logWeb)

	switch req.RequestURI {
	case "/", "/index.html":
		tmpl.Execute(w, reportData)
	default:
		http.ServeFile(w, req, "./static/html"+req.URL.Path)
	}
}

func reportHeaders(req *http.Request, log *log.Logger) []RequestHeader {
	var requestHeaders []RequestHeader

	log.Println("request received from", req.RemoteAddr+":", req.RequestURI)

	var headerReport string
	headerReport += "  Request headers received from " + req.RemoteAddr + ":" + "\n"
	for name, values := range req.Header {
		var header RequestHeader
		header.Name = name
		header.Values = values

		// log each header value before returning

		headerReport += "    [" + header.Name + "] \n"
		for _, value := range header.Values {
			headerReport += "      " + value + "\n"
		}

		requestHeaders = append(requestHeaders, header)
	}

	log.Println(headerReport)

	return requestHeaders
}

func reportEnvironmentVariables(log *log.Logger) []EnvironmentVariable {
	regEnvVarName := regexp.MustCompile(`=.*`)
	regEnvVarValue := regexp.MustCompile(`.*=`)
	var environmentVariables []EnvironmentVariable

	for _, entry := range os.Environ() {
		var envVariable EnvironmentVariable
		envVariable.Name = regEnvVarName.ReplaceAllString(entry, "")
		envVariable.Value = regEnvVarValue.ReplaceAllString(entry, "")
		environmentVariables = append(environmentVariables, envVariable)
	}

	return environmentVariables
}
