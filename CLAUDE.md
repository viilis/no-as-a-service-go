# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```sh
make build-run   # compile to tmp/main and execute
make watch       # live reload on *.go changes via wgo (prompts to install if missing)
make test        # go test ./...
make lint        # golangci-lint run ./... (prompts to install if missing)
```

## Architecture

Minimal Go application template. Entry point is `cmd/app/main.go`.

- `cmd/app/` — application entry point
- `tmp/` — compiled binary output (gitignored)
- `.env` / `.env.example` — environment variables loaded via `godotenv`
- `cliff.toml` — changelog generation config using [git-cliff](https://git-cliff.org) with conventional commits

The module name `tiny-go-template` in `go.mod` should be renamed when using this as a template base.
