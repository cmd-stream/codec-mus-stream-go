package test

import (
	"context"

	"github.com/cmd-stream/cmd-stream-go/core"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/ord"
	"github.com/mus-format/mus-stream-go/varint"
)

// Cmd1 is a simple command for testing.
type Cmd1 struct {
	X int
}

func (c Cmd1) Exec(ctx context.Context, receiver any, proxy core.Proxy) error {
	return nil
}

// Result1 is a simple result for testing.
type Result1 struct {
	X int
	L bool
}

func (r Result1) LastOne() bool {
	return r.L
}

// CmdSer is a MUS serializer for Cmd1.
type CmdSer struct{}

func (s CmdSer) Marshal(v core.Cmd[any], w mus.Writer) (n int, err error) {
	return varint.Int.Marshal(v.(Cmd1).X, w)
}

func (s CmdSer) Unmarshal(r mus.Reader) (v core.Cmd[any], n int, err error) {
	x, n1, err := varint.Int.Unmarshal(r)
	if err != nil {
		return nil, n1, err
	}
	return Cmd1{X: x}, n1, nil
}

func (s CmdSer) Size(v core.Cmd[any]) int {
	return varint.Int.Size(v.(Cmd1).X)
}

func (s CmdSer) Skip(r mus.Reader) (n int, err error) {
	return varint.Int.Skip(r)
}

// ResultSer is a MUS serializer for Result1.
type ResultSer struct{}

func (s ResultSer) Marshal(v core.Result, w mus.Writer) (n int, err error) {
	n1, err := varint.Int.Marshal(v.(Result1).X, w)
	if err != nil {
		return n1, err
	}
	n2, err := ord.Bool.Marshal(v.(Result1).L, w)
	return n1 + n2, err
}

func (s ResultSer) Unmarshal(r mus.Reader) (v core.Result, n int, err error) {
	x, n1, err := varint.Int.Unmarshal(r)
	if err != nil {
		return nil, n1, err
	}
	l, n2, err := ord.Bool.Unmarshal(r)
	if err != nil {
		return nil, n1 + n2, err
	}
	return Result1{X: x, L: l}, n1 + n2, nil
}

func (s ResultSer) Size(v core.Result) int {
	return varint.Int.Size(v.(Result1).X) + ord.Bool.Size(v.(Result1).L)
}

func (s ResultSer) Skip(r mus.Reader) (n int, err error) {
	n1, err := varint.Int.Skip(r)
	if err != nil {
		return n1, err
	}
	n2, err := ord.Bool.Skip(r)
	return n1 + n2, err
}
