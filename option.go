package stl

type Option func(*Stl)

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
