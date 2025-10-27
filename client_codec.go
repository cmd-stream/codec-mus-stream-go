package codec

import (
	"github.com/cmd-stream/core-go"
	"github.com/mus-format/mus-stream-go"
)

// NewClientCodec creates a new MUS-based streaming client codec for
// cmd-stream-go.
//
// Parameters:
//   - cmdSer:   A MUS serializer for Command types implementing core.Cmd[T].
//   - resultSer: A MUS serializer for Result types implementing core.Result.
func NewClientCodec[T any](cmdSer mus.Serializer[core.Cmd[T]],
	resultSer mus.Serializer[core.Result],
) ClientCodec[T] {
	return ClientCodec[T]{serT: cmdSer, serV: resultSer}
}

// ClientCodec defines a MUS-based streaming client codec for cmd-stream-go.
type ClientCodec[T any] = codec[core.Cmd[T], core.Result,
	mus.Serializer[core.Cmd[T]], mus.Serializer[core.Result]]
