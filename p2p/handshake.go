package p2p

type HandShakeFunc func(Peer) error

func NOPHandShake(Peer) error {
	return nil
}