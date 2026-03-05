# Implementation Plan

## Project Constraints
- [ ] Pure Go binary (no external system dependencies for PDF).
- [ ] Config via `config.yaml` only (no CLI flags).
- [ ] Hardcoded paths: `./content` (input), `./output` (build artifacts).
- [ ] Modes: `gocv` (CLI), `gocv serve` (Daemon).
- [ ] Commit at every loop iteration. Do not push. Do not tag.

## Current Status
- [x] Project backbone exists (HTTP server, graceful shutdown, config reading).
- [ ] Markdown parsing logic implemented.
- [ ] HTML Template engine integrated (Hugo-like theme selection).
- [ ] PDF Generation implemented (Pure Go library selected and integrated).
- [ ] CLI Mode (`gocv`) generates static files to `./output` and exits.
- [ ] Serve Mode (`gocv serve`) hosts HTML and serves PDF on demand.
- [ ] File Watcher implemented for live reload in Serve Mode.
- [ ] Dockerfile created for multi-stage build.

## Active Task
- [ ] Analyze existing backbone code and integrate Markdown parsing.

## Known Issues / Blockers
- [ ] Identify best Pure Go PDF library that supports HTML/CSS (or define CSS subset).

## Completed Log
- [x] Initial project structure defined.
- [x] Basic HTTP server and signal handling implemented.
