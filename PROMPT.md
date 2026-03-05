0a. Study `PLAN.md` to understand the current implementation status and next tasks.
0b. Study `AGENTS.md` for build, run, and validation rules.
0c. For reference, the application source code is in the root directory.
1. Your task is to implement functionality per the `PLAN.md` items. Choose the most important pending item. Before making changes, search the codebase (don't assume not implemented) to understand the current state.
2. After implementing functionality or resolving problems, run the validation commands listed in `AGENTS.md`. If functionality is missing, add it as per the project specifications (Pure Go, config.yaml, hardcoded paths).
3. When you discover issues or complete tasks, immediately update `PLAN.md` with your findings. Mark items as completed `[x]` or add new findings as unchecked `[ ]` items.
4. When the tests pass and the code is stable, `git add -A` then `git commit` with a message describing the changes. **You must commit at every loop iteration.**
5. **DO NOT push.** **DO NOT create git tags.**
99999. Important: When authoring documentation, capture the why — tests and implementation importance.
999999. Important: Single sources of truth, no migrations/adapters. If tests unrelated to your work fail, resolve them as part of the increment.
9999999. You may add extra logging if required to debug issues.
99999999. Keep `PLAN.md` current with learnings — future work depends on this to avoid duplicating efforts. Update especially after finishing your turn.
999999999. For any bugs you notice, resolve them or document them in `PLAN.md` even if it is unrelated to the current piece of work.
9999999999. Implement functionality completely. Placeholders and stubs waste efforts and time redoing the same work.
99999999999. If you find inconsistencies in the specifications then update the relevant documentation or `PLAN.md`.
999999999999. IMPORTANT: Keep `AGENTS.md` operational only — status updates and progress notes belong in `PLAN.md`. A bloated `AGENTS.md` pollutes every future loop's context.
9999999999999. Project Specifics:
    - No CLI flags. Configuration is via `config.yaml`.
    - Content path is hardcoded `./content`.
    - Output path is hardcoded `./output`.
    - Modes: `gocv` (CLI generate), `gocv serve` (HTTP server).
    - PDF Generation must be Pure Go (no external binaries).
