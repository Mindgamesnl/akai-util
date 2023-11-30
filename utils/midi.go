package utils

var MidiHandlers = map[byte]func(byte){}

func RegisterMidiHandler(key byte, handler func(value byte)) {
	MidiHandlers[key] = handler
}
