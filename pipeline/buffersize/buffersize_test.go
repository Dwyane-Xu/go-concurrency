package main

import "testing"

// BenchmarkSimpleLine-8   	     154	   7841618 ns/op	     216 B/op	       2 allocs/op
func BenchmarkSimpleLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleLine()
	}
}

// BenchmarkFanline_0-8    	     132	   9021728 ns/op	     626 B/op	       6 allocs/op
// BenchmarkFanline_1-8    	     146	   8236888 ns/op	     631 B/op	       6 allocs/op
// BenchmarkFanline_2-8    	     156	   7346030 ns/op	     537 B/op	       6 allocs/op
// BenchmarkFanline_3-8    	     166	   6838190 ns/op	     543 B/op	       6 allocs/op
// BenchmarkFanline_4-8    	     181	   6629093 ns/op	     681 B/op	       6 allocs/op
// BenchmarkFanline_5-8    	     184	   6482633 ns/op	     554 B/op	       6 allocs/op
// BenchmarkFanline_6-8    	     186	   6518898 ns/op	     549 B/op	       6 allocs/op
// BenchmarkFanline_7-8    	     177	   6668170 ns/op	     580 B/op	       6 allocs/op
// BenchmarkFanline_8-8    	     181	   6399748 ns/op	     570 B/op	       6 allocs/op
// BenchmarkFanline_9-8    	     186	   6504315 ns/op	     576 B/op	       6 allocs/op
// BenchmarkFanline_10-8   	     153	   7141870 ns/op	     587 B/op	       6 allocs/op
func BenchmarkFanline_0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fanline(0)
	}
}

func BenchmarkFanline_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fanline(1)
	}
}

func BenchmarkFanline_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fanline(2)
	}
}

func BenchmarkFanline_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fanline(3)
	}
}

func BenchmarkFanline_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fanline(4)
	}
}

func BenchmarkFanline_5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fanline(5)
	}
}

func BenchmarkFanline_6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fanline(6)
	}
}

func BenchmarkFanline_7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fanline(7)
	}
}

func BenchmarkFanline_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fanline(8)
	}
}

func BenchmarkFanline_9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fanline(9)
	}
}

func BenchmarkFanline_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fanline(10)
	}
}
