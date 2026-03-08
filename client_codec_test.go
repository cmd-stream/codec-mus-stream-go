package codec

import (
	"errors"
	"testing"

	"github.com/cmd-stream/core-go"
	cmocks "github.com/cmd-stream/core-go/test/mock"
	tmocks "github.com/cmd-stream/transport-go/test/mock"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testutil/mock"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestClientCodec(t *testing.T) {
	t.Run("Encoding should succeed", func(t *testing.T) {
		var (
			wantN      = 0
			wantCmd    = cmocks.NewCmd()
			wantWriter = tmocks.NewWriter()
			wantErr    = errors.New("encode error")
			cmdSer     = mock.NewSerializer[core.Cmd[any]]().RegisterMarshal(
				func(cmd core.Cmd[any], w mus.Writer) (n int, err error) {
					assertfatal.Equal[core.Cmd[any]](t, cmd, wantCmd)
					assertfatal.Equal[mus.Writer](t, w, wantWriter)
					return wantN, wantErr
				},
			)
			codec = NewClientCodec(cmdSer, nil)
		)
		n, err := codec.Encode(wantCmd, wantWriter)
		assertfatal.EqualError(t, err, wantErr)
		assertfatal.Equal(t, n, wantN)
	})

	t.Run("Decoding should succeed", func(t *testing.T) {
		var (
			wantResult = cmocks.NewResult()
			wantN      = 4
			wantReader = tmocks.NewReader()
			wantErr    = errors.New("decode error")
			resultSer  = mock.NewSerializer[core.Result]().RegisterUnmarshal(
				func(r mus.Reader) (result core.Result, n int, err error) {
					assertfatal.Equal[mus.Reader](t, r, wantReader)
					return wantResult, wantN, wantErr
				},
			)
			codec = NewClientCodec[any](nil, resultSer)
		)
		result, n, err := codec.Decode(wantReader)
		assertfatal.EqualError(t, err, wantErr)
		assertfatal.Equal(t, n, wantN)
		assertfatal.Equal[core.Result](t, result, wantResult)
	})
}
