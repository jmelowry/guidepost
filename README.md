# Guidepost

**Guidepost** is a lightweight developer tool for surfacing useful metadata about apps running in Kubernetes. It exposes links to logs, dashboards, CI pipelines, docs, and more — helping developers quickly orient themselves without digging through documentation.

---

## What it does

- Reads per-app JSON metadata files from S3
- Caches metadata in a local SQLite database
- Serves a simple HTTP API for querying metadata
- Includes a CLI for retrieving and opening app links
- Supports manual sync of individual apps from S3
- Allows templating of link formats to avoid repeating base URLs

---

## Example metadata (stored in S3)

Each app gets a single JSON file (e.g., `guidepost-metadata/my-app.json`):

```json
{
  "name": "my-app",
  "repo": "https://github.com/my-org/my-app",
  "links": {
    "logs": "/query?service=my-app",
    "dashboard": "/d/my-app",
    "docs": "/wiki/my-app"
  },
  "owner": "my-team@example.com"
}
```

The base URLs for logs, dashboards, docs, etc., are defined on the server side to reduce duplication.

---

## API (MVP)

- `GET /apps` – List all known apps
- `GET /apps/:name` – Get metadata for a specific app
- `POST /apps/:name/sync` – Force metadata reload from S3 (rate limited)

> ❗Note: We use the term "apps" instead of "services" to keep things open-ended — any workload exposing functionality can be registered.

---

## CLI (WIP)

```bash
guidepost info my-app                   # Show metadata for an app
guidepost get logs my-app              # Get logs link for an app
guidepost get dashboard my-app         # Get dashboard link
guidepost sync my-app                  # Trigger a sync from S3
```

---

## Requirements

- Go 1.21+
- Access to an S3 bucket for metadata (or local filesystem for dev/test)
- Kubernetes (for deploying the server and using RBAC with the CLI)

---

## Roadmap

- [ ] Load metadata from S3
- [ ] Cache in SQLite
- [ ] Serve HTTP API
- [ ] CLI with `info`, `get`, and `sync` commands
- [ ] Support for local filesystem as metadata source
- [ ] Rate limiting for manual syncs
- [ ] JSON schema validation
- [ ] Templated base URLs for link rendering

---

## License

MIT

---

## Contributing

This project is in early development. If you're interested in using or contributing to Guidepost, feel free to open an issue or start a discussion.
