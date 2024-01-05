package main

import (
	"fmt"
	"time"

	"github.com/rolacher/go-smartme/smartmeapi"
)

type ValCmd struct {
	DeviceId string    `arg help:"The device id in the form XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX."`
	Date     time.Time `arg optional format:"2006-01-02T15:04:05Z" help:"Get last values before the given date (form: YYYY-MM-DDTHH:MM:SSZ). Get all (last) values of the device if omitted."`
	EndDate  time.Time `arg optional format:"2006-01-02T15:04:05Z" help:"Get multiple values of a device, from date to end-date with interval between the values."`
	Interval int       `arg optional set:"defint=60" default:"${defint}" help:"Interval in minutes between to values, when date and end-date are given. Default is ${defint}."`
}

func (val *ValCmd) Run(globals *Globals) error {
	var err error
	if err = initializeApi(globals.Configfile); err != nil {
		return err
	}
	if val.Date.IsZero() {
		// Command called without a date as argument (only device id)
		var device *smartmeapi.ValuesData
		if device, err = smartmeapi.GetValues(val.DeviceId); err != nil {
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
		// Called with one date (after the device id) as argument
		if val.EndDate.IsZero() {
			var device *smartmeapi.ValuesData
			if device, err = smartmeapi.GetValuesInPast(val.DeviceId, val.Date); err != nil {
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
			// Called with two dates (after the device id) as arguments
			var device *[]smartmeapi.ValuesData
			if device, err = smartmeapi.GetValuesInPastMultiple(val.DeviceId, val.Date, val.EndDate, val.Interval); err != nil {
				return err
			}
			var json []byte
			json, err = unmarshall(device)
			if err != nil {
				fmt.Printf("Error unmarshall object: %v", err)
				return err
			}
			fmt.Printf("%s\n", json)
		}
	}

	return nil
}
