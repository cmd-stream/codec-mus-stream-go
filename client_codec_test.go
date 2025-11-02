package codec

import (
	"errors"
	"testing"

	"github.com/cmd-stream/core-go"
	cmocks "github.com/cmd-stream/testkit-go/mocks/core"
	tmocks "github.com/cmd-stream/testkit-go/mocks/transport"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/testdata/mock"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestClientCodec(t *testing.T) {
	t.Run("Encode", func(t *testing.T) {
		var (
			wantN      = 0
			wantCmd    = cmocks.NewCmd()
			wantWriter = tmocks.NewWriter()
			wantErr    = errors.New("encode error")
			cmdSer     = mock.NewSerializer[core.Cmd[any]]().RegisterMarshal(
				func(cmd core.Cmd[any], w mus.Writer) (n int, err error) {
					assertfatal.Equal[core.Cmd[any]](cmd, wantCmd, t)
					assertfatal.Equal[mus.Writer](w, wantWriter, t)
					return wantN, wantErr
				},
			)
			codec = NewClientCodec(cmdSer, nil)
		)
		n, err := codec.Encode(wantCmd, wantWriter)
		assertfatal.EqualError(err, wantErr, t)
		assertfatal.Equal(n, wantN, t)
	})

	t.Run("Decode", func(t *testing.T) {
		var (
			wantResult = cmocks.NewResult()
			wantN      = 4
			wantReader = tmocks.NewReader()
			wantErr    = errors.New("decode error")
			resultSer  = mock.NewSerializer[core.Result]().RegisterUnmarshal(
				func(r mus.Reader) (result core.Result, n int, err error) {
					assertfatal.Equal[mus.Reader](r, wantReader, t)
					return wantResult, wantN, wantErr
				},
			)
			codec = NewClientCodec[any](nil, resultSer)
		)
		result, n, err := codec.Decode(wantReader)
		assertfatal.EqualError(err, wantErr, t)
		assertfatal.Equal(n, wantN, t)
		assertfatal.Equal[core.Result](result, wantResult, t)
	})
}
