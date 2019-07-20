package matsum

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/mat"
)

const badTriangle = "matsum: invalid triangle"

// Sum returns the sum of the elements of the matrix.
func Sum(a mat.Matrix, f func(float64) float64) float64 {

	var sum float64
	aU, _ := untranspose(a)
	switch rma := aU.(type) {
	// // TODO: RawSymmetricer case
	// case mat.RawSymmetricer:
	// 	rm := rma.RawSymmetric()
	// 	for i := 0; i < rm.N; i++ {
	// 		// Diagonals count once while off-diagonals count twice.
	// 		sum += rm.Data[i*rm.Stride+i]
	// 		var s float64
	// 		for _, v := range rm.Data[i*rm.Stride+i+1 : i*rm.Stride+rm.N] {
	// 			s += v
	// 		}
	// 		sum += 2 * s
	// 	}
	// 	return sum
	case mat.RawTriangular:
		rm := rma.RawTriangular()
		var startIdx, endIdx int
		for i := 0; i < rm.N; i++ {
			// Start and end index for this triangle-row.
			switch rm.Uplo {
			case blas.Upper:
				startIdx = i
				endIdx = rm.N
			case blas.Lower:
				startIdx = 0
				endIdx = i + 1
			default:
				panic(badTriangle)
			}
			for _, v := range rm.Data[i*rm.Stride+startIdx : i*rm.Stride+endIdx] {
				sum += f(v)
			}
		}
		return sum
	case mat.RawMatrixer:
		rm := rma.RawMatrix()
		for i := 0; i < rm.Rows; i++ {
			for _, v := range rm.Data[i*rm.Stride : i*rm.Stride+rm.Cols] {
				sum += f(v)
			}
		}
		return sum
	case *mat.VecDense:
		rm := rma.RawVector()
		for i := 0; i < rm.N; i++ {
			sum += f(rm.Data[i*rm.Inc])
		}
		return sum
	default:
		r, c := a.Dims()
		for i := 0; i < r; i++ {
			for j := 0; j < c; j++ {
				sum += f(a.At(i, j))
			}
		}
		return sum
	}
}

// untranspose untransposes a matrix if applicable. If a is an Untransposer, then
// untranspose returns the underlying matrix and true. If it is not, then it returns
// the input matrix and false.
func untranspose(a mat.Matrix) (mat.Matrix, bool) {
	if ut, ok := a.(mat.Untransposer); ok {
		return ut.Untranspose(), true
	}
	return a, false
}
