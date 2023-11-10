package p2p

import "io"

type Decode interface {
	Decode(io.Reader, any) error
}
