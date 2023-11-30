package main

import (
	"akai-util/handles"
	"akai-util/utils"
	"fmt"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
	"github.com/longears/pixelslinger/midi"
	"strings"
)

func main() {
	var midiPath = ""
	var p = utils.FilePathWalkDir("/dev/snd")
	for i := range p {
		if strings.Contains(p[i], "midi") {
			midiPath = "/dev/snd/" + p[i]
		}
	}
	fmt.Println("Using midi " + midiPath)
	midiMessageChan := midi.GetMidiMessageStream(midiPath)
	midiState := midi.MidiState{}

	utils.InitPulse()
	handles.SetupHandlers()
	go systray.Run(onReady, nil)

	for {
		midiState.UpdateStateFromChannel(midiMessageChan)
		for i := range midiState.RecentMidiMessages {
			var m = midiState.RecentMidiMessages[i]
			fmt.Println(m)
			if m.Key == 0 {
				continue
			}

			var searchKey byte

			// specific for my controller, faders are all key 7 with index on channel
			if m.Key == 7 {
				searchKey = m.Channel
			} else {
				searchKey = m.Key
			}

			handler, found := utils.MidiHandlers[searchKey]
			if found {
				handler(m.Value)
			} else {
				fmt.Println("No handler for ", searchKey)
			}
		}

	}
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Akai Util")
	systray.SetTooltip("Akai Util")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon.Data)
}
