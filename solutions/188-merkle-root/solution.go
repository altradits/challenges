package piscine

import "crypto/sha256"

func doubleSHA188(b []byte) [32]byte {
	first := sha256.Sum256(b)
	return sha256.Sum256(first[:])
}

func MerkleRoot(txids [][32]byte) [32]byte {
	if len(txids) == 0 {
		return [32]byte{}
	}
	level := txids
	for len(level) > 1 {
		var next [][32]byte
		for i := 0; i < len(level); i += 2 {
			left := level[i]
			right := left
			if i+1 < len(level) {
				right = level[i+1]
			}
			var pair [64]byte
			copy(pair[:32], left[:])
			copy(pair[32:], right[:])
			next = append(next, doubleSHA188(pair[:]))
		}
		level = next
	}
	return level[0]
}
