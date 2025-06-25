package p2p

type HandShakeFunc func(Peer) error

func NOPHandShake(peer Peer) error {
	// No operation handshake function
	return nil
}