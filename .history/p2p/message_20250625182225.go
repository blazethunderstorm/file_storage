package p2p

const(
	incomingMessage = 0x1
	incomingStream= 0x2
)

type RPC struct{
	From string
	Payload []byte
	Stream bool
}