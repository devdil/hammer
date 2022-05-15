package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type httpRequest struct {
	isSecure    bool
	headers     map[string]string
	requestType string
	reqCtxt     requestContext
	wg          sync.WaitGroup
	responses   chan httpResponse
	httpClient  *http.Client
}

type httpResponse struct {
	statusCode int
	timeTook   int
}

func (h *httpRequest) IncrementRequestCount() {
	h.wg.Add(1)
}

func (h *httpRequest) Get() {
	// notify when its done, via waitgroup
	startTime := time.Now().UnixMilli()
	req, err := http.NewRequest("GET", h.reqCtxt.commandLineParams.url, nil)
	//for headerName, headerVal := range h.headers {
	//	req.Header.Add(headerName, headerVal)
	//}
	resp, err := h.httpClient.Do(req)
	if err != nil {
		log.Fatal("Something went wrong while fetching ", err)
	}
	defer resp.Body.Close()
	timeItTook := time.Now().UnixMilli() - startTime
	res := httpResponse{statusCode: resp.StatusCode, timeTook: int(timeItTook)}
	h.responses <- res
	h.wg.Done()
}

func (h *httpRequest) WaitForAllToComplete() {

}
