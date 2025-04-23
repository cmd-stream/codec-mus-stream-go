package cdc

import (
	"errors"
	"testing"

	"github.com/cmd-stream/base-go"
	bmock "github.com/cmd-stream/base-go/testdata/mock"
	tmock "github.com/cmd-stream/transport-go/testdata/mock"
	muss "github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestClientCodec(t *testing.T) {

	t.Run("Encode", func(t *testing.T) {
		var (
			wantCmd    = bmock.NewCmd()
			wantWriter = tmock.NewWriter()
			wantErr    = errors.New("encode error")
			cmdSer     = mock.NewSerializer[base.Cmd[any]]().RegisterMarshal(
				func(cmd base.Cmd[any], w muss.Writer) (n int, err error) {
					asserterror.Equal[base.Cmd[any]](cmd, wantCmd, t)
					asserterror.Equal[muss.Writer](w, wantWriter, t)
					return 0, wantErr
				},
			)
			codec = NewClientCodec[any](cmdSer, nil)
		)
		err := codec.Encode(wantCmd, wantWriter)
		asserterror.EqualError(err, wantErr, t)

	})

	t.Run("Decode", func(t *testing.T) {
		var (
			wantResult = bmock.NewResult()
			wantReader = tmock.NewReader()
			wantErr    = errors.New("decode error")
			resultSer  = mock.NewSerializer[base.Result]().RegisterUnmarshal(
				func(r muss.Reader) (result base.Result, n int, err error) {
					asserterror.Equal[muss.Reader](r, wantReader, t)
					return wantResult, 0, wantErr
				},
			)
			codec = NewClientCodec[any](nil, resultSer)
		)
		result, err := codec.Decode(wantReader)
		asserterror.Equal[base.Result](result, wantResult, t)
		asserterror.EqualError(err, wantErr, t)
	})

}

func TestServerCodec(t *testing.T) {

	t.Run("Encode", func(t *testing.T) {
		var (
			wantResult = bmock.NewResult()
			wantWriter = tmock.NewWriter()
			wantErr    = errors.New("encode error")
			resultSer  = mock.NewSerializer[base.Result]().RegisterMarshal(
				func(result base.Result, w muss.Writer) (n int, err error) {
					asserterror.Equal[base.Result](result, wantResult, t)
					asserterror.Equal[muss.Writer](w, wantWriter, t)
					return 0, wantErr
				},
			)
			codec = NewServerCodec[any](resultSer, nil)
		)
		err := codec.Encode(wantResult, wantWriter)
		asserterror.EqualError(err, wantErr, t)

	})

	t.Run("Decode", func(t *testing.T) {
		var (
			wantCmd    = bmock.NewCmd()
			wantReader = tmock.NewReader()
			wantErr    = errors.New("decode error")
			cmdSer     = mock.NewSerializer[base.Cmd[any]]().RegisterUnmarshal(
				func(r muss.Reader) (cmd base.Cmd[any], n int, err error) {
					asserterror.Equal[muss.Reader](r, wantReader, t)
					return wantCmd, 0, wantErr
				},
			)
			codec = NewServerCodec[any](nil, cmdSer)
		)
		cmd, err := codec.Decode(wantReader)
		asserterror.Equal[base.Cmd[any]](cmd, wantCmd, t)
		asserterror.EqualError(err, wantErr, t)
	})

}
