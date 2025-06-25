package p2p

const(
	incomingMessage = 0x1
	incomingStream= iota
)

type RPC struct{
	from string
	Payload []byte
	Stream bool
}