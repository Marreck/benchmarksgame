package main

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func BenchmarkKnucleotide(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		f, err := os.Open("input25000000.txt")
		if err != nil {
			b.Fatal(err)
		}
		os.Stdin = f

		dna := input()
		WriteFrequencies(dna, 1)
		WriteFrequencies(dna, 2)
		WriteCount(dna, "GGT")
		WriteCount(dna, "GGTA")
		WriteCount(dna, "GGTATT")
		WriteCount(dna, "GGTATTTTAATT")
		WriteCount(dna, "GGTATTTTAATTTATAGT")
	}
}

func BenchmarkReadStdin(b *testing.B) {
	b.ReportAllocs()

	f, err := os.Open("input25000000.txt")
	if err != nil {
		b.Fatal(err)
	}

	os.Stdin = f

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		f.Seek(0, 0)
		start := time.Now()
		_ = readStdin()
		fmt.Println(time.Since(start))
	}

}

func TestKnucleotide(t *testing.T) {

	f, err := os.Open("knucleotide-input.txt")
	if err != nil {
		t.Fatal(err)
	}
	os.Stdin = f

	dna := input()

	var output string
	output += WriteFrequencies(dna, 1) + "\n"
	output += WriteFrequencies(dna, 2) + "\n"
	output += WriteCount(dna, "GGT") + "\n"
	output += WriteCount(dna, "GGTA") + "\n"
	output += WriteCount(dna, "GGTATT") + "\n"
	output += WriteCount(dna, "GGTATTTTAATT") + "\n"
	output += WriteCount(dna, "GGTATTTTAATTTATAGT") + "\n"


	expectedOutput := `T 31.520
A 29.600
C 19.480
G 19.400

AT 9.922
TT 9.602
TA 9.402
AA 8.402
GA 6.321
TC 6.301
TG 6.201
GT 6.041
CT 5.961
AG 5.841
CA 5.461
AC 5.441
CC 4.041
CG 4.021
GC 3.701
GG 3.341

54	GGT
24	GGTA
4	GGTATT
0	GGTATTTTAATT
0	GGTATTTTAATTTATAGT
`


	if expectedOutput != output {
		t.Fatal("Invalid output. Expected:\n", expectedOutput, "\nGot:\n", output)
	}
}

// Test performance of adding to a map[int]*int. With go1.11 this is faster than adding to map[int]int when there a few different keys
func BenchmarkMapPointer(b *testing.B) {
	b.ReportAllocs()

	m := make(map[int]*int)

	for n := 0; n < b.N; n++ {
		key := n % 4
		pointer, ok := m[key]
		if !ok {
			pointer = new(int)
			m[key] = pointer
		}
		*pointer++
	}
}

func BenchmarkMap(b *testing.B) {
	b.ReportAllocs()

	m := make(map[int]int)

	for n := 0; n < b.N; n++ {
		key := n % 4
		m[key]++
	}
}