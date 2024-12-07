package overflow_test

import (
	"math"
	"testing"

	"github.com/g-utils/overflow"
)

// sample all possibilities of 8 bit numbers
// by checking against 64 bit numbers

func TestAlgorithms(t *testing.T) {

	errors := 0

	for a64 := int64(math.MinInt8); a64 <= int64(math.MaxInt8); a64++ {

		for b64 := int64(math.MinInt8); b64 <= int64(math.MaxInt8) && errors < 10; b64++ {

			a8 := int8(a64)
			b8 := int8(b64)

			if int64(a8) != a64 || int64(b8) != b64 {
				t.Fatal("LOGIC FAILURE IN TEST")
			}

			// ADDITION
			{
				r64 := a64 + b64

				// now the verification
				result, ok := overflow.Add(a8, b8)
				if ok && int64(result) != r64 {
					t.Errorf("failed to fail on %v + %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && int64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// SUBTRACTION
			{
				r64 := a64 - b64

				// now the verification
				result, ok := overflow.Sub(a8, b8)
				if ok && int64(result) != r64 {
					t.Errorf("failed to fail on %v - %v = %v instead of %v\n",
						a8, b8, result, r64)
				}
				if !ok && int64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// MULTIPLICATION
			{
				r64 := a64 * b64

				// now the verification
				result, ok := overflow.Mul(a8, b8)
				if ok && int64(result) != r64 {
					t.Errorf("failed to fail on %v * %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && int64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// DIVISION
			if b8 != 0 {
				r64 := a64 / b64

				// now the verification
				result, _, ok := overflow.Quotient(a8, b8)
				if ok && int64(result) != r64 {
					t.Errorf("failed to fail on %v / %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && result != 0 && int64(result) == r64 {
					t.Fail()
					errors++
				}
			}
		}
	}

}

func TestQuotient(t *testing.T) {
	q, r, ok := overflow.Quotient(100, 3)
	if r != 1 || q != 33 || !ok {
		t.Errorf("expected 100/3 => 33, r=1")
	}
	if _, _, ok = overflow.Quotient(1, 0); ok {
		t.Error("unexpected lack of failure")
	}
}

func TestUnsignedAlgorithms(t *testing.T) {

	errors := 0

	for a64 := uint64(0); a64 <= uint64(math.MaxUint8); a64++ {

		for b64 := uint64(0); b64 <= uint64(math.MaxUint8) && errors < 10; b64++ {

			a8 := uint8(a64)
			b8 := uint8(b64)

			if uint64(a8) != a64 || uint64(b8) != b64 {
				t.Fatal("LOGIC FAILURE IN TEST")
			}

			// ADDITION
			{
				r64 := a64 + b64

				// now the verification
				result, ok := overflow.Add(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v + %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// SUBTRACTION
			{
				r64 := a64 - b64

				// now the verification
				result, ok := overflow.Sub(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v - %v = %v instead of %v\n",
						a8, b8, result, r64)
				}
				if !ok && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// MULTIPLICATION
			{
				r64 := a64 * b64

				// now the verification
				result, ok := overflow.Mul(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v * %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// DIVISION
			if b8 != 0 {
				r64 := a64 / b64

				// now the verification
				result, _, ok := overflow.Quotient(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v / %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && result != 0 && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}
		}
	}
}

func TestUnsignedQuotient(t *testing.T) {
	q, r, ok := overflow.Quotient(100, 3)
	if r != 1 || q != 33 || !ok {
		t.Errorf("expected 100/3 => 33, r=1")
	}
	if _, _, ok = overflow.Quotient(1, 0); ok {
		t.Error("unexpected lack of failure")
	}
}

// Second set of tests for unsigned overflows in case my implementation is incorrect
// Source: github.com/rwxe/overflow
func TestAlgorithmsUintImpl2(t *testing.T) {

	errors := 0

	for a64 := uint64(0); a64 <= uint64(math.MaxUint8); a64++ {

		for b64 := uint64(0); b64 <= uint64(math.MaxUint8) && errors < 10; b64++ {

			a8 := uint8(a64)
			b8 := uint8(b64)

			if uint64(a8) != a64 || uint64(b8) != b64 {
				t.Fatal("LOGIC FAILURE IN TEST")
			}

			// ADDITION
			{
				r64 := a64 + b64

				// now the verification
				result, ok := overflow.Add(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v + %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// SUBTRACTION
			{
				r64 := a64 - b64

				// now the verification
				result, ok := overflow.Sub(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v - %v = %v instead of %v\n",
						a8, b8, result, r64)
				}
				if !ok && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// MULTIPLICATION
			{
				r64 := a64 * b64

				// now the verification
				result, ok := overflow.Mul(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v * %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}

			// DIVISION
			if b8 != 0 {
				r64 := a64 / b64

				// now the verifiggcation
				result, _, ok := overflow.Quotient(a8, b8)
				if ok && uint64(result) != r64 {
					t.Errorf("failed to fail on %v / %v = %v instead of %v\n",
						a8, b8, result, r64)
					errors++
				}
				if !ok && result != 0 && uint64(result) == r64 {
					t.Fail()
					errors++
				}
			}
		}
	}

}

func TestConvert(t *testing.T) {
	// int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64

	t.Run("positive: uint64->unsigned", func(t *testing.T) {
		t.Parallel()
		for i := uint64(0); i <= math.MaxUint8; i++ {
			testConvert4[uint64, uint64, uint32, uint16, uint8](t, i, true)
		}
		testConvert3[uint64, uint64, uint32, uint16](t, math.MaxUint16-1, true)
		testConvert3[uint64, uint64, uint32, uint16](t, math.MaxUint16, true)
		testConvert2[uint64, uint64, uint32](t, math.MaxUint32-1, true)
		testConvert2[uint64, uint64, uint32](t, math.MaxUint32, true)
		testConvert1[uint64, uint64](t, math.MaxUint64-1, true)
		testConvert1[uint64, uint64](t, math.MaxUint64, true)
	})

	t.Run("negative: uint64->unsigned", func(t *testing.T) {
		for i := uint64(math.MaxUint32 + 1); i <= uint64(math.MaxUint32+math.MaxUint8+1); i++ {
			testConvert3[uint64, uint32, uint16, uint8](t, i, false)
		}
		testConvert2[uint64, uint16, uint8](t, math.MaxUint16+1, false)
		testConvert1[uint64, uint8](t, math.MaxUint8+1, false)
	})

	t.Run("positive: uint64->signed", func(t *testing.T) {
		for i := uint64(0); i <= math.MaxInt8; i++ {
			testConvert4[uint64, int64, int32, int16, int8](t, i, true)
		}
		testConvert3[uint64, int64, int32, int16](t, math.MaxInt16-1, true)
		testConvert3[uint64, int64, int32, int16](t, math.MaxInt16, true)
		testConvert2[uint64, int64, int32](t, math.MaxInt32-1, true)
		testConvert2[uint64, int64, int32](t, math.MaxInt32, true)
		testConvert1[uint64, int64](t, math.MaxInt64-1, true)
		testConvert1[uint64, int64](t, math.MaxInt64, true)
	})

	t.Run("negative: uint64->signed", func(t *testing.T) {
		for i := uint64(math.MaxInt64 + 1); i <= uint64(math.MaxInt64+math.MaxUint8+1); i++ {
			testConvert4[uint64, int64, int32, int16, int8](t, i, false)
		}
		testConvert4[uint64, int64, int32, int16, int8](t, math.MaxUint64, false)
		testConvert4[uint64, int64, int32, int16, int8](t, math.MaxUint64-1, false)
		testConvert3[uint64, int32, int16, int8](t, math.MaxInt32+1, false)
		testConvert2[uint64, int16, int8](t, math.MaxInt16+1, false)
		testConvert1[uint64, int8](t, math.MaxInt8+1, false)
	})

	t.Run("positive: int64->signed", func(t *testing.T) {
		for i := int64(math.MinInt8); i <= math.MaxInt8; i++ {
			testConvert4[int64, int64, int32, int16, int8](t, i, true)
		}
		testConvert3[int64, int64, int32, int16](t, math.MinInt16, true)
		testConvert3[int64, int64, int32, int16](t, math.MinInt16+1, true)
		testConvert3[int64, int64, int32, int16](t, math.MaxInt16-1, true)
		testConvert3[int64, int64, int32, int16](t, math.MaxInt16, true)
		testConvert2[int64, int64, int32](t, math.MinInt32, true)
		testConvert2[int64, int64, int32](t, math.MinInt32+1, true)
		testConvert2[int64, int64, int32](t, math.MaxInt32-1, true)
		testConvert2[int64, int64, int32](t, math.MaxInt32, true)
		testConvert1[int64, int64](t, math.MinInt64, true)
		testConvert1[int64, int64](t, math.MinInt64+1, true)
		testConvert1[int64, int64](t, math.MaxInt64-1, true)
		testConvert1[int64, int64](t, math.MaxInt64, true)
	})

	t.Run("negative: int64->signed", func(t *testing.T) {
		for i := int64(math.MaxInt32 + 1); i <= int64(math.MaxInt32+math.MaxInt8+1); i++ {
			testConvert3[int64, int32, int16, int8](t, i, false)
		}
		testConvert2[int64, int16, int8](t, math.MaxInt16+1, false)
		testConvert1[int64, int8](t, math.MaxInt8+1, false)
	})

	t.Run("positive: int64->unsigned", func(t *testing.T) {
		for i := int64(0); i <= math.MaxUint8; i++ {
			testConvert4[int64, uint64, uint32, uint16, uint8](t, i, true)
		}
		testConvert3[int64, uint64, uint32, uint16](t, math.MaxUint16-1, true)
		testConvert3[int64, uint64, uint32, uint16](t, math.MaxUint16, true)
		testConvert2[int64, uint64, uint32](t, math.MaxUint32-1, true)
		testConvert2[int64, uint64, uint32](t, math.MaxUint32, true)
		testConvert1[int64, uint64](t, math.MaxInt64-1, true)
		testConvert1[int64, uint64](t, math.MaxInt64, true)
	})

	t.Run("negative: int64->unsigned", func(t *testing.T) {
		testConvert4[int64, uint64, uint32, uint16, uint8](t, math.MinInt64, false)
		testConvert4[int64, uint64, uint32, uint16, uint8](t, math.MinInt64+1, false)
		testConvert4[int64, uint64, uint32, uint16, uint8](t, -1, false)
		testConvert3[int64, uint32, uint16, uint8](t, math.MaxUint32+1, false)
		testConvert2[int64, uint16, int8](t, math.MaxUint16+1, false)
		for i := int64(math.MaxUint8 + 1); i <= math.MaxUint8+math.MaxUint8+2; i++ {
			testConvert1[int64, uint8](t, i, false)
		}
	})

	t.Run("negative: int32|16|8->uint64", func(t *testing.T) {
		testConvert1[int32, uint64](t, -1, false)
		testConvert1[int32, uint64](t, math.MinInt32, false)
		testConvert1[int16, uint64](t, -1, false)
		testConvert1[int16, uint64](t, math.MinInt16, false)
		testConvert1[int8, uint64](t, -1, false)
		testConvert1[int8, uint64](t, math.MinInt8, false)
	})

	// already covered by reverse conversion checks:
	// - "positive [u]int32|16|8 -> [u]int64"

	// impossible cases:
	// "negative uint32|16|8 -> [u]int64"
	// "negative [u]int32|16|8 -> int64"
}

func testConvert4[T overflow.Int, U1 overflow.Int, U2 overflow.Int, U3 overflow.Int, U4 overflow.Int](t *testing.T, a T, expectOk bool) {
	t.Helper()
	testConvert2[T, U1, U2](t, a, expectOk)
	testConvert2[T, U3, U4](t, a, expectOk)
}

func testConvert3[T overflow.Int, U1 overflow.Int, U2 overflow.Int, U3 overflow.Int](t *testing.T, a T, expectOk bool) {
	t.Helper()
	testConvert1[T, U1](t, a, expectOk)
	testConvert1[T, U2](t, a, expectOk)
	testConvert1[T, U3](t, a, expectOk)
}

func testConvert2[T overflow.Int, U1 overflow.Int, U2 overflow.Int](t *testing.T, a T, expectOk bool) {
	t.Helper()
	testConvert1[T, U1](t, a, expectOk)
	testConvert1[T, U2](t, a, expectOk)
}

func testConvert1[T overflow.Int, U overflow.Int](t *testing.T, a T, expectOk bool) {
	t.Helper()
	if b, ok := overflow.Convert[T, U](a); expectOk && !ok {
		t.Error("expected ok, got overflow: ", b)
	} else if !expectOk && ok {
		t.Error("expected overflow, got ok: ", b)
	} else if a2, ok2 := overflow.Convert[U, T](b); ok && !ok2 {
		t.Error("expected ok, got ok, but reverse conversion overflowed: ", a2)
	}
}

func BenchmarkConvert(b *testing.B) {
	for i := int64(-1_000_000_000); i < 1_000_000_000; i++ {
		_, _ = overflow.Convert[int64, uint64](i)
	}
}
