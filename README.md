# Install Golang with Homebrew:

  A good reference to introduce this intallation of GO lang with Homebrew.

      http://sourabhbajaj.com/mac-setup/Go/README.html

  $ brew update
  $ brew install golang

  Environment Setup:
  - add those lines into your .profile or .bashrc to export the required variables
	export GOPATH=$HOME/go-workspace # don't forget to change your path correctly!
	export GOROOT=/usr/local/opt/go/libexec
	export PATH=$PATH:$GOPATH/bin
	export PATH=$PATH:$GOROOT/bin

  - check go version
    $go version
