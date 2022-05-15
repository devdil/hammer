package main

type GenericLoad interface {
	Start()
	GenerateReport()
	Init()
}

type LoadRequest struct {
	reqContext requestContext
	loader     GenericLoad
}

func (loadReq *LoadRequest) Initialize() {
	protocol := loadReq.reqContext.commandLineParams.protocol

	if protocol == "http" {
		loadReq.loader = &httpLoad{requestContext: loadReq.reqContext}

	}

	loadReq.loader.Init()
}

func (loadReq LoadRequest) Start() {
	// loop until we have made all the requests
	loadReq.loader.Start()
}

func (loadReq LoadRequest) GenerateReport() {
	loadReq.loader.GenerateReport()
}
