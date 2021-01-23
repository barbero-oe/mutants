package dna

type Iterator interface {
	Next() (Sequence, bool)
}

type Sequence interface {
	Next() (rune, bool)
}
