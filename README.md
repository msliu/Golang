Install GO on mac

http://sourabhbajaj.com/mac-setup/Go/README.html 

A valid GOPATH is required to use the `go get` command.
If $GOPATH is not specified, $HOME/go will be used by default:
  https://golang.org/doc/code.html#GOPATH

You may wish to add the GOROOT-based install location to your PATH:
  export PATH=$PATH:/usr/local/opt/go/libexec/bin
==> Summary
🍺  /usr/local/Cellar/go/1.10: 8,150 files, 336.9MB
VincentLius-MBP:Golang vincentliu$ go version
go version go1.10 darwin/amd64
VincentLius-MBP:Golang vincentliu$ 