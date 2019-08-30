package combinations

import (
	"math/big"
)

func All(set []string) (subsets [][]string) {
	bigone := big.NewInt(1)
	length := uint(len(set))
	limit := big.NewInt(0).Lsh(bigone, length)
	bitmask := big.NewInt(0)

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := big.NewInt(1); subsetBits.Cmp(limit) < 0; subsetBits.Add(subsetBits, bigone) {
		var subset []string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			bitmask.Rsh(subsetBits, object)
			if bitmask.Uint64()&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}
