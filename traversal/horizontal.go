package traversal

import "barbero.oe/mutants/dna"

type Horizontal struct {
	Items [][]rune
	row   int
}

func (it *Horizontal) Next() (dna.Sequence, bool) {
	current := it.row
	if current == len(it.Items) {
		return nil, false
	}
	it.row++
	rowTraversal := RowTraversal{items: it.Items[current]}
	return &rowTraversal, true
}

type RowTraversal struct {
	items  []rune
	column int
}

func (it *RowTraversal) Next() (rune, bool) {
	current := it.column
	if current == len(it.items) {
		return 0, false
	}
	it.column++
	return it.items[current], true
}
