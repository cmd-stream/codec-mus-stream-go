package codec

import (
	"github.com/cmd-stream/core-go"
	"github.com/mus-format/mus-stream-go"
)

// NewServerCodec creates a new MUS-based streaming server codec for
// cmd-stream-go.
//
// Parameters:
//   - cmdSer:   A MUS serializer for Command types implementing core.Cmd[T].
//   - resultSer: A MUS serializer for Result types implementing core.Result.
func NewServerCodec[T any](cmdSer mus.Serializer[core.Cmd[T]],
	resultSer mus.Serializer[core.Result],
) ServerCodec[T] {
	return ServerCodec[T]{serT: resultSer, serV: cmdSer}
}

// ServerCodec defines a MUS-based streaming server codec for cmd-stream-go.
type ServerCodec[T any] = codec[core.Result, core.Cmd[T],
	mus.Serializer[core.Result], mus.Serializer[core.Cmd[T]]]
