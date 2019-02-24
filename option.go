package stl

type Option func(*Stl)

func OuterLoop(outer int) Option {
	return func(args *Stl) {
		args.outer = outer
		if outer < 0 {
			args.outer = outer * -1
		}
	}
}

func InnerLoop(inner int) Option {
	return func(args *Stl) {
		args.inner = inner
		if inner < 0 {
			args.inner = inner * -1
		}
	}
}

func SWindow(sWindow int) Option {
	return func(args *Stl) {
		args.sWindow = sWindow
		if sWindow < 0 {
			args.sWindow = sWindow * -1
		}
	}
}

func TWindow(tWindow int) Option {
	return func(args *Stl) {
		args.tWindow = tWindow
		if tWindow < 0 {
			args.tWindow = tWindow * -1
		}
	}
}

func LWindow(lWindow int) Option {
	return func(args *Stl) {
		args.lWindow = lWindow
		if lWindow < 0 {
			args.lWindow = lWindow * -1
		}
	}
}

func SDegree(sDegree int) Option {
	return func(args *Stl) {
		args.sDegree = sDegree
		if sDegree < 0 || sDegree > 2 {
			args.sDegree = 1
		}
	}
}

func TDegree(tDegree int) Option {
	return func(args *Stl) {
		args.tDegree = tDegree
		if tDegree < 0 || tDegree > 2 {
			args.tDegree = 1
		}
	}
}

func LDegree(lDegree int) Option {
	return func(args *Stl) {
		args.lDegree = lDegree
		if lDegree < 0 || lDegree > 2 {
			args.lDegree = 1
		}
	}
}

func SJump(sJump int) Option {
	return func(args *Stl) {
		args.sJump = sJump
		if sJump < 0 {
			args.sJump = sJump * -1
		}
	}
}

func TJump(tJump int) Option {
	return func(args *Stl) {
		args.tJump = tJump
		if tJump < 0 {
			args.tJump = tJump * -1
		}
	}
}

func LJump(lJump int) Option {
	return func(args *Stl) {
		args.lJump = lJump
		if lJump < 0 {
			args.lJump = lJump * -1
		}
	}
}

func CritFreq(critFreq float64) Option {
	return func(args *Stl) {
		args.critFreq = critFreq
		if critFreq < 0 {
			args.critFreq = critFreq * -1
		}
	}
}
