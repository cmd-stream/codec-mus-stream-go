# codec-mus-stream-go

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
