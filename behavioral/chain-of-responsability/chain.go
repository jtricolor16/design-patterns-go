package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ValidateName struct {
	nextMiddleware RequestMiddleware
}

func (v *ValidateName) SetNext(requestMiddleware RequestMiddleware) {
	v.nextMiddleware = requestMiddleware
}

func (v *ValidateName) next(request ChainElements) (ChainElements, error) {
	return v.nextMiddleware.Execute(request)
}

func (v *ValidateName) Execute(request ChainElements) (ChainElements, error) {
	if request.Name == "" {
		return request, fmt.Errorf("invalid name")
	}
	fmt.Printf("Valid Name: %s\n", request.Name)
	return v.next(request)
}

type ValidateHeaders struct {
	nextMiddleware RequestMiddleware
}

func (v *ValidateHeaders) SetNext(requestMiddleware RequestMiddleware) {
	v.nextMiddleware = requestMiddleware
}

func (v *ValidateHeaders) next(request ChainElements) (ChainElements, error) {
	return v.nextMiddleware.Execute(request)
}

func (v *ValidateHeaders) Execute(request ChainElements) (ChainElements, error) {
	header := request.Header
	auth, ok := header["Authorization"]
	if !ok || auth == "" {
		return request, fmt.Errorf("invalid authorization header")
	}
	fmt.Printf("Valid authorization header: %s\n", auth)
	return v.next(request)
}

type ExecuteRequest struct {
	nextMiddleware RequestMiddleware
}

func (v *ExecuteRequest) SetNext(requestMiddleware RequestMiddleware) {
	v.nextMiddleware = requestMiddleware
}

func (v *ExecuteRequest) next(request ChainElements) (ChainElements, error) {
	return request, nil
}

func (v *ExecuteRequest) Execute(request ChainElements) (ChainElements, error) {
	type Pokemon struct {
		Name string `json:"name,omitempty"`
		Id   int    `json:"id,omitempty"`
	}
	var pokemon Pokemon
	path := fmt.Sprintf("/api/v2/pokemon/%s", request.Name)
	urlInstance := &url.URL{Scheme: "https", Host: "pokeapi.co", Path: path}
	fmt.Println(urlInstance.String())
	req := &http.Request{URL: urlInstance}
	response, err := request.HttpInstance.Do(req)
	if err != nil {
		return request, err
	}

	defer response.Body.Close()

	if response.StatusCode >= http.StatusOK && response.StatusCode <= http.StatusBadRequest {
		err = json.NewDecoder(response.Body).Decode(&pokemon)
		if err != nil {
			return request, err
		}
		request.Response = pokemon
		return v.next(request)
	}

	return request, fmt.Errorf("invalid request: %d", response.StatusCode)
}
