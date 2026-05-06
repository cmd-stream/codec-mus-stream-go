package codec_test

import (
	"testing"

	cdctest "github.com/cmd-stream/codec-go/test"
	cdcmuss "github.com/cmd-stream/codec-mus-stream-go"
	"github.com/cmd-stream/codec-mus-stream-go/test"
)

func FuzzClientCodec_Decode(f *testing.F) {
	f.Add([]byte{0})
	f.Add([]byte{128, 1})

	f.Fuzz(func(t *testing.T, data []byte) {
		c := cdcmuss.NewClientCodec(test.CmdSer{}, test.ResultSer{})
		cdctest.FuzzDecode(c, data)
	})
}

func FuzzServerCodec_Decode(f *testing.F) {
	f.Add([]byte{0})
	f.Add([]byte{128, 1})

	f.Fuzz(func(t *testing.T, data []byte) {
		s := cdcmuss.NewServerCodec(test.CmdSer{}, test.ResultSer{})
		cdctest.FuzzDecode(s, data)
	})
}

func FuzzRoundTrip_Cmd(f *testing.F) {
	var (
		client = cdcmuss.NewClientCodec(test.CmdSer{}, test.ResultSer{})
		server = cdcmuss.NewServerCodec(test.CmdSer{}, test.ResultSer{})
	)

	f.Add(10)
	f.Fuzz(func(t *testing.T, x int) {
		cmd := test.Cmd1{X: x}
		cdctest.VerifyRoundTripCmd(t, client, server, cmd)
	})
}

func FuzzRoundTrip_Result(f *testing.F) {
	var (
		client = cdcmuss.NewClientCodec(test.CmdSer{}, test.ResultSer{})
		server = cdcmuss.NewServerCodec(test.CmdSer{}, test.ResultSer{})
	)

	f.Add(10, true)
	f.Fuzz(func(t *testing.T, x int, lastOne bool) {
		res := test.Result1{X: x, L: lastOne}
		cdctest.VerifyRoundTripResult(t, client, server, res)
	})
}
