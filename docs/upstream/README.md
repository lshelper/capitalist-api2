# Upstream documentation snapshot

This directory stores a repository snapshot of the official Capitalist API2 documentation:

`https://docs.capitalist.net/api/integration-api.html`

The raw HTML file is the source snapshot. The Markdown file is generated from the same HTML and exists to make reviews and future diffs easier.

Last refreshed: 2026-05-28.

To refresh the snapshot:

```bash
scripts/update-upstream-docs.sh
git diff -- docs/upstream
```

The refresh script does not use API credentials and must not read or write private environment files.
