package traversal

import "barbero.oe/mutants/dna"

type Vertical struct {
	Items  [][]rune
	column int
}

func (it *Vertical) Next() (dna.Sequence, bool) {
	current := it.column
	if current == len(it.Items) {
		return nil, false
	}
	it.column++
	columnTraversal := ColumnTraversal{it.Items, current, 0}
	return &columnTraversal, true
}

type ColumnTraversal struct {
	items  [][]rune
	column int
	row    int
}

func (it *ColumnTraversal) Next() (rune, bool) {
	current := it.row
	if current == len(it.items) {
		return 0, false
	}
	it.row++
	r := it.items[current][it.column]
	return r, true
}
