# noiferr

[![Go](https://github.com/a2not/noiferr/actions/workflows/go.yml/badge.svg)](https://github.com/a2not/noiferr/actions/workflows/go.yml)
[![License: MIT](https://img.shields.io/badge/github/license/a2not/noiferr)](https://opensource.org/licenses/MIT)

Static analysis tool for Go that checks if errors are properly handled with conditional branch


## Install
```
go get -u github.com/a2not/noiferr
```

## what this does

* check if the received `error` type is handled with conditional branch

## how to use it

```
make build
export PATH=$PATH:~/go/src/github.com/a2not/noiferr/cmd/noiferr
go vet -vettool=`which noiferr`
```
