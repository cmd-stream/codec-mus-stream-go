package cdc

import (
	"github.com/cmd-stream/core-go"
	"github.com/cmd-stream/transport-go"
	muss "github.com/mus-format/mus-stream-go"
)

type ClientCodec[T any] = Codec[core.Cmd[T], core.Result]

type ServerCodec[T any] = Codec[core.Result, core.Cmd[T]]

// NewClientCodec creates a new client codec.
func NewClientCodec[T any](cmdSer muss.Serializer[core.Cmd[T]],
	resultSer muss.Serializer[core.Result]) ClientCodec[T] {
	return Codec[core.Cmd[T], core.Result]{serT: cmdSer, serV: resultSer}
}

// NewClientCodec creates a new server codec.
func NewServerCodec[T any](resultSer muss.Serializer[core.Result],
	cmdSer muss.Serializer[core.Cmd[T]]) ServerCodec[T] {
	return Codec[core.Result, core.Cmd[T]]{serT: resultSer, serV: cmdSer}
}

// Codec represents a cmd-stream codec.
type Codec[T, V any] struct {
	serT muss.Serializer[T]
	serV muss.Serializer[V]
}

func (c Codec[T, V]) Encode(t T, w transport.Writer) (n int, err error) {
	n, err = c.serT.Marshal(t, w)
	return
}

func (c Codec[T, V]) Decode(r transport.Reader) (v V, n int, err error) {
	v, n, err = c.serV.Unmarshal(r)
	return
}
