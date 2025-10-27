package codec

import (
	"errors"
	"testing"

	"github.com/cmd-stream/core-go"
	cmock "github.com/cmd-stream/core-go/testdata/mock"
	tmock "github.com/cmd-stream/transport-go/testdata/mock"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestServerCodec(t *testing.T) {
	t.Run("Encode", func(t *testing.T) {
		var (
			wantResult = cmock.NewResult()
			wantWriter = tmock.NewWriter()
			wantN      = 3
			wantErr    = errors.New("encode error")
			resultSer  = mock.NewSerializer[core.Result]().RegisterMarshal(
				func(result core.Result, w mus.Writer) (n int, err error) {
					asserterror.Equal[core.Result](result, wantResult, t)
					asserterror.Equal[mus.Writer](w, wantWriter, t)
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
				func(r mus.Reader) (cmd core.Cmd[any], n int, err error) {
					asserterror.Equal[mus.Reader](r, wantReader, t)
					return wantCmd, wantN, wantErr
				},
			)
			codec = NewServerCodec(cmdSer, nil)
		)
		cmd, n, err := codec.Decode(wantReader)
		asserterror.EqualError(err, wantErr, t)
		asserterror.Equal(n, wantN, t)
		asserterror.Equal[core.Cmd[any]](cmd, wantCmd, t)
	})
}
