package cdc

import (
	"github.com/cmd-stream/base-go"
	"github.com/cmd-stream/transport-go"
	muss "github.com/mus-format/mus-stream-go"
)

type ClientCodec[T any] = Codec[base.Cmd[T], base.Result]

type ServerCodec[T any] = Codec[base.Result, base.Cmd[T]]

// NewClientCodec creates a new client codec.
func NewClientCodec[T any](cmdSer muss.Serializer[base.Cmd[T]],
	resultSer muss.Serializer[base.Result]) ClientCodec[T] {
	return Codec[base.Cmd[T], base.Result]{serT: cmdSer, serV: resultSer}
}

// NewClientCodec creates a new server codec.
func NewServerCodec[T any](resultSer muss.Serializer[base.Result],
	cmdSer muss.Serializer[base.Cmd[T]]) ServerCodec[T] {
	return Codec[base.Result, base.Cmd[T]]{serT: resultSer, serV: cmdSer}
}

// Codec represents a cmd-stream codec.
type Codec[T, V any] struct {
	serT muss.Serializer[T]
	serV muss.Serializer[V]
}

func (c Codec[T, V]) Encode(t T, w transport.Writer) (err error) {
	_, err = c.serT.Marshal(t, w)
	return
}

func (c Codec[T, V]) Decode(r transport.Reader) (v V, err error) {
	v, _, err = c.serV.Unmarshal(r)
	return
}
