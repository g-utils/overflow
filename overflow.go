package overflow

type Int interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Add[T Int](a, b T) (T, bool) {
	c := a + b
	if (c > a) == (b > 0) {
		return c, true
	}
	return c, false
}

func Sub[T Int](a, b T) (T, bool) {
	c := a - b
	if (c < a) == (b > 0) {
		return c, true
	}
	return c, false
}

func Mul[T Int](a, b T) (T, bool) {
	if a == 0 || b == 0 {
		return 0, true
	}
	c := a * b
	if (c < 0) == ((a < 0) != (b < 0)) {
		if c/b == a {
			return c, true
		}
	}
	return c, false
}

func Div[T Int](a, b T) (T, bool) {
	q, _, ok := Quotient(a, b)
	return q, ok
}

func Quotient[T Int](a, b T) (T, T, bool) {
	if b == 0 {
		return 0, 0, false
	}
	c := a / b
	ok := (c < 0) == ((a < 0) != (b < 0))
	return c, a % b, ok
}

func Convert[T Int, U Int](a T) (U, bool) {
	b := U(a)
	if T(b) == a && (a < 0) == (b < 0) {
		return b, true
	}
	return b, false
}
