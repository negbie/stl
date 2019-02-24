package stl

func cInterp(m []int, fits, slopes []float64, at []int) []float64 {
	var i, j int
	var u, h, u2, u3 float64
	nAt := len(at)
	ans := make([]float64, nAt)

	j = 0 // index of leftmost vertex
	for i = 0; i < nAt; i++ {
		if at[i] > m[j+1] {
			j++
		}
		h = float64((m[j+1] - m[j]))
		u = float64((at[i] - m[j])) / h
		u2 = u * u
		u3 = u2 * u
		ans[i] = (2*u3-3*u2+1)*fits[j] +
			(3*u2-2*u3)*fits[j+1] +
			(u3-2*u2+u)*slopes[j]*h +
			(u3-u2)*slopes[j+1]*h
	}
	return ans
}
