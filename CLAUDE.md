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

A small HTTP service (Go port of [no-as-a-service](https://github.com/hotheadhacker/no-as-a-service)) built on [Echo v5](https://echo.labstack.com) and scaffolded from [go-tiny-boilerplate](https://github.com/viilis/go-tiny-boilerplate). Entry point is `cmd/app/main.go`, which:

- loads `.env` via `godotenv` and reads `PORT`
- reads and JSON-unmarshals `reasons.json` (a flat array of ~1000 rejection strings) into memory at startup
- starts an Echo server with `Recover`, `CORS("*")`, and an in-memory `RateLimiter` (20 req burst) middleware
- exposes `GET /no`, which returns a random reason as `{"reason": ...}`

Layout:

- `cmd/app/` — application entry point
- `reasons.json` — flat JSON array of rejection strings, read at startup and served at random
- `http/no.http` — REST client request file for hitting the local server
- `tmp/` — compiled binary output (gitignored)
- `.env` / `.env.example` — environment variables loaded via `godotenv` (`PORT`)
- `cliff.toml` — changelog generation config using [git-cliff](https://git-cliff.org) with conventional commits
