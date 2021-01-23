package mutant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// atcg
func TestHumanDna(t *testing.T) {
	t.Run("human-1", isHuman([]string{"A", "T"}))
	t.Run("human-2", isHuman([]string{
		"ATCG",
		"TAGC",
		"TTTT",
		"ATCG"}))
	t.Run("human-3", isHuman([]string{
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGACGG",
		"GCGTCA",
		"TCACTG"}))
}

func TestMutantDna(t *testing.T) {
	t.Run("mutant-1", isMutant([]string{
		"AAAA",
		"TAGC",
		"TTTT",
		"ATCG"}))
}

func isHuman(dna []string) func(*testing.T) {
	return shouldBe(false, dna)
}

func isMutant(dna []string) func(*testing.T) {
	return shouldBe(true, dna)
}

func shouldBe(expected bool, dna []string) func(*testing.T) {
	return func(t *testing.T) {
		assert.Equal(t, expected, IsMutant(dna))
	}
}
