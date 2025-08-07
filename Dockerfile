# Using `distroless` for increased security and smaller image size
# Using `base` as this app requires libc = https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md
# Using `nonroot` for better container security
# Using specific `sha256` for reproduceable builds
FROM gcr.io/distroless/base-debian12:nonroot@sha256:dca858878977f01b80685240ca2cd9c1bb38958e2ab6f4ed5698c0ea23307143

COPY bin/app /app
COPY bin/migrations /migrations

CMD ["/app", "migrations"]
