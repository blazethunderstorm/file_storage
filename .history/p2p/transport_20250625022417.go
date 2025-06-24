package p2p



type Transport interface{
	// Dial connects to a peer with the given ID.
	Dial(peerID string) (Connection, error)
}