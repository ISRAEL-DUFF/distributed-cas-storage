package p2p

import "errors"

// ErrorInvalidHandshake is returned if the handshake between
// the local and remote node could not be esterblished
var ErrorInvalidHandshake = errors.New("invalid handshake")

type HandShakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error {
	return nil
}
