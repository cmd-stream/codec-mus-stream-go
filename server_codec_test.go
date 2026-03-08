package codec

import (
	"errors"
	"testing"

	"github.com/cmd-stream/core-go"
	cmock "github.com/cmd-stream/core-go/test/mock"
	tmock "github.com/cmd-stream/transport-go/test/mock"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testutil/mock"
	asserterror "github.com/ymz-ncnk/assert/error"
)

func TestServerCodec(t *testing.T) {
	t.Run("Encoding should succeed", func(t *testing.T) {
		var (
			wantResult = cmock.NewResult()
			wantWriter = tmock.NewWriter()
			wantN      = 3
			wantErr    = errors.New("encode error")
			resultSer  = mock.NewSerializer[core.Result]().RegisterMarshal(
				func(result core.Result, w mus.Writer) (n int, err error) {
					asserterror.Equal[core.Result](t, result, wantResult)
					asserterror.Equal[mus.Writer](t, w, wantWriter)
					return wantN, wantErr
				},
			)
			codec = NewServerCodec[any](nil, resultSer)
		)
		n, err := codec.Encode(wantResult, wantWriter)
		asserterror.EqualError(t, err, wantErr)
		asserterror.Equal(t, n, wantN)
	})

	t.Run("Decoding should succeed", func(t *testing.T) {
		var (
			wantCmd    = cmock.NewCmd()
			wantN      = 3
			wantReader = tmock.NewReader()
			wantErr    = errors.New("decode error")
			cmdSer     = mock.NewSerializer[core.Cmd[any]]().RegisterUnmarshal(
				func(r mus.Reader) (cmd core.Cmd[any], n int, err error) {
					asserterror.Equal[mus.Reader](t, r, wantReader)
					return wantCmd, wantN, wantErr
				},
			)
			codec = NewServerCodec(cmdSer, nil)
		)
		cmd, n, err := codec.Decode(wantReader)
		asserterror.EqualError(t, err, wantErr)
		asserterror.Equal(t, n, wantN)
		asserterror.Equal[core.Cmd[any]](t, cmd, wantCmd)
	})
}
