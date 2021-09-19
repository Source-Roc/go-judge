//go:build !linux
// +build !linux

package model

import (
	"io"
	"os"
)

func fileToByte(f *os.File) ([]byte, error) {
	if _, err := f.Seek(0, 0); err != nil {
		return nil, err
	}
	var s int64
	if fi, err := f.Stat(); err != nil {
		return nil, err
	} else {
		s = fi.Size()
	}
	c := make([]byte, s)
	if _, err := io.ReadFull(f, c); err != nil {
		return nil, err
	}
	f.Close()
	return c, nil
}

func releaseByte(b []byte) {
}
