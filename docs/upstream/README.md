# Upstream documentation snapshot

This directory stores a repository snapshot of the official Capitalist API2 documentation:

`https://docs.capitalist.net/api/integration-api.html`

The HTML file is the source snapshot. It normalizes Cloudflare's dynamic `data-cfemail` values so repeated refreshes do not produce unrelated diffs. The Markdown file is generated from the same HTML and exists to make reviews and future diffs easier.

Last refreshed: 2026-05-28.

To refresh the snapshot:

```bash
scripts/update-upstream-docs.sh
git diff -- docs/upstream
```

The refresh script does not use API credentials and must not read or write private environment files.
