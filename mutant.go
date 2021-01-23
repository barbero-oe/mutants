package mutant

import (
	"barbero.oe/mutants/concurrency"
	"barbero.oe/mutants/dna"
	"barbero.oe/mutants/traversal"
)

// IsMutant detects mutant dna if more than four equal consecutive letters are
// found
func IsMutant(dna []string) bool {
	runesDna := toRunes(dna)
	done := make(chan struct{})
	defer close(done)
	signals := findOn(done,
		&traversal.Horizontal{Items: runesDna},
		&traversal.Vertical{Items: runesDna},
		&traversal.ObliqueRight{Items: runesDna},
		&traversal.ObliqueLeft{Items: runesDna})
	return checkMutations(signals)
}

func toRunes(dna []string) [][]rune {
	runes := make([][]rune, len(dna))
	for i, row := range dna {
		runes[i] = []rune(row)
	}
	return runes
}

func findOn(done <-chan struct{}, strategies ...dna.Iterator) <-chan struct{} {
	outs := make([]<-chan struct{}, len(strategies))
	for i, strategy := range strategies {
		ch := findMutantSequences(done, strategy)
		outs[i] = ch
	}
	return concurrency.Merge(outs)
}

func checkMutations(signal <-chan struct{}) bool {
	count := 0
	for range signal {
		count++
		if count > 1 {
			return true
		}
	}
	return false
}

func findMutantSequences(done <-chan struct{}, dna dna.Iterator) <-chan struct{} {
	out := make(chan struct{})
	go find(done, out, dna)
	return out
}

func find(done <-chan struct{}, out chan<- struct{}, dna dna.Iterator) {
	defer close(out)
	for {
		sequence, hasSeq := dna.Next()
		if !hasSeq {
			break
		}
		var base rune
		var count int
		for {
			item, hasItem := sequence.Next()
			if !hasItem {
				break
			}
			select {
			case <-done:
				return
			default:
				if base != item {
					base = item
					count = 1
				} else {
					count++
					if count == 4 {
						out <- struct{}{}
					}
				}
			}
		}
	}
}
