package codec

import (
	"github.com/cmd-stream/cmd-stream-go/core"
	"github.com/mus-format/mus-stream-go"
)

// NewServerCodec creates a new MUS-based streaming server codec.
// It uses resultSer for encoding results and cmdSer for decoding commands.
//
// Parameters:
//   - cmdSer:   A MUS serializer for Command types.
//   - resultSer: A MUS serializer for Result types.
func NewServerCodec[T any](cmdSer mus.Serializer[core.Cmd[T]],
	resultSer mus.Serializer[core.Result],
) ServerCodec[T] {
	return ServerCodec[T]{serT: resultSer, serV: cmdSer}
}

// ServerCodec defines a MUS-based streaming server codec.
type ServerCodec[T any] = codec[core.Result, core.Cmd[T],
	mus.Serializer[core.Result], mus.Serializer[core.Cmd[T]]]
