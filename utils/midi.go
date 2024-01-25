package utils

import "fmt"

var MidiHandlers = map[byte]map[byte]func(byte){}
var MidiNoteHandlers = map[byte]map[byte]func(int64, int64){}

func RegisterMidiNoteOn(note byte, channel byte, handler func(int64, int64)) {
	if MidiNoteHandlers[note] == nil {
		MidiNoteHandlers[note] = map[byte]func(int64, int64){}
	}
	MidiNoteHandlers[note][channel] = handler
}

func FireMidiNoteOn(note byte, channel byte) {
	if MidiNoteHandlers[note] == nil {
		fmt.Println("No handler for ", note)
		return
	}
	if MidiNoteHandlers[note][channel] == nil {
		fmt.Println("No handler for ", note, channel)
		return
	}
	MidiNoteHandlers[note][channel](int64(note), int64(channel))
}

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
