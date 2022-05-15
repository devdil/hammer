package main

import (
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type httpRequest struct {
	isSecure    bool
	url         string
	headers     string
	requestType string
	reqCtxt     requestContext
}

type requestContext struct {
	result chan httpResponse
	wg     *sync.WaitGroup
}

type httpResponse struct {
	statusCode int
	timeTook   int
}

func (h httpRequest) Get() {
	// notify when its done, via waitgroup
	defer h.reqCtxt.wg.Done()
	startTime := time.Now().UnixMilli()
	resp, err := http.Get(h.url)
	if err != nil {
		log.Fatal("Something went wrong while fetching ", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Something went wrong while closing the get request ", err)
		}
	}(resp.Body)
	timeItTook := time.Now().UnixMilli() - startTime
	res := httpResponse{statusCode: resp.StatusCode, timeTook: int(timeItTook)}
	h.reqCtxt.result <- res
}
