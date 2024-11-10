package main

type Log interface {
	Info(message string)
	Error(message string)
}
