# overflow
 Check for integer overflows in Golang (Generic)

 # Install
 `go get github.com/acheong08/overflow`

 # Usage

```go
import (
  "github.com/acheong08/overflow"
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
Function signatures: https://pkg.go.dev/github.com/acheong08/overflow

```go
package overflow

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func Add[T number](a, b T) (T, bool)
func Sub[T number](a, b T) (T, bool)
func Mul[T number](a, b T) (T, bool)
func Div[T number](a, b T) (T, bool)
func Quotient[T number](a, b T) (T, T, bool)
```



# Copyright / License

This is a fork of https://github.com/JohnCGriffin/overflow. Generics were added to Golang in Go 1.18 (2022) and is used in place of build time code generators.

```
Copyright (c) 2017 John C. Griffin,

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
