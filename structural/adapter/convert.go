package main

import (
	"encoding/json"
	"strings"
)

type Convert interface {
	KeyValueToJson(string) (string, error)
}

type convertStringToJson struct{}

func NewConvertStringToJson() Convert {
	return &convertStringToJson{}
}

func (c *convertStringToJson) KeyValueToJson(request string) (string, error) {
	keyValues := strings.Split(request, ",")
	jsonStruct := map[string]string{}
	for _, keyValue := range keyValues {
		divideKeyValue := strings.Split(keyValue, ":")
		key := divideKeyValue[0]
		value := divideKeyValue[1]
		jsonStruct[key] = value
	}
	jsonRaw, err := json.Marshal(jsonStruct)
	if err != nil {
		return "", err
	}
	return string(jsonRaw), nil
}
