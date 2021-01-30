# Cards Deck with GO

We have a generated code in file suit_string.go using golib [stringer](https://golang.org/x/tools/cmd/stringer).

suit_string file has been created for you, if you happen to lose, follow the steps down below to generate a new one.

If you have created this project using modules and outside your $GOPATH, please do the following to be able to use the stringer lib

```bash
go mod init
export GOBIN=$PWD/bin
export PATH=$GOBIN:$PATH
```

Inside your go file that you want to generate code from, add the following at the top of you file

```go
//go:generate stringer -type=Suit,Rank
```

Then run go generate

```bash
go generate
```
