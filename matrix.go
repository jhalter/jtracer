package jtracer

type Matrix [][]float64

var IdentityMatrix = Matrix{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}

func (m Matrix) Multiply(m2 Matrix) Matrix {
	m3 := make(Matrix, len(m))

	for i := 0; i < len(m); i++ {
		m3[i] = make([]float64, len(m2[0]))
		for j := 0; j < len(m2[0]); j++ {
			for k := 0; k < len(m2); k++ {
				m3[i][j] += m[i][k] * m2[k][j]
			}
		}
	}

	return m3
}

func (m Matrix) MultiplyByTuple(t Tuple) *Tuple {
	var result []float64
	for row := 0; row < len(m); row++ {
		result = append(
			result,
			m[row][0]*t.X+
				m[row][1]*t.Y+
				m[row][2]*t.Z+
				m[row][3]*t.W,
		)
	}

	return &Tuple{X: result[0], Y: result[1], Z: result[2], W: result[3]}
}

func (m Matrix) Determinant() float64 {
	if len(m) == 2 && len(m[0]) == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}

	var result float64
	for col := 0; col < len(m); col++ {
		cofactor := m.Cofactor(0, col)
		result += m[0][col] * cofactor
	}
	return result
}

func (m Matrix) Submatrix(row, col int) Matrix {
	out := make(Matrix, len(m)-1)
	rr := 0
	// for each row in original matrix
	for i := 0; i < len(m); i++ {
		if i == row {
			continue
		}

		// create new output row
		out[rr] = make([]float64, len(m[0])-1)

		// for each column in the original matrix row
		rc := 0
		for j := 0; j < len(m[0]); j++ {
			if j != col {
				out[rr][rc] = m[i][j]
				rc += 1
			}
		}
		rr += 1
	}
	return out
}

func (m Matrix) Minor(row, col int) float64 {
	return m.Submatrix(row, col).Determinant()
}

func (m Matrix) Cofactor(row, col int) float64 {
	minor := m.Minor(row, col)
	if (row+col)%2 != 0 {
		return minor * -1.0
	}
	return minor
}

func (m Matrix) Inverse() Matrix {
	m2 := make(Matrix, len(m))
	for z := 0; z < len(m); z++ {
		m2[z] = make([]float64, len(m))
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			c := m.Cofactor(i, j)

			m2[j][i] = c / m.Determinant()
		}
	}

	return m2
}

func (m Matrix) Transpose() Matrix {
	out := make(Matrix, len(m))
	for i := 0; i < len(m); i++ {
		out[i] = make([]float64, len(m[0]))
		for j := 0; j < len(m[0]); j++ {
			out[i][j] = m[j][i]
		}
	}
	return out
}
