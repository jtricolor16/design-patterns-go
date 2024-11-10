package main

import "fmt"

func ConnectToNetwork(connectType TemplateConnectNetworkInterface) error {
	connectType.on()
	err := connectType.applyConnection()
	if err != nil {
		fmt.Println(err)
		return err
	}
	connectType.off()
	return nil
}

type defaultConnect struct {
	DeviceName string
}

func (d *defaultConnect) on() {
	fmt.Println("Press turn-on button ", d.DeviceName)
}

func (d *defaultConnect) applyConnection() error {
	return fmt.Errorf("you need to implement this function")
}

func (d *defaultConnect) off() {
	fmt.Println("Press turn-off button ", d.DeviceName)
}

type CellPhone struct {
	defaultConnect
}

func (c *CellPhone) applyConnection() error {
	fmt.Println("Connect in wi-fi ", c.DeviceName)
	return nil
}

func NewCellPhone(deviceName string) TemplateConnectNetworkInterface {
	return &CellPhone{
		defaultConnect{
			DeviceName: deviceName,
		},
	}
}

type PC struct {
	defaultConnect
}

func (p *PC) applyConnection() error {
	fmt.Println("Connect in wi-fi ", p.DeviceName)
	fmt.Println("Connect in rj45 cable ", p.DeviceName)
	return nil
}

func NewPC(deviceName string) TemplateConnectNetworkInterface {
	return &PC{
		defaultConnect{
			DeviceName: deviceName,
		},
	}
}
