package p2p

const(
	incomingMessage = 0
	incomingStream= iota
)

type RPC struct{
	from string
	Payload []byte
	Stream bool
}