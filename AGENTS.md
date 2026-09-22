# AGENTS.md

This file provides guidance to the AI agent when working with code in this repository.

## Project Overview

Go library package providing a goroutine job manager with panic recovery and graceful shutdown.
Module: `github.com/go-meridian/job` — package name is `jobmgr` (not `job`).

## Key Facts

- **Library only** — no `main` package, no binary output, no `cmd/` directory
- **Dependencies**: `github.com/go-meridian/logger` (zap-based), with local replace `../logger` in go.mod
- **Go version**: 1.27.1 (set in go.mod)
- **Singleton pattern**: `Init()` + `Mgr()` via `sync.Once`; `Init` is safe to call multiple times; **auto-init** on first use if `Init` not called
- **Logging**: uses `github.com/go-meridian/logger` package (`logger.Get()` returns `*zap.Logger`); no custom Logger interface

## Build & Verify

```bash
go build ./...
go vet ./...
```

No vendor directory, no tests yet. Add `_test.go` files alongside source.

## Code Conventions

- All doc comments and error messages in Chinese
- `PanicHandler` callback receives the full stack string; callers decide how to report
- `AddJob` captures panics; `AddJobNaked` does not — this is intentional, not a bug
- `StopAll` is safe to call multiple times (`sync.Once`); only logs `Error` on timeout (with remaining task count); package-level `StopAll(timeout)` also available
- `AddJob` panic output priority: `PanicHandler` > `logger.Error`（无 logger 时静默跳过）
- `Running()` returns current active goroutine count via `atomic.Int64`
