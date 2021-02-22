package huntingwords

var alphabetic = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "X", "W", "Y", "Z"}
var directiontsWeights = []string{"H", "H", "V", "H", "V", "D"}
var orientationsWeights = []string{"", "", "I"}
var colors = []string{"\033[1;36m", "\033[1;35m", "\033[1;34m", "\033[1;33m", "\033[1;32m", "\033[1;31m"}

func randomDirections(opt Options) (direction string) {
	switch {
	case opt.WordsInverse && opt.WordsDiagonal:
		direction = directiontsWeights[getInt(0, 5)]
		if direction != "D" {
			direction += orientationsWeights[getInt(0, 2)]
		}
		return
	case opt.WordsInverse && !opt.WordsDiagonal:
		direction = directiontsWeights[getInt(0, 4)] + orientationsWeights[getInt(0, 2)]
		return
	case !opt.WordsInverse && opt.WordsDiagonal:
		direction = directiontsWeights[getInt(0, 5)]
		return
	default:
		direction = directiontsWeights[getInt(0, 4)]
		return
	}
}

func isPositionRepeated(positions []position, board Board) bool {
	for _, p := range positions {
		if len(board.Grid[p.row][p.col].word) > 0 {
			return true
		}
	}
	return false
}

func createHorizontalPositions(word string, rowMax, columnMax int, inverseWord bool) []position {
	row := getInt(0, rowMax-1)
	col := getInt(0, (columnMax - len(word) - 1))

	wordIndex := 0
	if inverseWord {
		wordIndex = len(word) - 1
	}

	positions := make([]position, len(word))
	for i := 0; i < len(word); i++ {
		pos := position{row: row, col: col, letter: string(word[wordIndex])}
		positions[i] = pos
		col++
		if inverseWord {
			wordIndex--
		} else {
			wordIndex++
		}
	}
	return positions
}

func createVerticalPositions(word string, rowMax, columnMax int, inverseWord bool) []position {
	row := getInt(0, (rowMax - len(word) - 1))
	col := getInt(0, (columnMax - 1))

	wordIndex := 0
	if inverseWord {
		wordIndex = len(word) - 1
	}

	positions := make([]position, len(word))
	for i := 0; i < len(word); i++ {
		pos := position{row: row, col: col, letter: string(word[wordIndex])}
		positions[i] = pos
		row++
		if inverseWord {
			wordIndex--
		} else {
			wordIndex++
		}
	}
	return positions
}

func createDiagonalPositions(word string, rowMax, columnMax int) []position {
	row := getInt(0, (rowMax - len(word) - 1))
	col := getInt(0, (columnMax - len(word) - 1))

	wordIndex := 0

	positions := make([]position, len(word))
	for i := 0; i < len(word); i++ {
		pos := position{row: row, col: col, letter: string(word[wordIndex])}
		positions[i] = pos
		row++
		col++
		wordIndex++
	}
	return positions
}
