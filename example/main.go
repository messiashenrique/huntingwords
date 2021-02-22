package main

import h "github.com/messiashenrique/huntingwords"

func main() {
	o := h.Options{WordsInverse: true, WordsDiagonal: true}
	b1 := h.Board{Rows: 9, Cols: 9, Words: []string{"love", "clothes", "house", "golang", "peace"}, Options: o}
	b1.PrintRaw(true)
}
