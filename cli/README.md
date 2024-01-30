
# [![**Brooklyn**](https://brooklyn.apache.org/style/img/apache-brooklyn-logo-244px-wide.png)](http://brooklyn.apache.org/)

### Apache Brooklyn Client CLI

A command line client for [Apache Brooklyn](https://brooklyn.apache.org).

## Toolchain

The CLI tool is written in Go and should be obtained and built as a standard Go project. 
You will need the following tools to build it:

- Go (version 1.15 or higher).
- Maven (optional, used by the Brooklyn build process)


## Workspace Setup

Go is very particular about the layout of a source tree and the source repository, 
as it relies on this in the naming of packages.  

If you're familiar with Go and just want to develop the `br` tool itself you may simply work in your usual manner, 
using `go get github.com/apache/brooklyn-client/cli/br` and adding your own fork as a remote. 

`br` is built just like any other Go project. Dependencies are managed through Go modules.

## Compiling the code with Go for development purposes

Just use the regular Go build commands:

```bash
go build -o target/br ./br
```

The binary is now ready to use in `target/br`. 


## Testing 

The code includes a test script in the [test](test) directory. This deploys a Tomcat server on a location of your choice
and runs a number of tests against it, to verify that the br commands perform as expected.  To use this you must edit
the file "test_app.yaml" to change the location to your own value, and then invoke the test script like the following,
where the username and password need only be supplied if Brooklyn requires them:

```bash
sh test.sh  http://your-brooklyn-host:8081 myuser mypassword
```

Note, the tests are not yet comprehensive, and contributions are welcome.


## Building the code as part of the Brooklyn build process

For consistency with the other sub-projects of the overall [Brooklyn](https://github.com/apache/brooklyn) build, Maven
is used to perform the build when brooklyn-client is built as one of the sub-modules of Brooklyn, cross-compiling the code for a number of platform-architecture combinations.

Invoke the build script via Maven with one of 

  - ```mvn clean install```                                     build for all supported platforms
  - ```mvn -Dtarget=native clean install```                     build for the current platform
  - ```mvn -Dtarget=cross -Dos=OS -Darch=ARCH clean install```  build for platform with operating system OS and architecture ARCH

This builds the requested binaries into the `target/` directory, each in its own subdirectory with a name that includes 
the platform/architecture details, e.g. `bin/linux.386/br`. (When using this build process, the build script also writes the Go module cache into this directory.) The build installs a maven artifact to the maven repository,
consisting of a zip file containing all the binaries.  This artifact can be referenced in a POM as

```xml
<project>
    <groupId>org.apache.brooklyn</groupId>
    <artifactId>brooklyn-client-cli</artifactId>
    <classifier>bin</classifier>
    <type>zip</type>
    <version>1.2.0-SNAPSHOT</version>  <!-- BROOKLYN_VERSION -->
</project>

```

Most of the work is delegated to the `release/build.sh` script;
it is not normally necessary to use this, but if you need to know more,
try `release/build.sh -h` for more information.


## Usage

See instructions in the included [Runtime README](release/files/README) file.


----
Licensed to the Apache Software Foundation (ASF) under one 
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY 
KIND, either express or implied.  See the License for the 
specific language governing permissions and limitations
under the License.
