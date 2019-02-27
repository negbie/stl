package stl

import (
	"errors"
	"math"
)

type stl struct {
	outer    int
	inner    int
	sWindow  int
	lWindow  int
	tWindow  int
	sDegree  int
	lDegree  int
	tDegree  int
	sJump    int
	lJump    int
	tJump    int
	critFreq float64
}

// Decompose performs a STL decomposition.
func Decompose(series []float64, seasonality int, opts ...Option) (trend, seasonal, remainder []float64, err error) {
	tl := len(series)
	if tl < 11 {
		return nil, nil, nil, errors.New("series length must be at least 11")
	}
	if seasonality < 5 {
		return nil, nil, nil, errors.New("seasonality must be at least 5")
	}
	if tl <= 2*seasonality {
		return nil, nil, nil, errors.New("series length must be > 2 * seasonality")
	}

	// Default Options
	s := &stl{
		outer:    1,
		inner:    2,
		sWindow:  -1,
		tWindow:  -1,
		lWindow:  -1,
		sDegree:  1,
		tDegree:  1,
		lDegree:  1,
		sJump:    -1,
		tJump:    -1,
		lJump:    -1,
		critFreq: 0.05,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s.decompose(series, seasonality)
}

func (s *stl) decompose(series []float64, seasonality int) ([]float64, []float64, []float64, error) {
	nSeries := len(series)
	nPeriod := seasonality

	trend := make([]float64, nSeries)
	seasonal := make([]float64, nSeries)
	remainder := make([]float64, nSeries)
	for i := 0; i < nSeries; i++ {
		trend[i] = 0.0
		seasonal[i] = 0.0
		remainder[i] = 0.0
	}

	if s.lWindow == -1 {
		s.lWindow = nextOdd(float64(nPeriod))
	} else {
		s.lWindow = nextOdd(float64(s.lWindow))
	}

	s.sWindow = 10*nSeries + 1
	s.sDegree = 0
	s.sJump = int(math.Ceil(float64(s.sWindow) / 10.0))

	if s.tWindow == -1 {
		s.tWindow = nextOdd(1.5*float64(nPeriod)/(1.0-1.5/float64(s.sWindow)) + 0.5)
		//s.tWindow = calcTWindow(s.tDegree, s.sDegree, s.sWindow, nPeriod, s.critFreq)
	} else {
		s.tWindow = nextOdd(float64(s.tWindow))
	}

	if s.sJump == -1 {
		s.sJump = int(math.Ceil(float64(s.sWindow) / 10.0))
	}
	if s.tJump == -1 {
		s.tJump = int(math.Ceil(float64(s.tWindow) / 10.0))
	}
	if s.lJump == -1 {
		s.lJump = int(math.Ceil(float64(s.lWindow) / 10.0))
	}

	startIdx := nPeriod

	//cycleSubIndices will keep track of what part of the seasonal each observation belongs to
	cycleSubIndices := make([]int, nSeries)
	weight := make([]float64, nSeries)

	for i := 0; i < nSeries; i++ {
		cycleSubIndices[i] = i%nPeriod + 1
		weight[i] = 1.0
	}

	lenC := nSeries + 2*nPeriod
	C := make([]float64, lenC)
	D := make([]float64, nSeries)
	detrend := make([]float64, nSeries)

	tempSize := int(math.Ceil(float64(nSeries)/float64(nPeriod)) / 2)
	cycleSub := make([]float64, tempSize)
	subWeights := make([]float64, tempSize)
	cs1 := make([]int, nPeriod)
	cs2 := make([]int, nPeriod)

	for i := 0; i < nPeriod; i++ {
		cs1[i] = cycleSubIndices[i]
		cs2[i] = cycleSubIndices[nSeries-nPeriod+i]
	}

	ma3 := make([]float64, nSeries)
	L := make([]float64, nSeries)
	ljump := s.lJump
	tjump := s.tJump
	lenLev := int(math.Ceil(float64(nSeries) / float64(ljump)))
	lenTev := int(math.Ceil(float64(nSeries) / float64(tjump)))
	lEv := make([]int, lenLev)
	tEv := make([]int, lenTev)
	weightMeanAns := 0.0

	for oIter := 1; oIter <= s.outer; oIter++ {
		for iIter := 1; iIter <= s.inner; iIter++ {
			/** Step 1: detrending */
			for i := 0; i < nSeries; i++ {
				detrend[i] = series[i] - trend[i]
			}

			/** Step 2: smoothing of cycle-subseries */
			for i := 0; i < nPeriod; i++ {
				cycleSub = []float64{}
				subWeights = []float64{}
				for j := i; j < nSeries; j += nPeriod {
					if cycleSubIndices[j] == i+1 {
						cycleSub = append(cycleSub, detrend[j])
						subWeights = append(subWeights, weight[j])
					}
				}
				/*
				 C[c(cs1, cycleSubIndices, cs2) == i] <- rep(weighted.mean(cycleSub,
				 w = w[cycleSubIndices == i], na.rm = TRUE), cycleSub.length + 2)
				*/
				weightMeanAns = weightMean(cycleSub, subWeights)
				for j := i; j < nPeriod; j += nPeriod {
					if cs1[j] == i+1 {
						C[j] = weightMeanAns
					}
				}

				for j := i; j < nSeries; j += nPeriod {
					if cycleSubIndices[j] == i+1 {
						C[j+nPeriod] = weightMeanAns
					}
				}

				for j := 0; j < nPeriod; j++ {
					if cs2[j] == i+1 {
						C[j+nPeriod+nSeries] = weightMeanAns
					}
				}
			}

			/** Step 3: Low-pass filtering of collection of all the cycle-subseries
			# moving averages*/
			ma3 = cMa(C, nPeriod)

			for i, j := 0, 0; i < lenLev; i, j = i+1, j+ljump {
				lEv[i] = j + 1
			}

			if lEv[lenLev-1] != nSeries {
				tempLev := make([]int, lenLev+1)
				copy(tempLev, lEv)
				tempLev[lenLev] = nSeries
				L = loessSTL(nil, ma3, s.lWindow, s.lDegree, tempLev, weight, s.lJump)
			} else {
				L = loessSTL(nil, ma3, s.lWindow, s.lDegree, lEv, weight, s.lJump)
			}

			/** Step 4: Detrend smoothed cycle-subseries */
			/** Step 5: Deseasonalize */
			for i := 0; i < nSeries; i++ {
				seasonal[i] = C[startIdx+i] - L[i]
				D[i] = series[i] - seasonal[i]
			}

			/** Step 6: Trend Smoothing */
			for i, j := 0, 0; i < lenTev; i, j = i+1, j+tjump {
				tEv[i] = j + 1
			}

			if tEv[lenTev-1] != nSeries {
				tempTev := make([]int, lenTev+1)
				copy(tempTev, tEv)
				tempTev[lenTev] = nSeries
				trend = loessSTL(nil, D, s.tWindow, s.tDegree, tempTev, weight, s.tJump)
			} else {
				trend = loessSTL(nil, D, s.tWindow, s.tDegree, tEv, weight, s.tJump)
			}
		}
	}

	for i := 0; i < nSeries; i++ {
		remainder[i] = series[i] - trend[i] - seasonal[i]
	}

	return trend, seasonal, remainder, nil
}

func nextOdd(x float64) int {
	xx := int(math.Round(x))
	if xx%2 == 0 {
		return xx + 1
	}
	return xx
}

func weightMean(c []float64, w []float64) float64 {
	sum := 0.0
	sumW := 0.0
	len := len(c)
	for i := 0; i < len; i++ {
		if !math.IsNaN(c[i]) {
			sum += (c[i] * w[i])
			sumW += w[i]
		}
	}
	return sum / sumW
}
