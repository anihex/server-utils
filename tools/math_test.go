package tools

import (
	"math"
	"testing"
)

const radian = math.Pi / 180

func TestGCD(t *testing.T) {
	m := []int64{1920, 1280, 1200}
	n := []int64{1080, 720, 600}
	e := []int64{120, 80, 600}

	for i := range m {
		r := GCD(m[i], n[i])
		if r != e[i] {
			t.Errorf("Unexpected Value: %d instead of %d", r, e[i])
		}
	}
}

func TestMatMult1(t *testing.T) {
	m := []float64{
		0.5, 0, 0, 960,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	b := []float64{
		1920,
		1080,
		1,
		1,
	}

	e := []float64{
		1920,
		1080,
		1,
		1,
	}

	r, err := MatMult(m, b, 4, 1)
	if err != nil {
		t.Errorf("Unexpected Error: %v", err)
	}

	for i := range e {
		if e[i] != r[i] {
			t.Errorf("Error at %d: expected: %f; actual: %f", i, e[i], r[i])
		}
	}
}

func TestMatMult2(t *testing.T) {
	m := []float64{
		0.5, 0, 0, 960,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	b := []float64{
		0,
		0,
		1,
		1,
	}

	e := []float64{
		960,
		0,
		1,
		1,
	}

	r, err := MatMult(m, b, 4, 1)
	if err != nil {
		t.Errorf("Unexpected Error: %v", err)
	}

	for i := range e {
		if e[i] != r[i] {
			t.Errorf("Error at %d: expected: %f; actual: %f", i, e[i], r[i])
		}
	}
}

func TestRatio(t *testing.T) {
	w := []int64{
		1920,
		1280,
		1024,
		3840,
	}

	h := []int64{
		1080,
		720,
		768,
		2160,
	}

	ew := []int64{
		16,
		16,
		4,
		16,
	}

	eh := []int64{
		9,
		9,
		3,
		9,
	}

	tfr2 := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	tfr := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	tfr3, err := MatMult(tfr2, tfr, 4, 4)

	if err != nil {
		t.Error(err.Error())
		return
	}

	for i := range w {
		rw, rh, rew, reh, rerr := Ratio(w[i], h[i], tfr3)
		if rerr != nil {
			t.Errorf("unexpected error: %v", rerr)
		}

		if rw != ew[i] {
			t.Errorf(
				"w of %dth ratio incorrect: expected: %d actual: %d (w: %d)",
				i,
				ew[i],
				rw,
				rew,
			)
		}

		if rh != eh[i] {
			t.Errorf(
				"h of %dth ratio incorrect: expected: %d actual: %d (h: %d)",
				i,
				eh[i],
				rh,
				reh,
			)
		}
	}
}

func TestDeterminant(t *testing.T) {
	A := []float64{
		-1, 1, -8,
		1, -2, 0,
		3, 1, -3,
	}

	e := float64(-59)

	r := Determinant(A)

	if r != e {
		t.Errorf("expected: %f; actual: %f", e, r)
	}
}

func TestInverseAffine(t *testing.T) {
	A := []float64{
		-1, 1, -8, 0,
		1, -2, 0, 0,
		3, 1, -3, 0,
		0, 0, 0, 1,
	}

	Ai := []float64{
		-0.10169, 0.08475, 0.27119, 0,
		-0.05085, -0.45763, 0.13559, 0,
		-0.11864, -0.0678, -0.01695, 0,
		0, 0, 0, 1,
	}

	AiT, err := InverseAffine(A)
	if err != nil {
		t.Error(err.Error())
	}

	if len(AiT) != len(A) {
		t.Errorf(
			"result (%d) is not of the same length as the source (%d)",
			len(AiT),
			len(A),
		)
		return
	}

	for i := range AiT {
		v := math.Round(AiT[i]*100000) / 100000
		if v != Ai[i] {
			t.Errorf(
				"unexpected value at position %d: %f expecpted, got %f",
				i,
				Ai[i],
				v,
			)
		}
	}
}
