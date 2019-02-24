package stl

import "math"

func loessSTL(x, y []float64, span int, degree int, m []int, weights []float64, jump int) []float64 {
	n := len(y)
	if x == nil {
		x = make([]float64, n)
		for i := 0; i < n; i++ {
			x[i] = float64(i + 1)
		}
	}

	if weights == nil {
		weights = make([]float64, n)
		for i := 0; i < n; i++ {
			weights[i] = 1.0
		}
	}
	lenM := len(m)

	if span%2 == 0 {
		span++
	}

	s2 := (span + 1) / 2
	lIdx := make([]int, lenM)
	rIdx := make([]int, lenM)

	if n < span {
		for i := 0; i < lenM; i++ {
			lIdx[i] = 1
			rIdx[i] = n
		}
	} else {
		countSmallThanS2 := 0
		countLargeThanS2 := 0
		countLargeThanNMinusS2 := 0

		for i := 0; i < lenM; i++ {
			if m[i] < s2 {
				countSmallThanS2++
			} else if m[i] >= s2 && m[i] <= n-s2 {
				countLargeThanS2++
			} else {
				countLargeThanNMinusS2++
			}
		}
		i := 0
		for ; i < countSmallThanS2; i++ {
			lIdx[i] = 0
		}

		for ; i < countSmallThanS2+countLargeThanS2; i++ {
			lIdx[i] = m[i] - s2
		}

		for ; i < countSmallThanS2+countLargeThanS2+countLargeThanNMinusS2; i++ {
			lIdx[i] = n - span
		}

		for i := 0; i < lenM; i++ {
			rIdx[i] = lIdx[i] + span - 1
		}
	}

	maxDist := make([]float64, lenM)
	aa := 0.0
	bb := 0.0
	for i := 0; i < lenM; i++ {
		aa = math.Abs(float64(m[i]) - x[lIdx[i]])
		bb = math.Abs(x[rIdx[i]] - float64(m[i]))
		if aa > bb {
			maxDist[i] = aa
		} else {
			maxDist[i] = bb
		}
	}

	if span > n {
		for i := 0; i < len(maxDist); i++ {
			maxDist[i] = maxDist[i] + float64((span-n)/2)
		}
	}
	result, slopes := cLoess(x, y, degree, span, weights, m, lIdx, maxDist)

	// do interpolation
	res := make([]float64, len(result))
	copy(res, result)
	at := make([]int, n)
	for i := 0; i < n; i++ {
		at[i] = i + 1
	}

	if jump > 1 {
		res = cInterp(m, result, slopes, at)
	}
	return res
}

func cLoess(xx, yy []float64, degree, span int, ww []float64, m, lIdx []int, maxDist []float64) ([]float64, []float64) {

	n := len(xx)
	nM := len(m)

	x := make([]float64, span)
	w := make([]float64, span)
	xw := make([]float64, span)
	x2w := make([]float64, span)
	x3w := make([]float64, span)

	result := make([]float64, nM)
	slopes := make([]float64, nM)

	// variables for storing determinant intermediate values
	var (
		i, j                                       int
		r, tmp1, tmp2                              float64
		a, b, c, d, e, a1, b1, c1, a2, b2, c2, det float64
	)

	if span > n {
		span = n
	}

	// loop through all values of m
	for i = 0; i < nM; i++ {
		a = 0.0

		// get weights, x, and a
		for j = 0; j < span; j++ {
			w[j] = 0.0
			x[j] = xx[lIdx[i]+j] - float64(m[i])

			if x[j] > 0 {
				r = x[j]
			} else {
				r = -x[j]
			}

			// tricube
			tmp1 = r / maxDist[i]
			tmp2 = 1.0 - tmp1*tmp1*tmp1
			w[j] = tmp2 * tmp2 * tmp2

			// scale by user-defined weights
			w[j] = w[j] * ww[lIdx[i]+j]

			a = a + w[j]
		}

		if degree == 0 {
			// TODO: make sure denominator is not 0
			a1 = 1 / a
			for j = 0; j < span; j++ {
				// lIdx[j] = w[j] * a1;
				result[i] = result[i] + w[j]*a1*yy[lIdx[i]+j]
			}
		} else {
			// get xw, x2w, b, c for degree 1 or 2
			b = 0.0
			c = 0.0
			for j = 0; j < span; j++ {
				xw[j] = x[j] * w[j]
				x2w[j] = x[j] * xw[j]
				b = b + xw[j]
				c = c + x2w[j]
			}
			if degree == 1 {
				// TODO: make sure denominator is not 0
				det = 1 / (a*c - b*b)
				a1 = c * det
				b1 = -b * det
				c1 = a * det
				for j = 0; j < span; j++ {
					result[i] = result[i] + (w[j]*a1+xw[j]*b1)*yy[lIdx[i]+j]
					slopes[i] = slopes[i] + (w[j]*b1+xw[j]*c1)*yy[lIdx[i]+j]
				}
			} else {
				// TODO: make sure degree > 2 cannot be specified (and < 0 for that matter)
				// get x3w, d, and e for degree 2
				d = 0.0
				e = 0.0
				for j = 0; j < span; j++ {
					x3w[j] = x[j] * x2w[j]
					d = d + x3w[j]
					e = e + x3w[j]*x[j]
				}
				a1 = e*c - d*d
				b1 = c*d - e*b
				c1 = b*d - c*c
				a2 = c*d - e*b
				b2 = e*a - c*c
				c2 = b*c - d*a
				// TODO: make sure denominator is not 0
				det = 1 / (a*a1 + b*b1 + c*c1)
				a1 = a1 * det
				b1 = b1 * det
				c1 = c1 * det
				a2 = a2 * det
				b2 = b2 * det
				c2 = c2 * det
				for j = 0; j < span; j++ {
					result[i] = result[i] + (w[j]*a1+xw[j]*b1+x2w[j]*c1)*yy[lIdx[i]+j]
					slopes[i] = slopes[i] + (w[j]*a2+xw[j]*b2+x2w[j]*c2)*yy[lIdx[i]+j]
				}
			}
		}
	}
	return result, slopes
}
