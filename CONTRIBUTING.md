# Contributing

## Go project

Prerequisites:

- Go is installed
- `golangci-lint` is installed

### Running tests

```shell
make test
```

### Building an executable

```shell
make build
# or just
make
```

## Docs (MkDocs)

Prerequisites: Python 3 installed

Setup a Python virtual environment:

```shell
python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt
# exit venv with:
deactivate
```

Run MkDocs locally:

```shell
mkdocs serve
```
