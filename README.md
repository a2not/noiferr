# noiferr

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
