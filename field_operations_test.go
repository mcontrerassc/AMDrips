package main

import (
	"testing"
        "fmt" 
	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
)

func BenchmarkFieldAddition(b *testing.B) {
	sizes := []int{1000, 10000, 100000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
			a := make([]fr.Element, size)
			bb := make([]fr.Element, size)

			for i := 0; i < size; i++ {
				a[i].SetRandom()
				bb[i].SetRandom()
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				for j := 0; j < size; j++ {
					var c fr.Element
					c.Add(&a[j], &bb[j])
				}
			}
		})
	}
}

func BenchmarkFieldMultiplication(b *testing.B) {
	sizes := []int{1000, 10000, 100000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
			a := make([]fr.Element, size)
			bb := make([]fr.Element, size)

			for i := 0; i < size; i++ {
				a[i].SetRandom()
				bb[i].SetRandom()
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				for j := 0; j < size; j++ {
					var c fr.Element
					c.Mul(&a[j], &bb[j])
				}
			}
		})
	}
}

func BenchmarkFieldInversion(b *testing.B) {
	sizes := []int{1000, 10000, 100000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
			a := make([]fr.Element, size)

			for i := 0; i < size; i++ {
				a[i].SetRandom()
			}

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				for j := 0; j < size; j++ {
					var c fr.Element
					c.Inverse(&a[j])
				}
			}
		})
	}
}
