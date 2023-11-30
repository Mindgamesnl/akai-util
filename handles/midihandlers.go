package handles

import "akai-util/utils"

func SetupHandlers() {
	// spotify fader 1
	utils.RegisterMidiHandler(0, utils.PulseAppVolume("spotify"))

	// chrome fader 2
	utils.RegisterMidiHandler(1, utils.PulseAppVolume("Google Chrome"))

	// discord fader 3
	utils.RegisterMidiHandler(2, utils.PulseAppVolume("WEBRTC VoiceEngine"))

	// minecraft fader 4
	utils.RegisterMidiHandler(3, utils.PulseAppVolume("java"))
}
