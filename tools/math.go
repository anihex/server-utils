package tools

import (
	"fmt"
	"math"
)

// GCD calculates the "Greatest Common Divider". It uses the Euclid's algorithm.
// This works recursivly.
func GCD(m, n int64) int64 {
	if n == 0 {
		return m
	}

	return GCD(n, m%n)
}

// MatMult calculates the matrix multiplicaton of two metrices.
func MatMult(A, B []float64, wA, wB int) ([]float64, error) {
	hA := len(A) / wA
	hB := len(B) / wB

	wC := wB
	hC := hA

	if wA != hB {
		return nil, fmt.Errorf(
			"cols of the first matrix (%d) must match the rows of the the second matrix (%d)",
			wA,
			hB,
		)
	}

	if trail := len(A) % wA; trail != 0 {
		return nil, fmt.Errorf("first matrix is not valid; trailing %d cols", trail)
	}

	if trail := len(B) % wB; trail != 0 {
		return nil, fmt.Errorf("second matrix is not valid; trailing %d cols", trail)
	}

	C := make([]float64, wC*hC)

	for i := 0; i < hC; i++ {
		for k := 0; k < wC; k++ {
			for j := 0; j < wA; j++ {
				C[i*wC+k] += A[i*wA+j] * B[j*wB+k]
			}
		}
	}

	return C, nil
}

// Ratio calculates ratio of w, h in response to the transformation
func Ratio(w, h int64, trans []float64) (int64, int64, int64, int64, error) {
	if w < 1 || h < 1 {
		return 0, 0, 0, 0, fmt.Errorf("invalid parameters provided for w (%d) or h (%d)", w, h)
	}

	bBR, err := MatMult(
		trans,
		[]float64{float64(w), float64(h), 1, 1},
		4,
		1,
	) // b Bottom Right
	if err != nil {
		return 0, 0, 0, 0, err
	}

	bBL, err := MatMult(
		trans,
		[]float64{0, float64(h), 1, 1},
		4,
		1,
	) // b Bottom Right
	if err != nil {
		return 0, 0, 0, 0, err
	}

	bTL, err := MatMult(
		trans,
		[]float64{0, 0, 1, 1},
		4,
		1,
	) // b Top Left
	if err != nil {
		return 0, 0, 0, 0, err
	}

	bTR, err := MatMult(
		trans,
		[]float64{float64(w), 0, 1, 1},
		4,
		1,
	) // b Top Right
	if err != nil {
		return 0, 0, 0, 0, err
	}

	maxX := math.Max(
		math.Max(bBR[0], bTR[0]),
		math.Max(bBL[0], bTL[0]),
	)
	maxY := math.Max(
		math.Max(bBL[1], bBR[1]),
		math.Max(bTL[1], bTR[1]),
	)

	minX := math.Min(
		math.Min(bTL[0], bBL[0]),
		math.Min(bBR[0], bTR[0]),
	)
	minY := math.Min(
		math.Min(bTL[1], bTR[1]),
		math.Min(bBL[1], bBR[1]),
	)

	wr := int64(math.Abs(math.Round(maxX - minX)))
	hr := int64(math.Abs(math.Round(maxY - minY)))

	gcd := GCD(wr, hr)

	if gcd == 0 {
		gcd = 1
	}

	return wr / gcd, hr / gcd, wr, hr, nil
}

func Determinant(A []float64) float64 {
	result := (A[0]*A[4]*A[8] + A[1]*A[5]*A[6] + A[2]*A[3]*A[7]) -
		(A[6]*A[4]*A[2] + A[7]*A[5]*A[0] + A[8]*A[3]*A[1])

	return result
}

func InverseAffine(A []float64) (result []float64, err error) {
	tA := make([]float64, 9)
	tb := make([]float64, 3)
	tC := make([]float64, 9)
	result = make([]float64, 16)

	tA[0], tA[1], tA[2] = A[0], A[1], A[2]
	tA[3], tA[4], tA[5] = A[4], A[5], A[6]
	tA[6], tA[7], tA[8] = A[8], A[9], A[10]

	const a, b, c, d, e, f, g, h, i = 0, 1, 2, 3, 4, 5, 6, 7, 8

	fa := 1 / Determinant(tA)

	tC[a] = (tA[e]*tA[i] - tA[f]*tA[h]) * fa
	tC[b] = (tA[c]*tA[h] - tA[b]*tA[i]) * fa
	tC[c] = (tA[b]*tA[f] - tA[c]*tA[e]) * fa
	tC[d] = (tA[f]*tA[g] - tA[d]*tA[i]) * fa
	tC[e] = (tA[a]*tA[i] - tA[c]*tA[g]) * fa
	tC[f] = (tA[c]*tA[d] - tA[a]*tA[f]) * fa
	tC[g] = (tA[d]*tA[h] - tA[e]*tA[g]) * fa
	tC[h] = (tA[b]*tA[g] - tA[a]*tA[h]) * fa
	tC[i] = (tA[a]*tA[e] - tA[b]*tA[d]) * fa

	tb[0], tb[1], tb[2] = A[3], A[7], A[11]

	tb, err = MatMult(tC, tb, 3, 1)
	if err != nil {
		return nil, err
	}

	result[0] = tC[0]
	result[1] = tC[1]
	result[2] = tC[2]
	result[3] = -tb[0]
	result[4] = tC[3]
	result[5] = tC[4]
	result[6] = tC[5]
	result[7] = -tb[1]
	result[8] = tC[6]
	result[9] = tC[7]
	result[10] = tC[8]
	result[11] = -tb[2]
	result[15] = 1

	return result, nil
}
