package huntingwords

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strings"
)

//Board is the struc used to build the puzzle: hunting-words
type Board struct {
	Rows    int
	Cols    int
	Words   []string
	Clues   []string
	Grid    [][]cell
	Options Options
}

//Options for configuration the builder of board
type Options struct {
	WordsInverse  bool
	WordsDiagonal bool
}

type cell struct {
	letter     string
	word       string
	isSelected bool
}

type position struct {
	row    int
	col    int
	letter string
}

//PrintRaw returns in console the Grid of letters
func (board *Board) PrintRaw(color bool) {
	board.createBoard()
	for r, row := range board.Grid {
		tempRow := ""
		for c, cell := range row {
			word := board.Grid[r][c].word
			if len(word) > 0 && color {
				tempRow += colors[indexOf(word, board.Words)%5] + cell.letter + "\033[0m "
			} else {
				tempRow += cell.letter + " "
			}
		}
		fmt.Println(tempRow)
	}
}

func (board *Board) createBoard() {

	err := board.checkConditions()
	if err != nil {
		log.Fatalln("Error:", err)
	}

	board.Grid = make([][]cell, board.Rows)
	for i, word := range board.Words {
		board.Words[i] = strings.ToUpper(word)
	}
	for r := range board.Grid {
		board.Grid[r] = make([]cell, board.Cols)
	}

	board.insertRandomLetters()
	board.createPositions()
}

func (board *Board) insertRandomLetters() {
	for r := range board.Grid {
		for c := range board.Grid[r] {
			board.Grid[r][c].letter = getString(alphabetic)
		}
	}
}

func (board *Board) createPositions() {
	for _, word := range board.Words {
		createPositionForWord(word, board)
	}
}

func (board *Board) checkConditions() error {
	rowOrCol := int(math.Min(float64(board.Rows), float64(board.Cols)))
	if len(board.Words) >= rowOrCol {
		return errors.New("The number of Words should be lesser than the length of columns or rows")
	}
	for _, word := range board.Words {
		if (len(word) + 2) > rowOrCol {
			errorSpecific := fmt.Sprintf("The word\033[1;31m %s\033[0m should be 2 characters lesser than the length of columns or rows", word)
			return errors.New(errorSpecific)
		}
	}
	return nil
}

func createPositionForWord(word string, board *Board) {
	row := board.Rows - 1
	col := board.Cols - 1
	positions := make([]position, len(word))

	direction := randomDirections(board.Options)
	switch direction {
	case "H":
		positions = createHorizontalPositions(word, row, col, false)
	case "HI":
		positions = createHorizontalPositions(word, row, col, true)
	case "V":
		positions = createVerticalPositions(word, row, col, false)
	case "VI":
		positions = createVerticalPositions(word, row, col, true)
	case "D":
		positions = createDiagonalPositions(word, row, col)
	}

	if isPositionRepeated(positions, *board) {
		createPositionForWord(word, board)
		return
	}

	for _, p := range positions {
		board.Grid[p.row][p.col].letter = p.letter
		board.Grid[p.row][p.col].word = word
	}
}
