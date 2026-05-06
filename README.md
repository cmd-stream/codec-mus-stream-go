# codec-mus-stream

[![Go Reference](https://pkg.go.dev/badge/github.com/cmd-stream/codec-mus-stream-go.svg)](https://pkg.go.dev/github.com/cmd-stream/codec-mus-stream-go)
[![GoReportCard](https://goreportcard.com/badge/cmd-stream/codec-mus-stream-go)](https://goreportcard.com/report/github.com/cmd-stream/codec-mus-stream-go)
[![codecov](https://codecov.io/gh/cmd-stream/codec-mus-stream-go/graph/badge.svg)](https://codecov.io/gh/cmd-stream/codec-mus-stream-go)

**codec-mus-stream** provides a MUS-based streaming codec for [cmd-stream](https://github.com/cmd-stream/cmd-stream-go).

## How To

```go
import (
  "github.com/cmd-stream/cmd-stream-go/core"
  cdcmuss "github.com/cmd-stream/codec-mus-stream-go"
)

var (
  serverCodec = cdcmuss.NewServerCodec(cmds.CmdMUS, results.ResultMUS)
  clientCodec = cdcmuss.NewClientCodec(cmds.CmdMUS, results.ResultMUS)
)
```

Here, `cmds.CmdMUS` is a MUS serializer for the `core.Cmd` interface, and
`results.ResultMUS` is a MUS serializer for the `core.Result` interface.

## Example

See the [hello-world](https://github.com/cmd-stream/examples-go/tree/main/hello-world) 
example for a full demonstration of `codec-mus-stream`.

## Fuzz Testing

This library represents a tiny wrapper around `CmdMUS` and `ResultMUS` serializers. 
To assist with robustness testing, [codec_fuzz_test.go](codec_fuzz_test.go) 
demonstrates how to fuzz test your custom serializers in combination with the 
codec.

To run fuzz tests:

```bash
./fuzz.sh 1m
```