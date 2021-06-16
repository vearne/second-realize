package hashmap

import (
	"hash"
	"hash/crc64"
)

var (
	hasher hash.Hash64
)

func init() {
	hasher = crc64.New(crc64.MakeTable(crc64.ISO))
}

func HashCode(str string) uint64 {
	defer hasher.Reset()
	hasher.Write([]byte(str))
	return hasher.Sum64()
}
