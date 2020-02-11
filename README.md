[![CircleCI](https://circleci.com/gh/spatialcurrent/goprompt/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/goprompt/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/goprompt)](https://goreportcard.com/report/spatialcurrent/goprompt)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/goprompt?status.svg)](https://godoc.org/github.com/spatialcurrent/goprompt) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/goprompt/blob/master/LICENSE)

# goprompt

## Description

**goprompt** is a simple tool for prompting the user for input.

## Platforms

The following platforms are supported.  Pull requests to support other platforms are welcome!

| GOOS | GOARCH |
| ---- | ------ |
| darwin | amd64 |
| linux | amd64 |
| windows | amd64 |
| linux | arm64 |

## Releases

Find releases at [https://github.com/spatialcurrent/goprompt/releases](https://github.com/spatialcurrent/goprompt/releases).  You might want to rename your binary to just `goprompt` or `prompt`.  See the **Building** section below to build from scratch.

**Darwin**

- `goprompt_darwin_amd64` - CLI for Darwin on amd64 (includes `macOS` and `iOS` platforms)

**Linux**

- `goprompt_linux_amd64` - CLI for Linux on amd64
- `goprompt_linux_amd64` - CLI for Linux on arm64

**Windows**

- `goprompt_windows_amd64.exe` - CLI for Windows on amd64

## Usage

See the usage below or the following examples.

```shell
goprompt is a simple tool for prompting the user for input.

Usage:
  goprompt [--secret] [--json] [--question QUESTION]

Flags:
  -h, --help              help for goprompt
  -j, --json              validate input as JSON
  -q, --question string   the question for the prompt
  -s, --secret            use secret prompt
```

## Examples

```shell
...
mfa_serial=$(goprompt --secret --question "MFA SERIAL")
...
```

```shell
...
input=$(goprompt --secret --json --question "JSON Input")
echo "${input}" | jq .
...
```

## Building

Use `make help` to see help information for each target.

**CLI**

The `make build_cli` script is used to build executables for Linux and Windows.  Use `make install` for standard installation as a go executable.

**Changing Destination**

The default destination for build artifacts is `bin`, but you can change the destination with an environment variable.  For building on a Chromebook consider saving the artifacts in `/usr/local/go/bin`, e.g., `DEST=/usr/local/go/bin make build_cli`

## Testing

**CLI**

To run CLI testes use `make test_cli`, which uses [shUnit2](https://github.com/kward/shunit2).  If you recive a `shunit2:FATAL Please declare TMPDIR with path on partition with exec permission.` error, you can modify the `TMPDIR` environment variable in line or with `export TMPDIR=<YOUR TEMP DIRECTORY HERE>`. For example:

```
TMPDIR="/usr/local/tmp" make test_cli
```

**Go**

To run Go tests use `make test_go` (or `bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [ineffassign](https://github.com/gordonklaus/ineffassign), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

## Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/goprompt/blob/master/CONTRIBUTING.md) for how to get started.

## License

This work is distributed under the **MIT License**.  See **LICENSE** file.
