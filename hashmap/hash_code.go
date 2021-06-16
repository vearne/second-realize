package hashmap

import (
	"hash/maphash"
)

var hasher maphash.Hash

func HashCode(str string) uint64 {
	hasher.Reset()
	hasher.WriteString(str)
	return hasher.Sum64()
}
