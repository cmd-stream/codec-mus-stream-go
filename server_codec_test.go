package codec_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/cmd-stream/cmd-stream-go/core"
	cmock "github.com/cmd-stream/cmd-stream-go/test/mock/core"
	tmock "github.com/cmd-stream/cmd-stream-go/test/mock/transport"
	cdc "github.com/cmd-stream/codec-mus-stream-go"
	"github.com/mus-format/mus-stream-go"
	"github.com/mus-format/mus-stream-go/test/mock"
	assertfatal "github.com/ymz-ncnk/assert/fatal"
)

func TestServerCodec_Encode(t *testing.T) {
	var (
		wantResult = cmock.NewResult()
		wantWriter = tmock.NewWriter()
		wantN      = 3
		innerErr   = errors.New("encode error")
		wantErr    = fmt.Errorf(cdc.ErrorPrefix+"%w", innerErr)
		resultSer  = mock.NewSerializer[core.Result]()
	)
	resultSer.RegisterMarshal(
		func(result core.Result, w mus.Writer) (n int, err error) {
			assertfatal.Equal[core.Result](t, result, wantResult)
			assertfatal.Equal[mus.Writer](t, w, wantWriter)
			return wantN, innerErr
		},
	)
	codec := cdc.NewServerCodec[any](nil, resultSer)
	n, err := codec.Encode(wantResult, wantWriter)
	assertfatal.EqualError(t, err, wantErr)
	assertfatal.Equal(t, n, wantN)
}

func TestServerCodec_Decode(t *testing.T) {
	var (
		wantCmd    = cmock.NewCmd[any]()
		wantN      = 3
		wantReader = tmock.NewReader()
		innerErr   = errors.New("decode error")
		wantErr    = fmt.Errorf(cdc.ErrorPrefix+"%w", innerErr)
		cmdSer     = mock.NewSerializer[core.Cmd[any]]()
	)
	cmdSer.RegisterUnmarshal(
		func(r mus.Reader) (cmd core.Cmd[any], n int, err error) {
			assertfatal.Equal[mus.Reader](t, r, wantReader)
			return wantCmd, wantN, innerErr
		},
	)
	codec := cdc.NewServerCodec(cmdSer, nil)
	cmd, n, err := codec.Decode(wantReader)
	assertfatal.EqualError(t, err, wantErr)
	assertfatal.Equal(t, n, wantN)
	assertfatal.Equal[core.Cmd[any]](t, cmd, wantCmd)
}
