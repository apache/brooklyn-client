
# [![**Brooklyn**](https://brooklyn.apache.org/style/img/apache-brooklyn-logo-244px-wide.png)](http://brooklyn.apache.org/)

### Apache Brooklyn Client CLI Sub-Project

A command line client for [Apache Brooklyn](https://brooklyn.apache.org).

## Toolchain

You will need the following tools to build the CLI:
- Go (min version 1.5.1), with full cross-compiler support (see https://golang.org/dl).
  On Mac, if using Homebrew, use "brew install go --with-cc-all"
- godep (see https://github.com/tools/godep)

Optional:
- Maven (used by the Brooklyn build process)


## Build Pre-Requisites

- Ensure your [$GOPATH](http://golang.org/cmd/go/#hdr-GOPATH_environment_variable) is set correctly 
  to a suitable location for your Go code.
- Install Brooklyn CLI and dependencies (note the "-d" parameter, which instructs Go to download the files but not
  build the code).  
`go get -d github.com/apache/brooklyn-client/br`  
    
    
## A note on dependency management

The CLI has a small number of dependencies, notably on codegansta/cli.  To manage the version of dependencies, the CLI
code currently uses godep.  When contributing to the CLI it is important to be aware of the distinction.  To avoid
potentially bringing in new versions of dependencies, use "godep go" to build the code using the dependencies
saved in br/Godeps.  Alternatively, to bring in the latest versions of the dependencies, build with "go get", but in
that case remember to update the dependencies of the project using "godep save" along with your commit.

## Compiling the code with Go for development purposes


### Using saved dependencies
As Go dependendencies for godep are held in the main package directory ("br"), you need to build from that directory,
using godep:

```bash
cd $GOPATH/src/github.com/apache/brooklyn-client/br
godep go install 
```
This will build the "br" executable into $GOPATH/bin

### Updating the dependencies

To use the latest published versions of the dependencies simply use 
```bash
go get github.com/apache/brooklyn-client/br
```

When the code is ready to be committed, first update the saved dependencies with
```bash
cd $GOPATH/src/github.com/apache/brooklyn-client/br
godep save
```

## Testing 

The code includes a test script in the [test](test) directory. This deploys a Tomcat server on a location of your choice
and runs a number of tests against it, to verify that the br commands perform as expected.  To use this you must edit
the file "test_app.yaml" to change the location to your own value, and then invoke the test script like the following,
where the username and password need only be supplied if Brooklyn requires them:

```bash
    $ sh test.sh  http://your-brooklyn-host:8081 myuser mypassword
    exit 0
```

Note, the tests are not yet comprehensive, and contributions are welcome.

## Building the code for release

Either:
- Use the build script in the "release" folder directly (see its usage for details), or
- Invoke the build script via Maven with one of 
  - ```mvn clean install```                                     build for all supported platforms
  - ```mvn -Dtarget=native clean install```                     build for the current platform
  - ```mvn -Dtarget=cross -Dos=OS -Darch=ARCH clean install```  build for platform with operating system OS and architecture ARCH

This builds the requested binaries into the "target" directory, each with a file name that includes the version,
timestamp, and architecture details, e.g. br.0.9.0.20151218-195906.linux.amd64, and installs a zip file containing them
all as a maven artifact.  To run any of these as "br" of course you will need to create an alias or soft link.

## Running

See instructions in the included [Runtime README](README) file.

