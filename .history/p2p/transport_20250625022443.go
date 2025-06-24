package p2p



type Transport interface{
	// Dial connects to a peer with the given ID.
	Dial(peerID string) (Connection, error)
	// Listen starts listening for incoming connections.
	Listen() error
	// Accept waits for an incoming connection and returns it.
	Accept() (Connection, error)
	// Close closes the transport.
	Close() error
	// PeerID returns the ID of the local peer.
	PeerID() string
	// SetPeerID sets the ID of the local peer.
	SetPeerID(peerID string) error
	// IsConnected checks if a peer is connected.


	IsConnected(peerID string) bool
	// ConnectedPeers returns a list of connected peers.
	ConnectedPeers() []string
	// Send sends a message to a peer.
	Send(peerID string, message []byte) error
	// Receive receives a message from a peer.
	Receive(peerID string) ([]byte, error)
	// OnMessage registers a callback to be called when a message is received.			
	OnMessage(peerID string, callback func(message []byte)) error
	// OnConnect registers a callback to be called when a peer connects.
	OnConnect(peerID string, callback func()) error
		
}