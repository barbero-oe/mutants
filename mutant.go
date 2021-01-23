package mutant

// atcg

// IsMutant detects mutant dna if more than four equal consecutive letters are
// found
func IsMutant(dna []string) bool {
	done := make(chan struct{})
	defer close(done)
	horizontalSignals := findHorizontalSequences(done, dna)
	return checkMutations(horizontalSignals)
}

func checkMutations(foundSequence <-chan struct{}) bool {
	count := 0
	for range foundSequence {
		count++
		if count > 1 {
			return true
		}
	}
	return false
}

func findHorizontalSequences(done <-chan struct{}, dna []string) <-chan struct{} {
	out := make(chan struct{})
	go func() {
		defer close(out)
		for _, row := range dna {
			cancel := findOnRow(done, out, row)
			if cancel {
				return
			}
		}
	}()
	return out
}

func findOnRow(done <-chan struct{}, out chan<- struct{}, row string) bool {
	var base rune
	var count int
	for _, cell := range row {
		select {
		case <-done:
			return true
		default:
			if base != cell {
				base = cell
				count = 1
			} else {
				count++
				if count == 4 {
					out <- struct{}{}
				}
			}
		}
	}
	return false
}
