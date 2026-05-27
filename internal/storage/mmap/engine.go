package mmap

import (
	"errors"
	"os"
	"sync"

	"golang.org/x/sys/unix"
)

const (
	PageSize   = 4096
	RegionSize = 64 * 1024 * 1024
)

type Engine struct {
	mu     sync.Mutex
	file   *os.File
	data   []byte
	offset uint64
	size   uint64
}

func Open(path string, size uint64) (*Engine, error) {
	if size == 0 {
		size = RegionSize
	}

	file, err := os.OpenFile(
		path,
		os.O_RDWR|os.O_CREATE,
		0644,
	)

	if err != nil {
		return nil, err
	}

	if err := file.Truncate(int64(size)); err != nil {
		return nil, err
	}

	data, err := unix.Mmap(
		int(file.Fd()),
		0,
		int(size),
		unix.PROT_READ|unix.PROT_WRITE,
		unix.MAP_SHARED,
	)

	if err != nil {
		return nil, err
	}

	return &Engine{
		file: file,
		data: data,
		size: size,
	}, nil
}

func (e *Engine) Append(buf []byte) (uint64, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if len(buf) == 0 {
		return 0, errors.New("empty append")
	}

	next := e.offset + uint64(len(buf))

	if next >= e.size {
		e.offset = 0
		next = uint64(len(buf))
	}

	start := e.offset

	copy(e.data[start:next], buf)

	e.offset = next

	return start, nil
}

func (e *Engine) Read(offset uint64, size uint64) ([]byte, error) {
	if offset+size >= e.size {
		return nil, errors.New("out of bounds")
	}

	return e.data[offset : offset+size], nil
}

func (e *Engine) Sync() error {
	return unix.Msync(
		e.data,
		unix.MS_SYNC,
	)
}

func (e *Engine) Close() error {
	if err := unix.Munmap(e.data); err != nil {
		return err
	}

	return e.file.Close()
}
