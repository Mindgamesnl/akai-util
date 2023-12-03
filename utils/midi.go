package utils

import "fmt"

var MidiHandlers = map[byte]map[byte]func(byte){}

func RegisterMidiControlChange(controller byte, channel byte, handler func(value byte)) {
	if MidiHandlers[controller] == nil {
		MidiHandlers[controller] = map[byte]func(byte){}
	}
	MidiHandlers[controller][channel] = handler
}

func FireMidiControlChange(controller byte, channel byte, value byte) {
	if MidiHandlers[controller] == nil {
		fmt.Println("No handler for ", controller)
		return
	}
	if MidiHandlers[controller][channel] == nil {
		fmt.Println("No handler for ", controller, channel)
		return
	}
	MidiHandlers[controller][channel](value)
}
