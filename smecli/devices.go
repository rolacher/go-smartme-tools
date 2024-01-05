package main

import (
	"fmt"

	"github.com/rolacher/go-smartme/smartmeapi"
)

type DevCmd struct {
	DeviceId struct {
		DeviceId string `arg optional help:"Show parameters of this device. Without id lists all devices with their parameters."`
	} `arg`
}

func (dev *DevCmd) Run(globals *Globals) error {
	var err error
	if err = initializeApi(globals.Configfile); err != nil {
		return err
	}
	if dev.DeviceId.DeviceId != "" {
		// print structure of one device
		var device *smartmeapi.Device
		if device, err = smartmeapi.GetDevice(dev.DeviceId.DeviceId); err != nil {
			return err
		}
		var json []byte
		json, err = unmarshall(device)
		if err != nil {
			fmt.Printf("Error unmarshall object: %v", err)
			return err
		}
		fmt.Printf("%s\n", json)
	} else {
		// print plain list of all devices
		// var devices*smartmeapi.Devices
		var devices *[]smartmeapi.Device
		if devices, err = smartmeapi.GetDevices(); err != nil {
			return err
		}
		var json []byte
		json, err = unmarshall(devices)
		if err != nil {
			fmt.Printf("Error unmarshall object: %v", err)
			return err
		}
		fmt.Printf("%s\n", json)
	}

	return nil
}
