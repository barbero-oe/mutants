package mutant

import (
	"sync"

	"barbero.oe/mutants/dna"
	"barbero.oe/mutants/traversal"
)

// atcg

// IsMutant detects mutant dna if more than four equal consecutive letters are
// found
func IsMutant(dna []string) bool {
	runesDna := toRunes(dna)
	done := make(chan struct{})
	defer close(done)
	horizontalSignals := findMutantSequences(done, &traversal.Horizontal{Items: runesDna})
	verticalSignals := findMutantSequences(done, &traversal.Vertical{Items: runesDna})
	return checkMutations(horizontalSignals, verticalSignals)
}

func toRunes(dna []string) [][]rune {
	runes := make([][]rune, len(dna))
	for i, row := range dna {
		runes[i] = []rune(row)
	}
	return runes
}

func checkMutations(channels ...<-chan struct{}) bool {
	count := 0
	for range merge(channels...) {
		count++
		if count > 1 {
			return true
		}
	}
	return false
}

func merge(channels ...<-chan struct{}) <-chan struct{} {
	var wg sync.WaitGroup
	wg.Add(len(channels))
	out := make(chan struct{})
	consume := func(ch <-chan struct{}) {
		defer wg.Done()
		for range ch {
			out <- struct{}{}
		}
	}
	for _, ch := range channels {
		go consume(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
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
