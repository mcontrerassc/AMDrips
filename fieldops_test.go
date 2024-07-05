package main

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
)

func BenchmarkFieldAddition(b *testing.B) {
	a := make([]fr.Element, 1000)
	bb := make([]fr.Element, 1000)

	for i := 0; i < 1000; i++ {
		a[i].SetRandom()
		bb[i].SetRandom()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			var c fr.Element
			c.Add(&a[j], &bb[j])
		}
	}
}

func BenchmarkFieldMultiplication(b *testing.B) {
	a := make([]fr.Element, 1000)
	bb := make([]fr.Element, 1000)

	for i := 0; i < 1000; i++ {
		a[i].SetRandom()
		bb[i].SetRandom()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			var c fr.Element
			c.Mul(&a[j], &bb[j])
		}
	}
}

func BenchmarkFieldInversion(b *testing.B) {
	a := make([]fr.Element, 1000)

	for i := 0; i < 1000; i++ {
		a[i].SetRandom()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			var c fr.Element
			c.Inverse(&a[j])
		}
	}
}
