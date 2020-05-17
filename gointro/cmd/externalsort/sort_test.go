package main

import (
	"testing"
)

func BenchmarkCreatePipeline_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreatePipeline(infile, 100000000, 1)
	}
}

func BenchmarkCreatePipeline_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreatePipeline(infile, 100000000, 2)
	}
}

func BenchmarkCreatePipeline_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreatePipeline(infile, 100000000, 4)
	}
}

func BenchmarkCreatePipeline_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreatePipeline(infile, 100000000, 8)
	}
}

func BenchmarkCreatePipeline_16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreatePipeline(infile, 100000000, 6)
	}
}
