package wal

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

type Writer struct {
	mu   sync.Mutex
	file *os.File
	buf  *bufio.Writer
}

func Open(path string) (*Writer, error) {
	file, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		0644,
	)

	if err != nil {
		return nil, err
	}

	return &Writer{
		file: file,
		buf:  bufio.NewWriterSize(file, 1<<20),
	}, nil
}

func (w *Writer) Append(payload []byte) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	var lenbuf [8]byte

	binary.LittleEndian.PutUint64(
		lenbuf[:],
		uint64(len(payload)),
	)

	if _, err := w.buf.Write(lenbuf[:]); err != nil {
		return err
	}

	if _, err := w.buf.Write(payload); err != nil {
		return err
	}

	return nil
}

func (w *Writer) Sync() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if err := w.buf.Flush(); err != nil {
		return err
	}

	return w.file.Sync()
}

func (w *Writer) Close() error {
	if err := w.Sync(); err != nil {
		return err
	}

	return w.file.Close()
}
