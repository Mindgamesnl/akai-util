package led

import (
	"fmt"
	"github.com/rakyll/portmidi"
	"math/rand"
	"time"
)

var EMPTY = [][]bool{
	{false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false},
}

// matrix is 6 (height) x 9 (width)
func GetMidiCords(x, y int) (byte, byte) {
	var channel byte
	var key byte

	channel = 0
	key = 0

	if y == 1 {
		if x != 9 {
			key = 53
			channel = byte(x - 1)
		} else {
			key = 82
			channel = 0
		}
	}

	if y == 2 {
		if x != 9 {
			key = 54
			channel = byte(x - 1)
		} else {
			key = 83
			channel = 0
		}
	}

	if y == 3 {
		if x != 9 {
			key = 55
			channel = byte(x - 1)
		} else {
			key = 84
			channel = 0
		}
	}

	if y == 4 {
		if x != 9 {
			key = 56
			channel = byte(x - 1)
		} else {
			key = 85
			channel = 0
		}
	}

	if y == 5 {
		if x != 9 {
			key = 57
			channel = byte(x - 1)
		} else {
			key = 86
			channel = 0
		}
	}

	if y == 6 {
		if x != 9 {
			key = 52
			channel = byte(x - 1)
		} else {
			key = 81
			channel = 0
		}
	}

	if channel == 0 && key == 0 {
		fmt.Println("Invalid cords", x, y)
		return 0, 0
	}

	return channel, key
}

func SetLed(s *portmidi.Stream, channel int64, key int64, velocity int64, noteOn bool) error {
	var status int64
	if noteOn {
		status = 0x90 // Note On
	} else {
		status = 0x80 // Note Off
	}
	return s.WriteShort(status|channel, key, velocity)
}

func SetControlValue(s *portmidi.Stream, channel int64, control int64, value int64) error {
	return s.WriteShort(0xB0|channel, control, value)
}

// draw a 2d bool array to the midi matrix
func DrawMatrix(s *portmidi.Stream, matrix [][]bool) {
	for x := 0; x < 9; x++ {
		for y := 0; y < 6; y++ {
			channel, key := GetMidiCords(x+1, y+1)
			SetLed(s, int64(channel), int64(key), 127, matrix[y][x])
		}
	}
}

func DrawVumeters(s *portmidi.Stream) {

	for {
		updated, matrix := Render()
		if updated {
			DrawMatrix(s, matrix)
		}
		time.Sleep(time.Duration(MS_BETWEEN_UPDATES) * time.Millisecond)
	}
}

// draw a animated 3d cube to the midi matrix
func DrawRandomStars(s *portmidi.Stream) {
	// loop each second, make a random matrix and draw it
	for {
		matrix := GetRandomMatrix()
		DrawMatrix(s, matrix)
		time.Sleep(150 * time.Millisecond)
	}
}

// get a random 2d bool array
func GetRandomMatrix() [][]bool {
	var matrix [][]bool
	for x := 0; x < 9; x++ {
		var row []bool
		for y := 0; y < 6; y++ {
			row = append(row, rand.Intn(2) == 0)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
