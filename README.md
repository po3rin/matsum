# matsum

<img src="https://img.shields.io/badge/go-v1.12-blue.svg"/>

This is a package for adding elements after numerical calculation.

## Usage

```go
package main

import (
    "math"

    "github.com/po3rin/matsum"
    "gonum.org/v1/gonum/mat"
)

func main() {
    // creates matrix.
    A := mat.NewDense(2, 2, []float64{0, 0, 0, 0})

    // prepares applied function
    f := math.Exp

    // Do !!
    result := matsum.Sum(A, f)
    // result: 4

    // ...
}
```

## Todo

Add RawSymmetricer case ...
