package main

import (
	"fmt"
	"time"
)

type LoadRequest struct {
	reqContext      requestContext

}

func (loadReq *LoadRequest) Initialize() {
	protocol := loadReq.reqContext.commandLineParams.protocol

	if protocol == 'http' {
		// report generator should reference http report generator
		// load
		loadReq.reportGenerator := HttpReportGenerator{}
	}

}

func (loadReq LoadRequest) Start() {
	// loop until we have made all the requests
	startTimeSeconds := time.Now().Second()
	for {
		currentTimeSeconds := time.Now().Second()
		secondsSinceStarted := currentTimeSeconds - startTimeSeconds
		if secondsSinceStarted == loadReq.reqContext.commandLineParams.dur {
			break
		}

		for i := 1; i <= loadReq.reqContext.commandLineParams.dur; i++ {
			wg.Add(1)
			go GetHomePage("http://diljitpr.net", responses, &wg)
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n Collecting results...")
	wg.Wait()
	close(responses)
	s.Stop()
}

func (loadReq LoadRequest) GenerateReport() {
	loadReq.reportGenerator.GenerateReport()
	// create a channel of n requests
	responses := make(chan result, requestsPerSeconds*duration)

	// call report generator
	httpReportGenerator := HttpReportGenerator{responses: responses}
	httpReportGenerator.GenerateReport()
}
