package lapjv

// ToSquare squarify a matrix
func ToSquare[T Number](m [][]T) [][]T {
	rowsLen := len(m)
	if rowsLen == 0 {
		return m
	}

	colsLen := len((m)[0])

	diff := rowsLen - colsLen
	if diff == 0 {
		return m
	}

	size := rowsLen
	if colsLen > size {
		size = colsLen
	}

	matrix := make([][]T, size)

	for i := 0; i < size; i++ {
		matrix[i] = make([]T, size)
		for j := 0; j < size; j++ {
			if i < rowsLen && j < colsLen {
				matrix[i][j] = m[i][j]
			} else {
				matrix[i][j] = 0
			}

		}
	}

	return matrix
}
