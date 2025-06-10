package cdc

import (
	"errors"
	"testing"

	"github.com/cmd-stream/core-go"
	cmock "github.com/cmd-stream/core-go/testdata/mock"
	tmock "github.com/cmd-stream/transport-go/testdata/mock"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestClientCodec(t *testing.T) {

	t.Run("Encode", func(t *testing.T) {
		var (
			wantN      = 0
			wantCmd    = cmock.NewCmd()
			wantWriter = tmock.NewWriter()
			wantErr    = errors.New("encode error")
			cmdSer     = mock.NewSerializer[core.Cmd[any]]().RegisterMarshal(
				func(cmd core.Cmd[any], w muss.Writer) (n int, err error) {
					asserterror.Equal[core.Cmd[any]](cmd, wantCmd, t)
					asserterror.Equal[muss.Writer](w, wantWriter, t)
					return wantN, wantErr
				},
			)
			codec = NewClientCodec[any](cmdSer, nil)
		)
		n, err := codec.Encode(wantCmd, wantWriter)
		asserterror.EqualError(err, wantErr, t)
		asserterror.Equal(n, wantN, t)

	})

	t.Run("Decode", func(t *testing.T) {
		var (
			wantResult = cmock.NewResult()
			wantN      = 4
			wantReader = tmock.NewReader()
			wantErr    = errors.New("decode error")
			resultSer  = mock.NewSerializer[core.Result]().RegisterUnmarshal(
				func(r muss.Reader) (result core.Result, n int, err error) {
					asserterror.Equal[muss.Reader](r, wantReader, t)
					return wantResult, wantN, wantErr
				},
			)
			codec = NewClientCodec[any](nil, resultSer)
		)
		result, n, err := codec.Decode(wantReader)
		asserterror.EqualError(err, wantErr, t)
		asserterror.Equal(n, wantN, t)
		asserterror.Equal[core.Result](result, wantResult, t)
	})

}

func TestServerCodec(t *testing.T) {

	t.Run("Encode", func(t *testing.T) {
		var (
			wantResult = cmock.NewResult()
			wantWriter = tmock.NewWriter()
			wantN      = 3
			wantErr    = errors.New("encode error")
			resultSer  = mock.NewSerializer[core.Result]().RegisterMarshal(
				func(result core.Result, w muss.Writer) (n int, err error) {
					asserterror.Equal[core.Result](result, wantResult, t)
					asserterror.Equal[muss.Writer](w, wantWriter, t)
					return wantN, wantErr
				},
			)
			codec = NewServerCodec[any](nil, resultSer)
		)
		n, err := codec.Encode(wantResult, wantWriter)
		asserterror.EqualError(err, wantErr, t)
		asserterror.Equal(n, wantN, t)
	})

	t.Run("Decode", func(t *testing.T) {
		var (
			wantCmd    = cmock.NewCmd()
			wantN      = 3
			wantReader = tmock.NewReader()
			wantErr    = errors.New("decode error")
			cmdSer     = mock.NewSerializer[core.Cmd[any]]().RegisterUnmarshal(
				func(r muss.Reader) (cmd core.Cmd[any], n int, err error) {
					asserterror.Equal[muss.Reader](r, wantReader, t)
					return wantCmd, wantN, wantErr
				},
			)
			codec = NewServerCodec[any](cmdSer, nil)
		)
		cmd, n, err := codec.Decode(wantReader)
		asserterror.EqualError(err, wantErr, t)
		asserterror.Equal(n, wantN, t)
		asserterror.Equal[core.Cmd[any]](cmd, wantCmd, t)
	})

}
