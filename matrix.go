package mgl

// column-major order
type Matrix3f struct {
	m [3][3]float64
}

func NewMatrix(
	m00 float64, m01 float64, m02 float64,
	m10 float64, m11 float64, m12 float64,
	m20 float64, m21 float64, m22 float64,
) Matrix3f {
	return Matrix3f{
		m: [3][3]float64{
			{m00, m10, m20},
			{m01, m11, m21},
			{m02, m12, m22},
		},
	}
}

func NewMatrixFromVectors(a Vector3f, b Vector3f, c Vector3f) Matrix3f {
	return Matrix3f{
		m: [3][3]float64{
			{a.X, a.Y, a.Z},
			{b.X, b.Y, b.Z},
			{c.X, c.Y, c.Z},
		},
	}
}

func (m Matrix3f) Get(i int, j int) float64 {
	return m.m[j][i]
}

func (m Matrix3f) GetVector(j int) Vector3f {
	return Vector3f{
		X: m.m[j][0],
		Y: m.m[j][1],
		Z: m.m[j][2],
	}
}

func (m Matrix3f) MatrixMul(b Matrix3f) Matrix3f {
	return NewMatrix(
		m.Get(0, 0)*b.Get(0, 0)+m.Get(0, 1)*b.Get(1, 0)+m.Get(0, 2)*b.Get(2, 0),
		m.Get(0, 0)*b.Get(0, 1)+m.Get(0, 1)*b.Get(1, 1)+m.Get(0, 2)*b.Get(2, 1),
		m.Get(0, 0)*b.Get(0, 2)+m.Get(0, 1)*b.Get(1, 2)+m.Get(0, 2)*b.Get(2, 2),

		m.Get(1, 0)*b.Get(0, 0)+m.Get(1, 1)*b.Get(1, 0)+m.Get(1, 2)*b.Get(2, 0),
		m.Get(1, 0)*b.Get(0, 1)+m.Get(1, 1)*b.Get(1, 1)+m.Get(1, 2)*b.Get(2, 1),
		m.Get(1, 0)*b.Get(0, 2)+m.Get(1, 1)*b.Get(1, 2)+m.Get(1, 2)*b.Get(2, 2),

		m.Get(2, 0)*b.Get(0, 0)+m.Get(2, 1)*b.Get(1, 0)+m.Get(2, 2)*b.Get(2, 0),
		m.Get(2, 0)*b.Get(0, 1)+m.Get(2, 1)*b.Get(1, 1)+m.Get(2, 2)*b.Get(2, 1),
		m.Get(2, 0)*b.Get(0, 2)+m.Get(2, 1)*b.Get(1, 2)+m.Get(2, 2)*b.Get(2, 2),
	)
}

func (m Matrix3f) VectorMul(v Vector3f) Vector3f {
	return Vector3f{
		X: m.Get(0, 0)*v.X + m.Get(0, 1)*v.Y + m.Get(0, 2)*v.Z,
		Y: m.Get(1, 0)*v.X + m.Get(1, 1)*v.Y + m.Get(1, 2)*v.Z,
		Z: m.Get(2, 0)*v.X + m.Get(2, 1)*v.Y + m.Get(2, 2)*v.Z,
	}
}
