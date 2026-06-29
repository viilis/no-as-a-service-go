# Tiny Go application boilerplate

## Usage

**Build and run:**
```sh
make build-run
```

**Run with live reload** (watches `*.go` files and restarts on change):
```sh
make watch
```
If `wgo` is not installed, the target will prompt you to install it automatically.

**Run tests:**
```sh
make test
```

**Lint:**
```sh
make lint
```
If `golangci-lint` is not installed, the target will prompt you to install it automatically.

## Under the hood
- [wgo - watcher-go](https://github.com/bokwoon95/wgo) - Live reload for Go apps
