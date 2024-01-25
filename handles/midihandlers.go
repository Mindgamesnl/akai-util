package handles

import (
	"akai-util/commands"
	"akai-util/utils"
)

func SetupHandlers() {
	// spotify fader 1
	utils.RegisterMidiControlChange(7, 0, utils.PulseAppVolume(0, "spotify"))

	// chrome fader 2
	utils.RegisterMidiControlChange(7, 1, utils.PulseAppVolume(1, "Google Chrome"))

	utils.RegisterMidiControlChange(7, 2, utils.PulseAppVolume(2, "java"))

	// Firefox fader 3
	utils.RegisterMidiControlChange(7, 3, utils.PulseAppVolume(3, "Firefox"))

	// discord fader 3
	utils.RegisterMidiControlChange(7, 4, utils.PulseAppVolume(4, "WEBRTC VoiceEngine"))

	utils.RegisterMidiNoteOn(82, 0, commands.Script("/home/mats/Desktop/wallpapers", "random.sh"))
	utils.RegisterMidiNoteOn(83, 0, commands.Script("/home/mats/Desktop/wallpapers", "download-and-set.sh"))
}
