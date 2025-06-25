package p2p

const(
	incomingMessage = 1x0
	incomingStream= iota
)

type RPC struct{
	from string
	Payload []byte
	Stream bool
}