package led

import (
	"fmt"
)

var (
	totalFaders        = 0
	faderPositions     = map[int]int{}
	updated            = false
	drawingTextFrames  = 0
	drawingText        = false
	textToDraw         = ""
	MS_BETWEEN_UPDATES = 20
	includeLetterSpace = true
)

type Country string

var (
	rlName  = "Mats"
	igName  = "ToetMats"
	age     = 22
	origin  = Country("The Netherlands, Den Haag")
	type_   = "Student"
	hobbies = []string{"Skating", "Partying W/ friends", "Programming"}
	bands   = []string{"Bring Me The Horizon", "Vistas", "Neck Deep", "The Snuts"}
)

type Person struct {
	rlName  string
	igName  string
	age     int
	origin  Country
	type_   string
	hobbies []string
	bands   []string
}

func SetFaderPosition(fader int, position int) {
	faderPositions[fader] = position
	totalFaders = len(faderPositions)
	updated = true

	// draw for 1.5 seconds, based on MS_BETWEEN_UPDATES
	drawingTextFrames = 1500 / MS_BETWEEN_UPDATES
	drawingText = true
	textToDraw = fmt.Sprintf("%d", position)
	includeLetterSpace = position < 100
}

func Render() (bool, [][]bool) {
	if drawingTextFrames > 0 && !updated {
		drawingTextFrames--
		return false, nil
	} else if drawingTextFrames == 0 {
		drawingText = false
		// request redraw
		updated = true
	}

	if totalFaders == 0 || !updated {
		return false, nil
	} else {
		updated = false
	}

	if drawingText {
		return true, CharsToMatrix(textToDraw, includeLetterSpace, false)
	}

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

	return true, matrix
}
