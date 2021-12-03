package util

import (
	"crypto/md5"
	"crypto/sha256"
	"github.com/google/uuid"
	"testing"
)

func BenchmarkSha256(b *testing.B) {
	target := []byte(uuid.New().String())
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sha256.Sum256(target)
	}
}

func BenchmarkMd5(b *testing.B) {
	target := []byte(uuid.New().String())
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		md5.Sum(target)
	}
}
