# logger
> 这只是一个handler的实现

## install

```bash
go get -u github.com/ndsky1003/sloghandler
```


#### usage
```golang
	handler := sloghandler.NewFastTextHandler(sloghandler.Options().SetAddSource(true).SetReplaceAttr(func(groups []string, a slog.Attr) slog.Attr {
//在这里修改指定的打印逻辑,比如passwd,用******替换
		if a.Key == "a" {
			a.Value = slog.StringValue("hello")
		}
		return a
	}).SetExtractorAttr(func(ctx context.Context, r *slog.Record) {
// 根据项目在这里注入traceid
		r.AddAttrs(slog.String("extracted", "value"))
	}).SetForceDebugFn(func(ctx context.Context) bool {
//强制某些人,不被打印等级控制
		return true
	}))
	handler.SetAddSource(false) //动态修改打印行号
	handler.SetLevel(slog.LevelError) //动态修改打印逻辑
```
