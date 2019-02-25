package popcount_test

import (
	"github.com/forht/gopl/ch2/popcount"
	"testing"
)

const TN uint64 = 0x1234567890ABCDEF

func TestPopCounts(t *testing.T) {
	if popcount.PopCount(TN) != popcount.PopCountLoop(TN) {
		t.Error("PopCountLoop != PopCount")
	}
	if popcount.PopCount(TN) != popcount.PopCountShift(TN) {
		t.Error("PopCountShift != PopCount")
	}
	if popcount.PopCount(TN) != popcount.PopCountClear(TN) {
		t.Error("PopCountClear != PopCount")
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(TN)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(TN)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountShift(TN)
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountClear(TN)
	}
}
