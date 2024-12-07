# overflow
 Check for integer overflows in Golang (Generic)

 # Install
 `go get github.com/g-utils/overflow`

 # Usage

```go
import (
  "github.com/g-utils/overflow"
  "math"
)

func main(){
  a, b := math.MaxInt64 - 5, 6
  sum, ok := overflow.Add(a, b)
  if !ok {
    panic("Overflow occured!")
  }
}
```
Function signatures: https://pkg.go.dev/github.com/g-utils/overflow

```go
package overflow

type Int interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Add[T Int](a, b T) (T, bool)
func Sub[T Int](a, b T) (T, bool)
func Mul[T Int](a, b T) (T, bool)
func Div[T Int](a, b T) (T, bool)
func Quotient[T Int](a, b T) (T, T, bool)
func Convert[T Int](a, b T) (T, T, bool)
```

# Copyright / License

This is a fork of https://github.com/JohnCGriffin/overflow. Generics were added to Golang in Go 1.18 (2022) and is used in place of build time code generators.

Tests for unsigned integers have also been added from github.com/rwxe/overflow in case my logic is incorrect.
