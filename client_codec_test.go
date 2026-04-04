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

func TestClientCodec_Encode(t *testing.T) {
	var (
		wantN      = 0
		wantCmd    = cmock.NewCmd[any]()
		wantWriter = tmock.NewWriter()
		innerErr   = errors.New("encode error")
		wantErr    = fmt.Errorf(cdc.ErrorPrefix+"%w", innerErr)
		cmdSer     = mock.NewSerializer[core.Cmd[any]]()
	)
	cmdSer.RegisterMarshal(
		func(cmd core.Cmd[any], w mus.Writer) (n int, err error) {
			assertfatal.Equal[core.Cmd[any]](t, cmd, wantCmd)
			assertfatal.Equal[mus.Writer](t, w, wantWriter)
			return wantN, innerErr
		},
	)
	codec := cdc.NewClientCodec(cmdSer, nil)
	n, err := codec.Encode(wantCmd, wantWriter)
	assertfatal.EqualError(t, err, wantErr)
	assertfatal.Equal(t, n, wantN)
}

func TestClientCodec_Decode(t *testing.T) {
	var (
		wantResult = cmock.NewResult()
		wantN      = 4
		wantReader = tmock.NewReader()
		innerErr   = errors.New("decode error")
		wantErr    = fmt.Errorf(cdc.ErrorPrefix+"%w", innerErr)
		resultSer  = mock.NewSerializer[core.Result]()
	)
	resultSer.RegisterUnmarshal(
		func(r mus.Reader) (result core.Result, n int, err error) {
			assertfatal.Equal[mus.Reader](t, r, wantReader)
			return wantResult, wantN, innerErr
		},
	)
	codec := cdc.NewClientCodec[any](nil, resultSer)
	result, n, err := codec.Decode(wantReader)
	assertfatal.EqualError(t, err, wantErr)
	assertfatal.Equal(t, n, wantN)
	assertfatal.Equal[core.Result](t, result, wantResult)
}
