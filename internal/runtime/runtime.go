package runtime

import "runtime"

func Tune() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
