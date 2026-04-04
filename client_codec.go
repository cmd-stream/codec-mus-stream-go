package codec

import (
	"github.com/cmd-stream/cmd-stream-go/core"
	"github.com/mus-format/mus-stream-go"
)

// NewClientCodec creates a new MUS-based streaming client codec.
// It uses cmdSer for encoding Commands and resultSer for decoding Results.
//
// Parameters:
//   - cmdSer:   A MUS serializer for Command types.
//   - resultSer: A MUS serializer for Result types.
func NewClientCodec[T any](cmdSer mus.Serializer[core.Cmd[T]],
	resultSer mus.Serializer[core.Result],
) ClientCodec[T] {
	return ClientCodec[T]{serT: cmdSer, serV: resultSer}
}

// ClientCodec defines a MUS-based streaming client codec.
type ClientCodec[T any] = codec[core.Cmd[T], core.Result,
	mus.Serializer[core.Cmd[T]], mus.Serializer[core.Result]]
