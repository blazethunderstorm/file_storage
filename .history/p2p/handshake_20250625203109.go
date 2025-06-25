package p2p

type HandShakeFunc func(Peer) error

func NOPHandShake(peer Peer) error {
	return nil
}