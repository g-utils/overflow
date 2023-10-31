package overflow

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Add[T number](a, b T) (T, bool) {
	c := a + b
	if (c > a) == (b > 0) {
		return c, true
	}
	return c, false
}

func Sub[T number](a, b T) (T, bool) {
	c := a - b
	if (c < a) == (b > 0) {
		return c, true
	}
	return c, false
}

func Mul[T number](a, b T) (T, bool) {
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
func Div[T number](a, b T) (T, bool) {
	q, _, ok := Quotient(a, b)
	return q, ok
}
func Quotient[T number](a, b T) (T, T, bool) {
	if b == 0 {
		return 0, 0, false
	}
	c := a / b
	status := (c < 0) == ((a < 0) != (b < 0))
	return c, a % b, status
}
