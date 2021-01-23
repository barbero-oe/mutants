package mutant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// atcg
func TestHumanDna(t *testing.T) {
	t.Run("human-1", isHuman([]string{"A", "T"}))
	t.Run("human-2", isHuman([]string{
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGACGG",
		"GCGTCA",
		"TCACTG"}))
}

func isHuman(dna []string) func(*testing.T) {
	return func(t *testing.T) {
		assert.Equal(t, false, IsMutant(dna))
	}
}
