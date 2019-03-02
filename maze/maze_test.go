package main

import (
	"bytes"
	"testing"
)

func BenchmarkParallel(b *testing.B) {
	parallel = true
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			main()
		}
	})
}

func BenchmarkNormal(b *testing.B) {
	main()
}
