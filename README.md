[![ci](https://github.com/sinlov-go/go-common-lib/actions/workflows/ci.yml/badge.svg)](https://github.com/sinlov-go/go-common-lib/actions/workflows/ci.yml)
[![license](https://img.shields.io/github/license/sinlov-go/go-common-lib)](https://github.com/sinlov-go/go-common-lib)
[![go mod version](https://img.shields.io/github/go-mod/go-version/sinlov-go/go-common-lib?label=go.mod)](https://github.com/sinlov-go/go-common-lib)
[![GoDoc](https://godoc.org/github.com/sinlov-go/go-common-lib?status.png)](https://godoc.org/github.com/sinlov-go/go-common-lib/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov-go/go-common-lib)](https://goreportcard.com/report/github.com/sinlov-go/go-common-lib)
[![codecov](https://codecov.io/gh/sinlov-go/go-common-lib/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov-go/go-common-lib)
[![GitHub release)](https://img.shields.io/github/v/release/sinlov-go/go-common-lib)](https://github.com/sinlov-go/go-common-lib/releases)

## for what

- this project used to github golang lib project

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/sinlov-go/go-common-lib)](https://github.com/sinlov-go/go-common-lib/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息

## depends

in go mod project

```bash
# warning use private git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "https://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q https://github.com/sinlov-go/go-common-lib.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/sinlov-go/go-common-lib
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/sinlov-go/go-common-lib | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## Features

- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## env

- minimum go version: go 1.18
- change `go 1.18`, `^1.18`, `1.18.10` to new go version

### libs

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.8.4  |
| https://github.com/sebdah/goldie    | v2.5.3  |

## usage

- use as `pkg`

```go
import (
    "github.com/sinlov-go/go-common-lib/pkg/date"
    "github.com/sinlov-go/go-common-lib/pkg/filepath_plus"
    "github.com/sinlov-go/go-common-lib/pkg/string_tools"
    "github.com/sinlov-go/go-common-lib/pkg/folder"
    "github.com/sinlov-go/go-common-lib/pkg/reflect_kit"
    "github.com/sinlov-go/go-common-lib/pkg/struct_kit"
)
```

# dev

```bash
# It needs to be executed after the first use or update of dependencies.
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run at env dev use cmd/main.go
$ make dev
```

- ci to fast check

```bash
# check style at local
$ make style

# run ci at local
$ make ci
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```
