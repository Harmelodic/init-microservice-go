# Using `distroless` for increased security and smaller image size
# Using `base` as this app requires libc = https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md
# Using `nonroot` for better container security
# Using specific `sha256` for reproduceable builds
FROM gcr.io/distroless/base-debian12:nonroot@sha256:4ae8d0163a6f04d96f36e41324d76f00744f0db7545b6d04039c9e6fa1df77f3

COPY bin/app /app
COPY bin/migrations /migrations

CMD ["/app", "migrations"]
