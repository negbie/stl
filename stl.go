package stl

type Stl struct {
	periodPoints   int
	dataPoints     int
	iter           int
	robustnessIter int
	sWindow        int
	lWindow        int
	tWindow        int
	sDegree        int
	lDegree        int
	tDegree        int
	sJump          int
	lJump          int
	tJump          int
	critFreq       float64
}

type Option func(*Stl)

func PeriodPoints(periodPoints int) Option {
	return func(args *Stl) { args.periodPoints = periodPoints }
}

func DataPoints(dataPoints int) Option {
	return func(args *Stl) { args.dataPoints = dataPoints }
}

func Iter(iter int) Option {
	return func(args *Stl) { args.iter = iter }
}

func RobustnessIter(robustnessIter int) Option {
	return func(args *Stl) { args.robustnessIter = robustnessIter }
}

func SWindow(sWindow int) Option {
	return func(args *Stl) { args.sWindow = sWindow }
}

func LWindow(lWindow int) Option {
	return func(args *Stl) { args.lWindow = lWindow }
}

func TWindow(tWindow int) Option {
	return func(args *Stl) { args.tWindow = tWindow }
}

func SDegree(sDegree int) Option {
	return func(args *Stl) { args.sDegree = sDegree }
}

func LDegree(lDegree int) Option {
	return func(args *Stl) { args.lDegree = lDegree }
}

func TDegree(tDegree int) Option {
	return func(args *Stl) { args.tDegree = tDegree }
}

func SJump(sJump int) Option {
	return func(args *Stl) { args.sJump = sJump }
}

func LJump(lJump int) Option {
	return func(args *Stl) { args.lJump = lJump }
}

func TJump(tJump int) Option {
	return func(args *Stl) { args.tJump = tJump }
}

func CritFreq(critFreq float64) Option {
	return func(args *Stl) { args.critFreq = critFreq }
}
