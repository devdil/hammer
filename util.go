package main

import (
	"encoding/json"
	"errors"
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
