package stl

func cMa(x []float64, nP int) []float64 {
	i, n := 0, len(x)
	nn := n - nP*2
	nnP := nP
	maTmp := 0.0

	ans := make([]float64, n-2*nnP)
	ma := make([]float64, nn+nnP+1)
	ma2 := make([]float64, nn+2)

	maTmp = 0
	for i = 0; i < nnP; i++ {
		maTmp = maTmp + x[i]
	}
	ma[0] = maTmp / float64(nnP)

	for i = nnP; i < nn+2*nnP; i++ {
		maTmp = maTmp - x[i-nnP] + x[i]
		ma[i-nnP+1] = maTmp / float64(nnP)
	}

	maTmp = 0
	for i = 0; i < nnP; i++ {
		maTmp = maTmp + ma[i]
	}
	ma2[0] = maTmp / float64(nnP)

	for i = nnP; i < nn+nnP+1; i++ {
		maTmp = maTmp - ma[i-nnP] + ma[i]
		ma2[i-nnP+1] = maTmp / float64(nnP)
	}

	maTmp = 0

	for i = 0; i < 3; i++ {
		maTmp = maTmp + ma2[i]
	}
	ans[0] = maTmp / 3

	for i = 3; i < nn+2; i++ {
		maTmp = maTmp - ma2[i-3] + ma2[i]
		ans[i-2] = maTmp / 3
	}
	return ans
}
