# Using `distroless` for increased security and smaller image size
# Using `base` as this app requires libc = https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md
# Using `nonroot` for better container security
# Using specific `sha256` for reproduceable builds
FROM gcr.io/distroless/base-debian12:nonroot@sha256:956eee19d77039968b05209dce21e43c84fb2bae7644a2b0546b36996c96e305

COPY bin/app /app
COPY bin/migrations /migrations

CMD ["/app", "migrations"]
