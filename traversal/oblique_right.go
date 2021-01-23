package traversal

import (
	"barbero.oe/mutants/dna"
)

type ObliqueRight struct {
	Items  [][]rune
	column int
	row    int
}

func (it *ObliqueRight) Next() (dna.Sequence, bool) {
	if it.row == 0 {
		current := it.column
		if current == len(it.Items) {
			it.row++
			return &DiagonalRight{it.Items, 0, 1}, true
		}
		it.column++
		return &DiagonalRight{it.Items, current, 0}, true
	}
	current := it.row
	if current == len(it.Items) {
		return nil, false
	}
	it.row++
	return &DiagonalRight{it.Items, 0, current}, true
}

type DiagonalRight struct {
	items  [][]rune
	column int
	row    int
}

func (it *DiagonalRight) Next() (rune, bool) {
	if it.row == len(it.items) || it.column == len(it.items) {
		return 0, false
	}
	currentRow, currentColumn := it.row, it.column
	it.row++
	it.column++
	return it.items[currentRow][currentColumn], true
}
