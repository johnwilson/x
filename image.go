package x

import (
	"io"
)

type Image interface {
	io.ReadWriteCloser
}
