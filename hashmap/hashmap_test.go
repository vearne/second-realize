package hashmap

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

func init() {
	rand.Seed(0)
}
func benchmark(b *testing.B, set, get, delete int) {
	h := NewHashMap(4, 1)
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for k := 0; k < set*100; k++ {
			wg.Add(1)
			go func() {
				key := strconv.Itoa(rand.Int())
				value := key
				h.Set(key, value)
				wg.Done()
			}()
		}
		for k := 0; k < get*100; k++ {
			wg.Add(1)
			go func() {
				key := strconv.Itoa(rand.Int())
				h.Get(key)
				wg.Done()
			}()
		}
		for k := 0; k < delete*100; k++ {
			wg.Add(1)
			go func() {
				key := strconv.Itoa(rand.Int())
				h.Delete(key)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSetMore(b *testing.B) { benchmark(b, 9, 1, 1) }

func BenchmarkGetMore(b *testing.B) { benchmark(b, 1, 9, 1) }

func BenchmarkEqual(b *testing.B) { benchmark(b, 5, 5, 5) }
