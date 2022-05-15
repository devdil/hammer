package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

func ParseHeadersToMap(headers string) (map[string]string, error) {
	headersMap := map[string]string{}
	err := json.Unmarshal([]byte(headers), &headersMap)
	if err != nil {
		log.Println("Could not parse headers to map,", err)
		return nil, errors.New("could not parse headers into native map")
	}
	return headersMap, nil
}

func PrintWelcomeMessage(params cmdParams) {
	fmt.Println("                          \\`. \n.--------------.___________) \\\n|//////////////|___________[ ]\n`--------------'           ) (\n                           '-'")
	fmt.Println("\n█░█ ▄▀█ █▀▄▀█ █▀▄▀█ █▀▀ █▀█\n█▀█ █▀█ █░▀░█ █░▀░█ ██▄ █▀▄")
	fmt.Println("Hammer Load Test Tool ")
	fmt.Println("---------------------------------")
	fmt.Println("-----Parameters------------------")
	fmt.Println("Requests Per Second: ", params.rps)
	fmt.Println("Duration: ", params.dur)
	fmt.Println("---------------------------------")
	fmt.Println("Hammering ...")

}
