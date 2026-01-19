## What's New and Useful in Go 1.25.0

If you are upgrading an existing module, set the toolchain version with:

- `go mod edit -go=1.25` — updates `go.mod` to target Go 1.25.0.
- Or install the toolchain from https://golang.org/dl/ to build with the Go 1.25.0 features and optimizations.

### Handy commands

- `go doc -http=:6060` — serves package documentation for the current module on a local web server.
- `go vet ./...` — now reports misplaced `sync.WaitGroup.Add` calls and suggests `net.JoinHostPort` instead of manual host/port string concatenation.

### Runtime and stdlib highlights

- Container-aware `GOMAXPROCS`: better respects CPU limits inside containers so workloads stay within their requested resources.
- `testing/synctest`: stabilized in Go 1.25 (experimental in 1.24) to help make concurrent tests deterministic.
  - Bubble: `synctest.Test` runs code in an isolated goroutine group.
  - Fake time: sleeps advance virtual time instead of blocking real time.
  - Quiescence: `synctest.Wait()` blocks until all goroutines in the bubble finish or are durably blocked.
- `encoding/json/v2`: offers new features and performance improvements over `encoding/json`. See https://pkg.go.dev/encoding/json/v2 for details.
- New `os.Root` methods for filesystem operations under a rooted view:
  - `Root.Chmod`, `Root.Chown`, `Root.Chtimes`, `Root.Lchown`
  - `Root.Link`, `Root.MkdirAll`, `Root.ReadFile`, `Root.ReadLink`
  - `Root.RemoveAll`, `Root.Rename`, `Root.Symlink`, `Root.WriteFile`
