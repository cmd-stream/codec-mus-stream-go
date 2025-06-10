package cdc

import (
	"fmt"
	"reflect"

	"github.com/cmd-stream/core-go"
	"github.com/cmd-stream/transport-go"
	muss "github.com/mus-format/mus-stream-go"
)

type ClientCodec[T any] = codec[core.Cmd[T], core.Result, muss.Serializer[core.Cmd[T]], muss.Serializer[core.Result]]

type ServerCodec[T any] = codec[core.Result, core.Cmd[T], muss.Serializer[core.Result], muss.Serializer[core.Cmd[T]]]

type TypedClientCodec[T any, C core.Cmd[T], R core.Result] = typedClientCodec[T, C, R, muss.Serializer[C], muss.Serializer[R]]

type TypedServerCodec[T any, C core.Cmd[T], R core.Result] = typedServerCodec[T, C, R, muss.Serializer[C], muss.Serializer[R]]

// NewClientCodec creates a new server codec.
func NewServerCodec[T any](cmdSer muss.Serializer[core.Cmd[T]],
	resultSer muss.Serializer[core.Result]) ServerCodec[T] {
	return ServerCodec[T]{serT: resultSer, serV: cmdSer}
}

// NewClientCodec creates a new client codec.
func NewClientCodec[T any](cmdSer muss.Serializer[core.Cmd[T]],
	resultSer muss.Serializer[core.Result]) ClientCodec[T] {
	return ClientCodec[T]{serT: cmdSer, serV: resultSer}
}

func NewTypedServerCodec[T any, C core.Cmd[T], R core.Result](
	cmdSer muss.Serializer[C], resultSer muss.Serializer[R]) TypedServerCodec[T, C, R] {
	return TypedServerCodec[T, C, R]{serC: cmdSer, serR: resultSer}
}

func NewTypedClientCodec[T any, C core.Cmd[T], R core.Result](
	cmdSer muss.Serializer[C], resultSer muss.Serializer[R]) TypedClientCodec[T, C, R] {
	return TypedClientCodec[T, C, R]{serC: cmdSer, serR: resultSer}
}

type codec[T, V any, K muss.Serializer[T], M muss.Serializer[V]] struct {
	serT K
	serV M
}

func (c codec[T, V, K, M]) Encode(t T, w transport.Writer) (n int, err error) {
	n, err = c.serT.Marshal(t, w)
	return
}

func (c codec[T, V, K, M]) Decode(r transport.Reader) (v V, n int, err error) {
	v, n, err = c.serV.Unmarshal(r)
	return
}

type typedClientCodec[T any, C core.Cmd[T], R core.Result, K muss.Serializer[C], M muss.Serializer[R]] struct {
	serC K
	serR M
}

func (c typedClientCodec[T, C, R, K, M]) Encode(cmd core.Cmd[T], w transport.Writer) (n int, err error) {
	cmd1, ok := cmd.(C)
	if !ok {
		err = fmt.Errorf("codec expects Command of type %v, but got %v", reflect.TypeFor[C](), reflect.TypeOf(cmd))
		return
	}
	n, err = c.serC.Marshal(cmd1, w)
	return
}

func (c typedClientCodec[T, C, R, K, M]) Decode(r transport.Reader) (result core.Result, n int, err error) {
	result, n, err = c.serR.Unmarshal(r)
	return
}

type typedServerCodec[T any, C core.Cmd[T], R core.Result, K muss.Serializer[C], M muss.Serializer[R]] struct {
	serC K
	serR M
}

func (c typedServerCodec[T, C, R, K, M]) Encode(result core.Result, w transport.Writer) (n int, err error) {
	result1, ok := result.(R)
	if !ok {
		err = fmt.Errorf("codec expects Result of type %v, but got %v", reflect.TypeFor[C](), reflect.TypeOf(result))
		return
	}
	n, err = c.serR.Marshal(result1, w)
	return
}

func (c typedServerCodec[T, C, R, K, M]) Decode(r transport.Reader) (cmd core.Cmd[T], n int, err error) {
	cmd, n, err = c.serC.Unmarshal(r)
	return
}
