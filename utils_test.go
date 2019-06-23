package gofft

import (
	"math"
	"math/rand"
	"testing"
)

func TestIsPow2(t *testing.T) {
	// 1. Test all powers of 2 up to 2^30
	for i := 0; i < 31; i++ {
		x := 1 << uint32(i)
		r := IsPow2(x)
		if r != true {
			t.Errorf("IsPow2(%d), got: %t, expected: %t", x, r, true)
		}
	}

	// 2. Test all non-powers of 2 up to 2^15
	n := 1
	for x := 0; x < (1 << 15); x++ {
		if x == n {
			n <<= 1
			continue
		}
		r := IsPow2(x)
		if r != false {
			t.Errorf("IsPow2(%d), got: %t, expected: %t", x, r, false)
		}
	}
}

func TestNextPow2(t *testing.T) {
	for i := 0; i < 30; i++ {
		// 1. Test all powers of 2 up to 2^29
		x := 1 << uint32(i)
		r := NextPow2(x)
		if r != x {
			t.Errorf("NextPow2(%d), got: %d, expected: %d", x, r, x)
		}
		// 2. Test powers of 2 plus one
		r = NextPow2(x + 1)
		if r != 2*x {
			t.Errorf("NextPow2(%d+1), got: %d, expected: %d", x, r, 2*x)
		}
		// 3. Test random number between here and next power of 2
		if x > 1 {
			n := rand.Intn(x-1) + 1
			r = NextPow2(x + n)
			if r != 2*x {
				t.Errorf("NextPow2(%d+%d), got: %d, expected: %d", x, n, r, 2*x)
			}
		}
	}
}

func TestZeroPad(t *testing.T) {
	for i := 0; i < 1000; i++ {
		// Test random lengths between 0 and 10000, and random paddings between 0 and 1000
		N1 := rand.Intn(10000)
		N2 := N1 + rand.Intn(1000)
		x1 := complexRand(N1)
		x2 := ZeroPad(x1, N2)
		if len(x1) != N1 {
			t.Errorf("ZeroPad old array length, got: %d, expected: %d", len(x1), N1)
		}
		for j := 0; j < N1; j++ {
			if x1[j] != x2[j] {
				t.Errorf("ZeroPad copied section, got: x2[j] = %v, expected: x2[j] = %v", x2[j], x1[j])
			}
		}
		for j := N1; j < N2; j++ {
			if x2[j] != 0 {
				t.Errorf("ZeroPad padded section, got: x2[j] = %v, expected: x2[j] = %v", x2[j], 0)
			}
		}
	}
}

func TestFloat64ToComplex128Array(t *testing.T) {
	// Test random arrays of length 0 to 1000
	for i := 0; i < 1000; i++ {
		a := floatRand(i)
		b := Float64ToComplex128Array(a)
		if len(a) != len(b) {

		}
		for j := 0; j < i; j++ {
			if a[j] != real(b[j]) {
				t.Errorf("Float64ToComplex128Array, got: real(b[j]) = %v, expected: real(b[j]) = %v", real(b[j]), a[j])
			}
			if imag(b[j]) != 0 {
				t.Errorf("Float64ToComplex128Array, got: imag(b[j]) = %v, expected: imag(b[j]) = 0", imag(b[j]))
			}
		}
	}
}

func TestComplex128ToFloat64Array(t *testing.T) {
	// Test random arrays of length 0 to 1000
	for i := 0; i < 1000; i++ {
		a := complexRand(i)
		b := Complex128ToFloat64Array(a)
		for j := 0; j < i; j++ {
			if real(a[j]) != b[j] {
				t.Errorf("Complex128ToFloat64Array, got: b[j] = %v, expected: b[j] = %v", b[j], real(a[j]))
			}
		}
	}
}

func TestRoundFloat64Array(t *testing.T) {
	// Test random arrays of length 0 to 1000
	for i := 0; i < 1000; i++ {
		a := floatRand(i)
		b := make([]float64, i)
		copy(b, a)
		RoundFloat64Array(b)
		for j := 0; j < i; j++ {
			if math.Round(a[j]) != b[j] {
				t.Errorf("RoundFloat64Array, got: math.Round(a[j]) = %v, expected: math.Round(a[j]) = %v", math.Round(a[j]), b[j])
			}
		}
	}
}
