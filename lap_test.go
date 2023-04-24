package lapjv_test

import (
	"log"
	"testing"

	"github.com/philistino/lapjv"
	"github.com/stretchr/testify/assert"
)

func createDiagonalMatrix(m, n int) [][]int {
	matrix := make([][]int, n)

	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)

		for j := 0; j < n; j++ {
			matrix[i][j] = i * j
		}
	}

	return matrix
}

func createDiagonalMatrixFloat(m, n int) [][]float64 {
	matrix := make([][]float64, n)

	for i := 0; i < m; i++ {
		matrix[i] = make([]float64, n)

		for j := 0; j < n; j++ {
			matrix[i][j] = float64(i) * float64(j)
		}
	}

	return matrix
}

func TestLapjvSampleMatrix(t *testing.T) {
	matrix := newMatrix(4, 4)
	matrix[0] = []int{5, 2, 9, 2}
	matrix[1] = []int{6, 4, 6, 2}
	matrix[2] = []int{2, 4, 5, 1}
	matrix[3] = []int{2, 4, 5, 1}

	result := lapjv.Lapjv(matrix)
	assert.Equal(t, []int{1, 2, 3, 0}, result.InRow)
	assert.Equal(t, []int{3, 0, 1, 2}, result.InCol)
	assert.Equal(t, float64(11), result.Cost)
}

func TestLapjvSampleMatrixSame(t *testing.T) {
	matrix := newMatrix(4, 4)
	matrix[0] = []int{2, 4, 5, 2}
	matrix[1] = []int{2, 4, 5, 2}
	matrix[2] = []int{2, 4, 5, 2}
	matrix[3] = []int{2, 4, 5, 2}

	result := lapjv.Lapjv(matrix)
	assert.Equal(t, []int{3, 0, 1, 2}, result.InRow)
	assert.Equal(t, []int{1, 2, 3, 0}, result.InCol)
	assert.Equal(t, float64(13), result.Cost)
}

func TestResult(t *testing.T) {
	m := createDiagonalMatrix(10, 10)
	result := lapjv.Lapjv(m)

	assert.Equal(t, float64(120), result.Cost)
	assert.Equal(t, 7, result.InCol[2])
	assert.Equal(t, 4, result.InCol[5])
	assert.Equal(t, 9, result.InRow[0])
}

func TestLapjvDiagonalMatrix(t *testing.T) {
	matrix := createDiagonalMatrix(10, 10)
	for _, row := range matrix {
		log.Println(row)
	}
	result := lapjv.Lapjv(matrix)
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, result.InRow)
	assert.Equal(t, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, result.InCol)
	assert.Equal(t, float64(120), result.Cost)
}

func TestLapjvSampleMatrixFloat(t *testing.T) {
	matrix := newMatrixFloat(4, 4)
	matrix[0] = []float64{5, 2, 9, 2}
	matrix[1] = []float64{6, 4, 6, 2}
	matrix[2] = []float64{2, 4, 5, 1}
	matrix[3] = []float64{2, 4, 5, 1}

	result := lapjv.Lapjv(matrix)
	assert.Equal(t, []int{1, 2, 3, 0}, result.InRow)
	assert.Equal(t, []int{3, 0, 1, 2}, result.InCol)
	assert.Equal(t, float64(11), result.Cost)
}

func BenchmarkSolveInt100(b *testing.B) {
	matrix := createDiagonalMatrix(100, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lapjv.Lapjv(matrix)
	}
}

func BenchmarkSolveFloat100(b *testing.B) {
	matrix := createDiagonalMatrixFloat(100, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lapjv.Lapjv(matrix)
	}
}
