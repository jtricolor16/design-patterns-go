package main

import (
	"fmt"
	"net/http"
)

type ChainElements struct {
	Name         string
	Header       map[string]string
	Response     any
	HttpInstance http.Client
}

func main() {
	fmt.Println("CoR")
	request := ChainElements{
		Name:         "squirtle",
		Header:       map[string]string{"Authorization": "any"},
		HttpInstance: *http.DefaultClient,
	}
	validateName := &ValidateName{}
	validateHeaders := &ValidateHeaders{}
	executeRequest := &ExecuteRequest{}
	validateName.SetNext(validateHeaders)
	validateHeaders.SetNext(executeRequest)
	var err error
	request, err = validateName.Execute(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(request.Response)
}
