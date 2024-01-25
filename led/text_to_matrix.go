package led

func CharsToMatrix(input string, spaceBetweenChars bool, inverted bool) [][]bool {

	var matrix = make([][]bool, 6)
	for i := range matrix {
		matrix[i] = make([]bool, 9)
	}

	var widthPerChar = 3
	if spaceBetweenChars {
		widthPerChar = 4
	}

	for i, char := range input {
		if i > 8 {
			break
		}
		// draw each character with one column of space between
		var charMatrix = charToMatrix(char)
		for y := 0; y < 6; y++ {
			for x := 0; x < 3; x++ {
				// make sure the character is not drawn off the screen
				if x+i*widthPerChar > 8 {
					continue
				}
				matrix[y][x+i*widthPerChar] = charMatrix[y][x]
			}
		}
	}

	if inverted {
		for y := 0; y < 6; y++ {
			for x := 0; x < 9; x++ {
				matrix[y][x] = !matrix[y][x]
			}
		}
	}

	return matrix
}

func charToMatrix(char rune) [][]bool {
	switch char {
	case 'A':
		return [][]bool{
			{false, true, false},
			{true, false, true},
			{true, true, true},
			{true, false, true},
			{true, false, true},
			{true, false, true},
		}
	case 'B':
		return [][]bool{
			{true, true, false},
			{true, false, true},
			{true, false, true},
			{true, true, false},
			{true, false, true},
			{true, true, false},
		}
	case 'C':
		return [][]bool{
			{false, true, true},
			{true, false, false},
			{true, false, false},
			{true, false, false},
			{true, false, false},
			{false, true, true},
		}
	case 'D':
		return [][]bool{
			{true, true, false},
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{true, true, false},
		}
	case 'E':
		return [][]bool{
			{true, true, true},
			{true, false, false},
			{true, true, false},
			{true, false, false},
			{true, false, false},
			{true, true, true},
		}
	case 'F':
		return [][]bool{
			{true, true, true},
			{true, false, false},
			{true, true, false},
			{true, false, false},
			{true, false, false},
			{true, false, false},
		}
	case 'G':
		return [][]bool{
			{false, true, true},
			{true, false, false},
			{true, false, false},
			{true, false, true},
			{true, false, true},
			{false, true, true},
		}

	case 'H':
		return [][]bool{
			{true, false, true},
			{true, false, true},
			{true, true, true},
			{true, false, true},
			{true, false, true},
			{true, false, true},
		}
	case 'I':
		return [][]bool{
			{true, true, true},
			{false, true, false},
			{false, true, false},
			{false, true, false},
			{false, true, false},
			{true, true, true},
		}
	case 'J':
		return [][]bool{
			{false, false, true},
			{false, false, true},
			{false, false, true},
			{false, false, true},
			{true, false, true},
			{false, true, false},
		}
	case 'K':
		return [][]bool{
			{true, false, true},
			{true, false, true},
			{true, true, false},
			{true, false, true},
			{true, false, true},
			{true, false, true},
		}
	case 'L':
		return [][]bool{
			{true, false, false},
			{true, false, false},
			{true, false, false},
			{true, false, false},
			{true, false, false},
			{true, true, true},
		}
	case 'M':
		return [][]bool{
			{true, false, true},
			{true, true, true},
			{true, true, true},
			{true, false, true},
			{true, false, true},
			{true, false, true},
		}
	case 'N':
		return [][]bool{
			{true, false, true},
			{true, false, true},
			{true, true, true},
			{true, true, true},
			{true, false, true},
			{true, false, true},
		}
	case 'O':
		return [][]bool{
			{false, true, false},
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{false, true, false},
		}
	case 'P':
		return [][]bool{
			{true, true, false},
			{true, false, true},
			{true, false, true},
			{true, true, false},
			{true, false, false},
			{true, false, false},
		}
	case 'Q':
		return [][]bool{
			{false, true, false},
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{true, true, false},
			{false, false, true},
		}
	case 'R':
		return [][]bool{
			{true, true, false},
			{true, false, true},
			{true, false, true},
			{true, true, false},
			{true, false, true},
			{true, false, true},
		}
	case 'S':
		return [][]bool{
			{false, true, true},
			{true, false, false},
			{false, true, false},
			{false, false, true},
			{false, false, true},
			{true, true, false},
		}
	case 'T':
		return [][]bool{
			{true, true, true},
			{false, true, false},
			{false, true, false},
			{false, true, false},
			{false, true, false},
			{false, true, false},
		}
	case 'U':
		return [][]bool{
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{false, true, false},
		}
	case 'V':
		return [][]bool{
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{false, true, false},
			{false, true, false},
		}
	case 'W':
		return [][]bool{
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{true, true, true},
			{true, true, true},
			{true, false, true},
		}
	case 'X':
		return [][]bool{
			{true, false, true},
			{true, false, true},
			{false, true, false},
			{false, true, false},
			{true, false, true},
			{true, false, true},
		}
	case 'Y':
		return [][]bool{
			{true, false, true},
			{true, false, true},
			{false, true, false},
			{false, true, false},
			{false, true, false},
			{false, true, false},
		}
	case 'Z':
		return [][]bool{
			{true, true, true},
			{false, false, true},
			{false, true, false},
			{true, false, false},
			{true, false, false},
			{true, true, true},
		}
	case '0':
		return [][]bool{
			{false, true, false},
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{true, false, true},
			{false, true, false},
		}
	case '1':
		return [][]bool{
			{false, true, false},
			{true, true, false},
			{false, true, false},
			{false, true, false},
			{false, true, false},
			{true, true, true},
		}
	case '2':
		return [][]bool{
			{false, true, false},
			{true, false, true},
			{false, false, true},
			{false, true, false},
			{true, false, false},
			{true, true, true},
		}
	case '3':
		return [][]bool{
			{true, true, false},
			{false, false, true},
			{false, false, true},
			{false, true, false},
			{false, false, true},
			{true, true, false},
		}
	case '4':
		return [][]bool{
			{false, false, true},
			{false, true, true},
			{true, false, true},
			{true, true, true},
			{false, false, true},
			{false, false, true},
		}
	case '5':
		return [][]bool{
			{true, true, true},
			{true, false, false},
			{true, true, false},
			{false, false, true},
			{false, false, true},
			{true, true, false},
		}
	case '6':
		return [][]bool{
			{false, true, false},
			{true, false, false},
			{true, true, false},
			{true, false, true},
			{true, false, true},
			{false, true, false},
		}
	case '7':
		return [][]bool{
			{true, true, true},
			{false, false, true},
			{false, false, true},
			{false, true, false},
			{false, true, false},
			{false, true, false},
		}
	case '8':
		return [][]bool{
			{false, true, false},
			{true, false, true},
			{false, true, false},
			{true, false, true},
			{true, false, true},
			{false, true, false},
		}
	case '9':
		return [][]bool{
			{false, true, false},
			{true, false, true},
			{true, false, true},
			{false, true, true},
			{false, false, true},
			{false, true, false},
		}
	case ' ':
		return [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
			{false, false, false},
			{false, false, false},
			{false, false, false},
		}
	}
	return [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}
}
