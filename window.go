package stlplus

import "math"

func calcTWindow(tDg, sDg, nS, nP int, omega float64) int {
	if tDg == 0 {
		tDg = 1
	}
	if sDg == 0 {
		sDg = 1
	}

	ns := float64(nS)
	np := float64(nP)

	coefsA := [][]float64{{0.000103350651767650, 3.81086166990428e-6}, {-0.000216653946625270, 0.000708495976681902}}
	coefsB := [][]float64{{1.42686036792937, 2.24089552678906}, {-3.1503819836694, -3.30435316073732}, {5.07481807116087, 5.08099438760489}}
	coefsC := [][]float64{{1.66534145060448, 2.33114333880815}, {-3.87719398039131, -1.8314816166323}, {6.46952900183769, 1.85431548427732}}
	// estimate critical frequency for seasonal
	betac0 := coefsA[0][0] + coefsA[1][0]*omega
	betac1 := coefsB[0][0] + coefsB[1][0]*omega + coefsB[2][0]*omega*omega
	betac2 := coefsC[0][0] + coefsC[1][0]*omega + coefsC[2][0]*omega*omega

	fC := (1.0 - (betac0 + betac1/ns + betac2/(ns*ns))) / np

	betat0 := coefsA[0][0] + coefsA[1][0]*omega
	betat1 := coefsB[0][0] + coefsB[1][0]*omega + coefsB[2][0]*omega*omega
	betat2 := coefsC[0][0] + coefsC[1][0]*omega + coefsC[2][0]*omega*omega

	betat00 := betat0 - fC

	return nextOdd((-betat1 - math.Sqrt(betat1*betat1-4.0*betat00*betat2)) / (2.0 * betat00))
}
