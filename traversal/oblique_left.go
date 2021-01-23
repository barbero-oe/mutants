package traversal

import (
	"barbero.oe/mutants/dna"
)

type ObliqueLeft struct {
	Items  [][]rune
	column int
	row    int
}

func (it *ObliqueLeft) Next() (dna.Sequence, bool) {
	if it.row == 0 {
		current := it.column
		if current == len(it.Items) {
			it.row++
			return &DiagonalLeft{it.Items, 0, 1}, true
		}
		it.column++
		return &DiagonalLeft{it.Items, current, 0}, true
	}
	current := it.row
	if current == len(it.Items) {
		return nil, false
	}
	it.row++
	return &DiagonalLeft{it.Items, len(it.Items) - 1, current}, true
}

type DiagonalLeft struct {
	items  [][]rune
	column int
	row    int
}

func (it *DiagonalLeft) Next() (rune, bool) {
	if it.column == -1 || it.row == len(it.items) {
		return 0, false
	}
	currentRow, currentColumn := it.row, it.column
	it.row++
	it.column--
	return it.items[currentRow][currentColumn], true
}
