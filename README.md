# codec-mus-stream-go

[![Go Reference](https://pkg.go.dev/badge/github.com/cmd-stream/codec-mus-stream-go.svg)](https://pkg.go.dev/github.com/cmd-stream/codec-mus-stream-go)
[![GoReportCard](https://goreportcard.com/badge/cmd-stream/codec-mus-stream-go)](https://goreportcard.com/report/github.com/cmd-stream/codec-mus-stream-go)
[![codecov](https://codecov.io/gh/cmd-stream/codec-mus-stream-go/graph/badge.svg)](https://codecov.io/gh/cmd-stream/codec-mus-stream-go)

**codec-mus-stream** provides a MUS-based streaming codec for [cmd-stream](https://github.com/cmd-stream/cmd-stream-go).

## How To

```go
import (
  "github.com/cmd-stream/cmd-stream-go/core"
  cdc "github.com/cmd-stream/codec-mus-stream-go"
)

var (
  serverCodec = cdc.NewServerCodec(cmds.CmdMUS, results.ResultMUS)
  clientCodec = cdc.NewClientCodec(cmds.CmdMUS, results.ResultMUS)
)
```

Here, `cmds.CmdMUS` is a MUS serializer for the `core.Cmd` interface, and
`results.ResultMUS` is a MUS serializer for the `core.Result` interface.

## Example

A full example of how to use **codec-mus-stream** can be found 
[here](https://github.com/cmd-stream/examples-go/tree/main/hello-world).
