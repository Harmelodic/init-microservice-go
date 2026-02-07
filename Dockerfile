# Using `distroless` for increased security and smaller image size
# Using `base` as this app requires libc = https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md
# Using `nonroot` for better container security
# Using specific `sha256` for reproduceable builds
FROM gcr.io/distroless/base-debian12:nonroot@sha256:107333192f6732e786f65df4df77f1d8bfb500289aad09540e43e0f7b6a2b816

COPY bin/app /app
COPY bin/migrations /migrations

CMD ["/app", "migrations"]
