package util

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"testing"
)

// sizes are the input lengths (in bytes) each algorithm is benchmarked against.
var sizes = []int{64, 1024, 1024 * 1024}

// makePayload builds a deterministic, reproducible input of n bytes.
func makePayload(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(i * 31)
	}
	return b
}

// benchHash runs the given hash function across every configured input size as
// sub-benchmarks (e.g. BenchmarkSHA256/1024B). SetBytes lets `go test` report
// throughput (MB/s) alongside ns/op.
func benchHash(b *testing.B, fn func([]byte)) {
	for _, size := range sizes {
		payload := makePayload(size)
		b.Run(fmt.Sprintf("%dB", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				fn(payload)
			}
		})
	}
}

func BenchmarkMD5(b *testing.B)    { benchHash(b, func(p []byte) { md5.Sum(p) }) }
func BenchmarkSHA1(b *testing.B)   { benchHash(b, func(p []byte) { sha1.Sum(p) }) }
func BenchmarkSHA256(b *testing.B) { benchHash(b, func(p []byte) { sha256.Sum256(p) }) }
func BenchmarkSHA512(b *testing.B) { benchHash(b, func(p []byte) { sha512.Sum512(p) }) }
