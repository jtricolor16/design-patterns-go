package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TV struct {
	Volume int
	On     bool
}

const (
	maxVolume     = 20
	minVolume     = 0
	defaultVolume = 10
	folder        = "data"
)

func ChangeVolume(value bool) {
	var tv TV
	fileAddress := fmt.Sprintf("%s/tv.json", folder)
	content, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		fmt.Println("error to load tv file")
	}
	_ = json.Unmarshal(content, &tv)
	if !tv.On {
		fmt.Println("The TV is off")
		return
	}
	if value && tv.Volume != maxVolume {
		tv.Volume++
		fmt.Printf("The volume is %d\n", tv.Volume)
		jsonValue, _ := json.Marshal(tv)
		ioutil.WriteFile(fileAddress, jsonValue, 0644)
		return
	}
	if !value && tv.Volume != minVolume {
		tv.Volume--
		fmt.Printf("The volume is %d\n", tv.Volume)
		jsonValue, _ := json.Marshal(tv)
		ioutil.WriteFile(fileAddress, jsonValue, 0644)
		return
	}
	fmt.Printf("The volume was not changed %d\n", tv.Volume)
}

func TurnOnTv(value bool) {
	var tv TV
	fileAddress := fmt.Sprintf("%s/tv.json", folder)
	content, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		fmt.Println("error to load tv file")
		tv.Volume = defaultVolume
	} else {
		_ = json.Unmarshal(content, &tv)
	}

	tv.On = value
	fmt.Printf("The TV is %v\n", tv.On)
	jsonValue, _ := json.Marshal(tv)
	ioutil.WriteFile(fileAddress, jsonValue, 0644)
}
