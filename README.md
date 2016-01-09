# synacor-go

This is an implementation of the
[Synacor challenge](https://challenge.synacor.com/) using Golang.

## running the project

Assuming you have Go [installed](https://golang.org/dl/):

    go get github.com/nmuth/synacor-go
    cd $GOPATH/github.com/nmuth/synacor-go
    go build
    ./synacor-go data/challenge.bin

You can execute arbitrary files with `synacor-go $file`, so go nuts. Included is
a Ruby script for generating a "Hello, world!" binary.
