package lapjv

import "math"

// MaxValue is the maximum cost allowed in the matrix
const MaxValue = math.MaxInt

// NewResult instantiates an allocated Result
func NewResult(dim int) *Result {
	return &Result{
		InRow: make([]int, dim),
		InCol: make([]int, dim),
	}
}

// Result returns by the LAPJV
type Result struct {
	// Total cost
	Cost float64
	// Assignments in row
	InRow []int
	// Assignments in col
	InCol []int
}

type Number interface {
	int | float64 | float32
}

// Lapjv is a naive port of the Jonker Volgenant Algorithm from C++ to Go
func Lapjv[T Number](matrix [][]T) *Result {
	var unassignedfound bool
	var i, imin, numfree, prvnumfree, freerow int
	var j, j1, j2, endofpath, last, low, up int
	var i0 int
	var usubmin T
	var v2 T
	var min T
	var umin T
	var h T

	dim := len(matrix)
	result := NewResult(dim)
	u := make([]T, dim)
	v := make([]T, dim)
	free := make([]int, dim)
	collist := make([]int, dim)
	matches := make([]int, dim)
	pred := make([]int, dim)
	d := make([]T, dim)
	Max := T(MaxValue)

	// skipping L53-54
	for j := dim - 1; j >= 0; j-- {
		min = matrix[0][j]
		imin = 0
		for i := 1; i < dim; i++ {
			if matrix[i][j] < min {
				min = matrix[i][j]
				imin = i
			}
		}

		v[j] = min
		matches[imin]++
		if matches[imin] == 1 {
			result.InRow[imin] = j
			result.InCol[j] = imin
		} else {
			result.InCol[j] = -1
		}
	}

	for i := 0; i < dim; i++ {
		if matches[i] == 0 {
			free[numfree] = i
			numfree++
		} else if matches[i] == 1 {
			j1 = result.InRow[i]
			min = Max
			for j := 0; j < dim; j++ {
				if j != j1 && matrix[i][j]-v[j] < min {
					min = matrix[i][j] - v[j]
				}
			}
			v[j1] -= min
		}
	}

	for loopcmt := 0; loopcmt < 2; loopcmt++ {
		k := 0
		prvnumfree = numfree
		numfree = 0
		for k < prvnumfree {
			i = free[k]
			k++
			umin = matrix[i][0] - v[0]
			j1 = 0
			usubmin = Max

			for j := 1; j < dim; j++ {
				h = matrix[i][j] - v[j]

				if h < usubmin {
					if h >= umin {
						usubmin = h
						j2 = j
					} else {
						usubmin = umin
						umin = h
						j2 = j1
						j1 = j
					}
				}
			}

			i0 = result.InCol[j1]
			if umin < usubmin {
				v[j1] = v[j1] - (usubmin - umin)
			} else if i0 >= 0 {
				j1 = j2
				i0 = result.InCol[j2]
			}

			result.InRow[i] = j1
			result.InCol[j1] = i
			if i0 >= 0 {
				if umin < usubmin {
					k--
					free[k] = i0
				} else {
					free[numfree] = i0
					numfree++
				}
			}
		}
	}

	for f := 0; f < numfree; f++ {
		freerow = free[f]
		for j := 0; j < dim; j++ {
			d[j] = matrix[freerow][j] - v[j]
			pred[j] = freerow
			collist[j] = j
		}

		low = 0
		up = 0
		unassignedfound = false

		for !unassignedfound {
			if up == low {
				last = low - 1
				min = d[collist[up]]
				up++

				for k := up; k < dim; k++ {
					j = collist[k]
					h = d[j]
					if h <= min {
						if h < min {
							up = low
							min = h
						}
						collist[k] = collist[up]
						collist[up] = j
						up++
					}
				}

				for k := low; k < up; k++ {
					if result.InCol[collist[k]] < 0 {
						endofpath = collist[k]
						unassignedfound = true
						break
					}
				}
			}

			if !unassignedfound {
				j1 = collist[low]
				low++
				i = result.InCol[j1]
				h = matrix[i][j1] - v[j1] - min

				for k := up; k < dim; k++ {
					j = collist[k]
					v2 = matrix[i][j] - v[j] - h

					if v2 < d[j] {
						pred[j] = i

						if v2 == min {
							if result.InCol[j] < 0 {
								endofpath = j
								unassignedfound = true
								break
							} else {
								collist[k] = collist[up]
								collist[up] = j
								up++
							}
						}

						d[j] = v2
					}
				}
			}
		}

		for k := 0; k <= last; k++ {
			j1 = collist[k]
			v[j1] += d[j1] - min
		}

		i = freerow + 1
		for i != freerow {
			i = pred[endofpath]
			result.InCol[endofpath] = i
			j1 = endofpath
			endofpath = result.InRow[i]
			result.InRow[i] = j1
		}
	}

	lapcost := T(0)
	for i := 0; i < dim; i++ {
		j = result.InRow[i]
		u[i] = matrix[i][j] - v[j]
		lapcost += matrix[i][j]
	}

	result.Cost = float64(lapcost)
	return result
}
