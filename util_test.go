package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseHeadersToMapShouldPassWithRightParams(t *testing.T) {
	testHeaders := `{"Content-Type" : "application/json"}`
	headersToMap, err := ParseHeadersToMap(testHeaders)

	assert.Nil(t, err)
	assert.Equal(t, headersToMap["Content-Type"], "application/json")

}

func TestParseHeadersToMapShouldFailWithIncorectJson(t *testing.T) {
	testHeaders := `{"Content-Type" : }`
	_, err := ParseHeadersToMap(testHeaders)

	if assert.NotNil(t, err) {
		assert.Equal(t, "could not parse headers into native map", err.Error())
	}

}
