# codec-mus-stream-go

**codec-mus-stream-go** provides a MUS-based streaming codec for [cmd-stream-go](https://github.com/cmd-stream/cmd-stream-go).

## Example

```go
var (
  serverCodec = codec.NewServerCodec(cmds.CmdMUS, results.ResultMUS)
  clientCodec = codec.NewClientCodec(cmds.CmdMUS, results.ResultMUS)
)
```

Here, `cmds.CmdMUS` is a MUS serializer for the `core.Cmd` interface, and
`results.ResultMUS` is a MUS serializer for the `core.Result` interface.
