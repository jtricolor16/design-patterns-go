package main

type RequestMiddleware interface {
	Execute(request ChainElements) (ChainElements, error)
	SetNext(requestMiddleware RequestMiddleware)
	next(request ChainElements) (ChainElements, error)
}
