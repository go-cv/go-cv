# Implementation Plan

## Project Constraints
- [x] Pure Go binary (no external system dependencies for PDF).
- [x] Config via `config.yaml` only (no CLI flags).
- [x] Hardcoded paths: `./content` (input), `./output` (build artifacts).
- [x] Modes: `gocv` (CLI), `gocv serve` (Daemon).
- [x] Commit at every loop iteration. Do not push. Do not tag. (Workflow instruction)

## Current Status
- [x] Project backbone exists (HTTP server, graceful shutdown, config reading).
- [x] Markdown parsing logic implemented (using goldmark).
- [x] HTML Template engine integrated (Hugo-like theme selection).
- [x] PDF Generation implemented (using go-pdf/fpdf - Pure Go library).
- [x] CLI Mode (`gocv`) generates static files to `./output` and exits.
- [x] Serve Mode (`gocv serve`) hosts HTML and serves PDF on demand.
- [x] File Watcher implemented for live reload in Serve Mode (using fsnotify).
- [x] Dockerfile created for multi-stage build (re-created: was missing from repo).
- [x] go.mod dependencies corrected (marked direct deps properly).

## Active Task
- [x] Analyze existing backbone code and integrate Markdown parsing.
- [x] Integrate HTML template engine with theme selection.
- [x] Implement Serve Mode with file watching for live reload.
- [x] Create Dockerfile for multi-stage build.
- [x] All tasks complete - project ready for testing.

## Known Issues / Blockers
- [x] Identify best Pure Go PDF library that supports HTML/CSS (or define CSS subset).
  - Selected go-pdf/fpdf - generates PDFs programmatically from markdown AST.
  - Note: Not HTML-to-PDF conversion; walks goldmark AST and renders directly.
- [x] Current output is raw HTML fragments without full HTML document structure.

## Completed Log
- [x] Initial project structure defined.
- [x] Basic HTTP server and signal handling implemented.
- [x] CLI mode vs Serve mode distinction implemented (checks os.Args[1] for "serve").
- [x] Content reading from `./content` directory implemented.
- [x] Markdown to HTML conversion using goldmark library.
- [x] Theme system with templates/themes/{theme}/base.html structure.
- [x] Config.yaml theme setting (defaults to "default").
- [x] Output now generates complete HTML documents with styling.
