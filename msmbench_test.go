package main_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	bls12377 "github.com/consensys/gnark-crypto/ecc/bls12-377"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
)

func BenchmarkMSM(b *testing.B) {
	sizes := []int{100, 1000, 10000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size-%d", size), func(b *testing.B) {
			// Generate random scalars and points
			scalars := make([]fr.Element, size)
			points := make([]bls12377.G1Affine, size)

			// Get the generator point
			generator := bls12377.G1Affine{}
			generator.Set(&bls12377.G1Affine{})

			for i := 0; i < size; i++ {
				// Generate a random scalar
				scalars[i].SetRandom()

				// Perform scalar multiplication with the generator point
				var tmp bls12377.G1Jac
				tmp.ScalarMultiplicationAffine(&generator, scalars[i].BigInt(new(big.Int)))

				// Convert to affine coordinates
				points[i].FromJacobian(&tmp)
			}

			// Reset the timer before the actual MSM operation
			b.ResetTimer()

			// Run the benchmark
			for i := 0; i < b.N; i++ {
				var result bls12377.G1Jac
				_, _ = result.MultiExp(points, scalars, ecc.MultiExpConfig{})
			}
		})
	}
}

