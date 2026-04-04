// Package codec provides a MUS-based streaming codec implementation for
// cmd-stream.
package codec

import (
	"fmt"

	"github.com/cmd-stream/cmd-stream-go/transport"
	"github.com/mus-format/mus-stream-go"
)

const ErrorPrefix = "codecmuss: "

// codec is a generic MUS-based streaming codec implementation.
// T is the type for encoding, V is the type for decoding.
// K and M are the corresponding MUS serializers.
type codec[T, V any, K mus.Serializer[T], M mus.Serializer[V]] struct {
	serT K
	serV M
}

// Encode writes a value of type T to the given transport.Writer.
// Returns the number of bytes written and any error.
func (c codec[T, V, K, M]) Encode(t T, w transport.Writer) (n int, err error) {
	n, err = c.serT.Marshal(t, w)
	if err != nil {
		err = fmt.Errorf(ErrorPrefix+"%w", err)
	}
	return
}

// Decode reads a value of type V from the given transport.Reader.
// Returns the decoded value, number of bytes read, and any error.
func (c codec[T, V, K, M]) Decode(r transport.Reader) (v V, n int, err error) {
	v, n, err = c.serV.Unmarshal(r)
	if err != nil {
		err = fmt.Errorf(ErrorPrefix+"%w", err)
	}
	return
}
