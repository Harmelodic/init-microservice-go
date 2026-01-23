# Using `distroless` for increased security and smaller image size
# Using `base` as this app requires libc = https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md
# Using `nonroot` for better container security
# Using specific `sha256` for reproduceable builds
FROM gcr.io/distroless/base-debian12:nonroot@sha256:748987b77724fc1a9cf1acc8afb7f2cecff8cd91f24f43423647f1896eb7262f

COPY bin/app /app
COPY bin/migrations /migrations

CMD ["/app", "migrations"]
