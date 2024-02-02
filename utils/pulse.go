package utils

import (
	"akai-util/led"
	"fmt"
	"github.com/sqp/pulseaudio"
	"log"
)

var pacClient *pulseaudio.Client
var pa *PAClient
var trackedSinks = []string{}

func InitPulse() {
	pacClient, err := pulseaudio.New()
	if err != nil {
		fmt.Println("Error connecting to pulseaudio")
		log.Fatal(err)
	}
	pa = NewPAClient(pacClient)
}

func PulseVolumeNonTrackedApps(value byte) {
	err := pa.RefreshStreams()
	if err != nil {
		fmt.Println("Error refreshing streams", err)
		return
	}

	ival := float32(value)
	ival = ival / 127 * 100

	led.SetFaderPosition(5, int(ival))

	fmt.Println(" Setting non tracked apps to", ival)
	err2 := pa.ProcessVolumeAction("*", ival)
	if err2 != nil {
		fmt.Println("Error setting volume", err2)
		return
	}
}

func PulseAppVolume(displayFader int, applicationName string) func(byte) {
	trackedSinks = append(trackedSinks, applicationName)
	fmt.Println("Registering app " + applicationName)
	return func(value byte) {
		err := pa.RefreshStreams()
		if err != nil {
			fmt.Println("Error refreshing streams", err)
			return
		}

		ival := float32(value)
		ival = ival / 127 * 100

		led.SetFaderPosition(displayFader, int(ival))

		fmt.Println(" Setting "+applicationName+" to", ival)
		err2 := pa.ProcessVolumeAction(applicationName, ival)
		if err2 != nil {
			fmt.Println("Error setting volume", err2)
			return
		}
	}
}
