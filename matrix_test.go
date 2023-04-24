package lapjv_test

import (
	"testing"

	"github.com/philistino/lapjv"
	"github.com/stretchr/testify/assert"
)

func TestSquareMatrixToSquare(t *testing.T) {
	squareMatrix := newMatrix(3, 3)
	squareMatrix[0] = []int{1, 2, 3}
	squareMatrix[1] = []int{1, 2, 3}
	squareMatrix[2] = []int{1, 2, 3}

	result := lapjv.ToSquare(squareMatrix)

	assert.Equal(t, squareMatrix, result)
}

func TestVerticalMatrixToSquare(t *testing.T) {
	verticalMatrix := newMatrix(2, 1)
	verticalMatrix[0] = []int{1}
	verticalMatrix[1] = []int{1}

	expectedMatrix := newMatrix(2, 2)
	expectedMatrix[0] = []int{1, 0}
	expectedMatrix[1] = []int{1, 0}

	result := lapjv.ToSquare(verticalMatrix)

	assert.Equal(t, expectedMatrix, result)
}

func TestHorizontalMatrixToSquare(t *testing.T) {
	horizontalMatrix := newMatrix(1, 2)
	horizontalMatrix[0] = []int{1, 2}

	expectedMatrix := newMatrix(2, 2)
	expectedMatrix[0] = []int{1, 2}
	expectedMatrix[1] = []int{0, 0}

	result := lapjv.ToSquare(horizontalMatrix)

	assert.Equal(t, expectedMatrix, result)
}

func newMatrix(x, y int) [][]int {
	rows := make([][]int, x)

	for i := range rows {
		rows[i] = make([]int, y)
	}

	return rows
}

func newMatrixFloat(x, y int) [][]float64 {
	rows := make([][]float64, x)

	for i := range rows {
		rows[i] = make([]float64, y)
	}

	return rows
}

func TestEmptyMatrix(t *testing.T) {
	matrix := make([][]int, 0)
	result := lapjv.ToSquare(matrix)

	assert.Equal(t, matrix, result)
}
