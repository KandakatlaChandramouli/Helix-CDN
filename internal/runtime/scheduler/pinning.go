package scheduler

import "runtime"

func Tune() {
	runtime.GOMAXPROCS(
		runtime.NumCPU(),
	)
}
