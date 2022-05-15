package main

import (
	"fmt"
	"net/http"
	"time"
)

type httpLoad struct {
	httpReport     HttpReportGenerator
	requestContext requestContext
	httpResponses  chan httpResponse
	httpConnector  *httpRequest
}

func (h *httpLoad) Init() {
	// create a channel
	rps := h.requestContext.commandLineParams.rps
	dur := h.requestContext.commandLineParams.dur
	h.httpResponses = make(chan httpResponse, rps*dur)
	// pass ^ above to report generator
	h.httpReport = HttpReportGenerator{responses: h.httpResponses}
	// compose headers
	headersInMap, _ := ParseHeadersToMap(h.requestContext.commandLineParams.headers)
	httpClient := &http.Client{}
	h.httpConnector = &httpRequest{reqCtxt: h.requestContext, requestType: h.requestContext.commandLineParams.methodType, headers: headersInMap, responses: h.httpResponses, httpClient: httpClient}
}

func (h *httpLoad) Start() {

	switch methodType := h.requestContext.commandLineParams.methodType; methodType {
	case "GET":
		// loop until we have made all the requests
		startTimeSeconds := time.Now().Second()
		for {
			currentTimeSeconds := time.Now().Second()
			secondsSinceStarted := currentTimeSeconds - startTimeSeconds
			if secondsSinceStarted == h.requestContext.commandLineParams.dur {
				break
			}

			for i := 1; i <= h.requestContext.commandLineParams.rps; i++ {
				h.httpConnector.IncrementRequestCount()
				go h.httpConnector.Get()
			}
			time.Sleep(1 * time.Second)
		}
		h.WaitForAllToComplete()
		close(h.httpConnector.responses)

	default:
		fmt.Println("Http type not supported yet!")
	}

}

func (h *httpLoad) WaitForAllToComplete() {
	h.httpConnector.wg.Wait()
}

func (h httpLoad) GenerateReport() {
	h.httpReport.GenerateReport()
}
