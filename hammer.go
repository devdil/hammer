package main

import (
	"flag"
)

func main() {

	/*  Start of parsing command line parameters */
	rps := flag.Int("rps", 1, "Requests per second. Usage rps=1 where 1 is a positive integer.")
	dur := flag.Int("dur", 1, "Duration in seconds dur=1 where is 1 is a valid duration in seconds.")
	protocol := flag.String("prot", "http", "Protocol used http, tcp or udp.")
	url := flag.String("url", "url", "Target url for load testing.")
	methodType := flag.String("method", "get", "If protocol used is http, possible options are get, put, update, delete. Default is get.")
	headers := flag.String("headers", "{}", "headers in string format. example: `{'Content-Type' : 'application/json'`")
	/*  End of parsing command line parameters */
	flag.Parse()

	//TODO
	// validate cmd params

	cmdParams := cmdParams{rps: *rps, dur: *dur, protocol: *protocol, methodType: *methodType, url: *url, headers: *headers}

	// print welcome message
	PrintWelcomeMessage(cmdParams)

	// Initialize load engine and start
	reqContext := requestContext{commandLineParams: cmdParams}
	loadRequest := LoadRequest{reqContext: reqContext}
	loadRequest.Initialize()
	loadRequest.Start()
	loadRequest.GenerateReport()

}
