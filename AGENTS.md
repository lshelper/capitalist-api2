# Capitalist API2 Agent Notes

## Updating Against Official Documentation

When asked to reread the original official documentation link, compare changes, and update the code, follow this workflow.

Official documentation:

`https://docs.capitalist.net/api/integration-api.html`

1. Check repository state first.
   - Run `git status --short --branch`.
   - Do not overwrite unrelated local changes.
   - Note whether local commits are ahead of `origin/main`.

2. Refresh the stored upstream documentation snapshot.
   - Run `scripts/update-upstream-docs.sh`.
   - Review `git diff -- docs/upstream`.
   - Identify official documentation changes from the saved snapshot.
   - Keep `docs/upstream/integration-api.html` and `docs/upstream/integration-api.md` committed when the official docs changed.
   - Add a dated note to `docs/upstream/changes.md`.

3. Analyze what changed before editing code.
   - Compare new endpoints, request fields, response fields, auth/signature rules, callbacks, currencies, payment channels, limits, warnings, and examples.
   - Pay special attention to naming differences between payment channel types and currency codes.
   - Preserve official endpoint casing exactly.

4. Update shared repository documentation.
   - Update `docs/api-surface.md`.
   - Update the `Documentation sync status` date in root `README.md`.
   - Document important behavioral notes, especially raw-body signature rules and caching/expiry warnings.

5. Update all maintained clients when API surface changes.
   - TypeScript client.
   - Go client.
   - PHP client.
   - Bash client.
   - IntelliJ HTTP collection.
   - Keep each language's existing style and naming conventions.

6. Add or update examples.
   - Add HTTP `.http` requests for new endpoints.
   - Add Bash helpers for new endpoint groups.
   - Update README snippets where useful.
   - Add environment variables to `clients/http/http-client.env.json.example` when needed.

7. Add tests where the repo has an existing test pattern.
   - TypeScript: add `node:test` coverage for request paths and response parsing when practical.
   - Go: add `httptest` coverage for request paths and response parsing when practical.
   - PHP: at minimum run syntax checks and existing tests unless a local client test harness exists.
   - Bash: run syntax checks for all `.sh` files and existing signature tests.

8. Run verification before committing.
   - `scripts/update-upstream-docs.sh`
   - `git diff --check`
   - `cd clients/typescript && npm test`
   - `cd clients/go && go test ./...`
   - `cd clients/php && php -l src/Client.php && php tests/SignatureTest.php`
   - `cd clients/bash && find . -name '*.sh' -print0 | xargs -0 -n1 bash -n && bash tests/signature_test.sh`

9. Commit in meaningful chunks.
   - Prefer one commit for upstream documentation snapshot changes.
   - Prefer a separate commit for code/client/example updates.
   - Keep commits focused and reviewable.

10. Do a review pass after implementation.
    - Re-read the diff.
    - Check official endpoint names and casing.
    - Check URL/path escaping.
    - Check raw-body signature behavior for GET/POST/file requests.
    - Check tests cover the highest-risk behavior.
    - Report any residual test gaps.

11. Push only when asked.
    - Check `git status --short --branch`.
    - Push the current branch to `origin` if requested.
