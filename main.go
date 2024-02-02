package main

import (
	"akai-util/handles"
	"akai-util/led"
	"akai-util/obs"
	"akai-util/utils"
	"fmt"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/rakyll/portmidi"
	"strings"
)

func main() {

	portmidi.Initialize()
	totalDevices := portmidi.CountDevices()

	var readId portmidi.DeviceID
	var writeId portmidi.DeviceID

	// loop over devices
	for i := 0; i < totalDevices; i++ {
		deviceInfo := portmidi.Info(portmidi.DeviceID(i))
		// does the name contain "Akai"?
		if strings.Contains(deviceInfo.Name, "Akai") {
			// is it an input device?
			if deviceInfo.IsInputAvailable {
				readId = portmidi.DeviceID(i)
			}
			// is it an output device?
			if deviceInfo.IsOutputAvailable {
				writeId = portmidi.DeviceID(i)
			}
		}
	}

	// listen to the input device
	in, err := portmidi.NewInputStream(readId, 1024)
	if err != nil {
		panic(err)
	}

	out, outErr := portmidi.NewOutputStream(writeId, 1024, 0)
	if outErr != nil {
		panic(outErr)
	}

	go led.DrawVumeters(out)
	go led.StartBlinkLoop(out)
	obs.Init(out)

	// subscribe to the in.Listen() channel
	go func() {
		for e := range in.Listen() {
			handleMidiIn(e)
		}
	}()

	var midiPath = ""
	var p = utils.FilePathWalkDir("/dev/snd")
	for i := range p {
		if strings.Contains(p[i], "midi") {
			midiPath = "/dev/snd/" + p[i]
		}
	}
	fmt.Println("Using midi " + midiPath)

	utils.InitPulse()
	handles.SetupHandlers()
	systray.Run(onReady, nil)

}

func handleMidiIn(event portmidi.Event) {
	// convert Data1 and Data2 to channel, key, value
	switch status := event.Status & 0xF0; status {
	case 0x80:
		fmt.Printf("Note Off - Channel: %d, Key: %d, Velocity: %d", event.Status&0x0F, event.Data1, event.Data2)
	case 0x90:
		fmt.Printf("Note On - Channel: %d, Key: %d, Velocity: %d", event.Status&0x0F, event.Data1, event.Data2)
		utils.FireMidiNoteOn(byte(event.Data1), byte(event.Status&0x0F))
	// Add more cases for other MIDI event types as needed
	case 0xB0:
		fmt.Println(fmt.Sprintf("Control Change - Channel: %d, Controller Number: %d, Value: %d", event.Status&0x0F, event.Data1, event.Data2))
		utils.FireMidiControlChange(byte(event.Data1), byte(event.Status&0x0F), byte(event.Data2))
		return
	default:
		fmt.Printf("Unknown MIDI Event - Status: %d, Data1: %d, Data2: %d", event.Status, event.Data1, event.Data2)
	}
	fmt.Println()
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Akai Util")
	systray.SetTooltip("Akai Util")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon.Data)
}
