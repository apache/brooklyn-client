# Brooklyn CLI

## Compiling

1. Ensure your [$GOPATH](http://golang.org/cmd/go/#hdr-GOPATH_environment_variable) is set correctly,
   e.g. to some location where Go does its working, such as `~/gocode` .
2. Get and build the cli source code: `go get github.com/brooklyncentral/brooklyn-cli/br`
3. Run it from `$GOPATH/bin/br`
4. Thereafter if you want to do code changes, 
   link the `$GOPATH/src/github.com/brooklyncentral/brooklyn-cli`
   with the directory where you want to keep your git repositories.
   (TODO: clarify best practice for this, including how to combine
   it with a Brooklyn all-projects build)

## Running

First, log in to your Brooklyn instance with:

    $ ./br login URL [USER PASSWORD]

See the help command for info on all commands:

    $ ./br help

And for help on individual commands:

    $ ./br help COMMAND

