package utils

import (
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

func PulseAppVolume(applicationName string) func(byte) {
	trackedSinks = append(trackedSinks, applicationName)
	fmt.Println("Registering app " + applicationName)
	return func(value byte) {
		err := pa.RefreshStreams()
		if err != nil {
			fmt.Println("Error refreshing streams", err)
			return
		}
		fmt.Println("Setting "+applicationName+" to", value)
		err2 := pa.ProcessVolumeAction(applicationName, float32(value))
		if err2 != nil {
			fmt.Println("Error setting volume", err2)
			return
		}
	}
}
