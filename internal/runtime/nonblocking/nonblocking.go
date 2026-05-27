package nonblocking

import "golang.org/x/sys/unix"

func Set(fd int) error {
	return unix.SetNonblock(
		fd,
		true,
	)
}
