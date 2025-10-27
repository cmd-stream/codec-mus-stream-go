// Package codec provides a MUS-based streaming codec implementation for
// cmd-stream-go.
package codec

import (
	"github.com/cmd-stream/transport-go"
	"github.com/mus-format/mus-stream-go"
)

type codec[T, V any, K mus.Serializer[T], M mus.Serializer[V]] struct {
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
