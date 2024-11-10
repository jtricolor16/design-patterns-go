package main

type TemplateConnectNetworkInterface interface {
	on()
	applyConnection() error
	off()
}
