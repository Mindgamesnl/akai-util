package led

var (
	totalFaders    = 0
	faderPositions = map[int]int{}
)

func SetFaderPosition(fader int, position int) {
	faderPositions[fader] = position
	totalFaders = len(faderPositions)
}

func DrawVuToMatrix() [][]bool {
	if totalFaders == 0 {
		return EMPTY
	}
	// draw VU meters for each fader
	// 9 * 6 matrix
	// fit the entire width by dividing the width by the number of faders
	// then draw a bar for each fader
	// the bar is the height of the fader position
	// the bar is the width of the width divided by the number of faders
	// the bar is drawn from the bottom up

	var widthPerFader = 1
	var matrix = make([][]bool, 6)
	for i := range matrix {
		matrix[i] = make([]bool, 9)
	}

	for fader, position := range faderPositions {
		// remap position 0-100 to 0-6
		var height = int(float32(position) / 100 * 6)
		// height = y 0-6
		// width widthPerFader offset by fader

		// fill Y from top to bottom
		for y := 0; y < height; y++ {
			for x := 0; x < widthPerFader; x++ {
				matrix[5-y][x+fader*widthPerFader] = true
			}
		}
	}

	return matrix
}
