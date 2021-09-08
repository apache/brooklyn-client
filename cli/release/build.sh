#!/usr/bin/env bash
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#  http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

echo "Setting up go environment"
go env -w GO111MODULE=auto
export GO386='softfloat'
echo "GO386=$GO386"
#
# Constants
#
OSVALUES="darwin freebsd linux netbsd openbsd windows"
ARCHVALUES="386 amd64"
BRNAME="br"
PROJECT_PACKAGE="github.com/apache/brooklyn-client/cli"
CLI_EXECUTABLE="${PROJECT_PACKAGE}/${BRNAME}"

START_TIME=$(date +%s)

#
# Globals
#
os=""
arch=""
all=""
outdir=`pwd`/target
sourcedir=`pwd`
label=""
timestamp=""

builds=(
  darwin/amd64
  linux/386
  windows/386
)

show_help() {
	echo "Usage:	$0 [-d <OUTPUTDIR>] [-l <LABEL>] [-t] -s <SOURCEDIR>"
	echo "	$0 -o <OS> -a <ARCH> [-d <DIRECTORY>] [-l <LABEL>] [-t] -s <SOURCEDIR>"
	echo "	$0 -A [-d <OUTPUTDIR>] [-l <LABEL>] [-t] -s <SOURCEDIR>"
	echo "	$0 -h"
	echo
		cat <<-EOH
	 -A  Build for default OS/ARCH combinations
	 -a  Set ARCH to build for
	 -d  Set output directory
	 -h  Show help
	 -l  Set label text for including in filename
	 -o  Set OS to build for
	 -t  Set timestamp for including in filename
	 -s  Source directory

EOH

	echo $OSVALUES | awk 'BEGIN{printf("Supported OS:\n")};{for(i=1;i<=NF;i++){printf("\t%s\n",$i)}}'
	echo $ARCHVALUES | awk 'BEGIN{printf("Supported ARCH:\n")};{for(i=1;i<=NF;i++){printf("\t%s\n",$i)}}'
	echo Default build:
	for build in "${builds[@]}" ; do
	    printf "\t%s\n" $build
	done
}

while [ $# -gt 0 ]; do
	case $1 in 
	-h|help)
		show_help
		exit 0
		;;
	-d)
		if [ $# -lt 2 ]; then
			show_help
			echo "Value for OUTPUTDIR must be provided"
			exit 1
		fi
		outdir="$2"
		shift 2
		;;
	-s)
		if [ $# -lt 2 ]; then
			show_help
			echo "Value for SOURCEDIR must be provided"
			exit 1
		fi
		sourcedir="$2"
		shift 2
		;;
	-o)
		if [ $# -lt 2 ]; then
			show_help
			echo "Value for OS must be provided"
			exit 1
		fi
		os="$2"
		shift 2
		;;
	-a)
		if [ $# -lt 2 ]; then
			show_help
			echo "Value for ARCH must be provided"
			exit 1
		fi
		arch="$2"
		shift 2
		;;
	-A)
		all="all"
		shift 1
		;;
	-l)
		if [ $# -lt 2 ]; then
			show_help
			echo "Value for LABEL must be provided"
			exit 1
		fi
		label=".$2"
		shift 2
		;;
	-t)
		timestamp=`date +.%Y%m%d-%H%M%S`
		shift
		;;
	*)
		show_help
		echo "Unrecognised parameter: $1"
		exit 1
		;;
	esac
done

echo "Starting build.sh (brooklyn-client go build script)"

# Set GOPATH to ${outdir} so that Go module cache gets put there.
export GOPATH="${outdir}"

# set GOCACHE to avoid
# "failed to initialize build cache at /.cache/go-build: mkdir /.cache: permission denied"
# on CI builds
export GOCACHE="${outdir}/.cache/go-build"

#
# Test if go is available
#
if ! command -v go >/dev/null 2>&1 ; then
  cat 1>&2 << \
--MARKER--

ERROR: Go language binaries not found (running "go")

The binaries for Go must be installed to build the brooklyn-client CLI.
See golang.org for more information, or run maven with '-Dno-go-client' to skip.

--MARKER--
  exit 1
fi

GO_VERSION=`go version | awk '{print $3}'`
echo GO_VERSION is ${GO_VERSION}
GO_V=`echo $GO_VERSION | sed 's/^go1\.\([0-9][0-9]*\).*/\1/'`
# test if not okay so error shows if regex above not matched
if ! (( "$GO_V" >= 15 )) ; then
  cat 1>&2 << \
--MARKER--

ERROR: Incompatible Go language version: $GO_VERSION

Go version 1.15 or higher is required to build the brooklyn-client CLI.
See golang.org for more information, or run maven with '-Dno-go-client' to skip.

--MARKER--
  exit 1
fi

if !(( "$GO_V" >= 16 )) ; then
  export GO386=387
  echo "Updated to use GO386=$GO386 due GO version compatibility as $GO_GO_VERSION does not support 'softfloat'"
fi

mkdir -p $outdir

# Disable use of C code modules (causes problems with cross-compiling)
export CGO_ENABLED=0

if [ -n "$all" -a \( -n "$os" -o -n "$arch" \) ]; then
	show_help
	echo "OS and ARCH must not be combined with ALL"
	exit 1
fi

if [ \( -n "$os" -a -z "$arch" \) -o \( -z "$os" -a -n "$arch" \) ]; then
	show_help
	echo "OS and ARCH must be specified"
	exit 1
fi

# Run unit tests
echo "Run unit tests"
go test ./... || exit $?

mkdir -p ${outdir}/bin

# build requested file
function build_cli () {
    local filepath=$1
    mkdir -p ${filepath%/*}
    go build -ldflags "-s" -o $filepath $CLI_EXECUTABLE || return $?
}

# Do a build for one platorm, usage like: build_for_platform darwin/amd64
function build_for_platform () {
    local os=${1%/*}
    local arch=${1#*/}
    local BINARY=${BRNAME}
    if [ "windows" = $os ] ; then
        BINARY=${BINARY}.exe
    fi
    GOOS="$os" GOARCH="$arch" build_cli "${outdir}/bin/$os.$arch/${BINARY}${label}" || return $?
}

# Build as instructed
if [ -z "$os" -a -z "$all" ]; then
	echo "Building $BRNAME for native OS/ARCH"
	build_cli "${outdir}/bin/${BRNAME}${label}${timestamp}" || exit $?
elif [ -z "$all" ]; then
	validos=`expr " $OSVALUES " : ".* $os "`
	if [ "$validos" -eq 0 ]; then
		show_help
		echo "Unrecognised OS: $os"
		exit 1
	fi
	validarch=`expr " $ARCHVALUES " : ".* $arch "`
	if [ "$validarch" -eq 0 ]; then
		show_help
		echo "Unrecognised ARCH: $arch"
		exit 1
	fi
	echo "Building $BRNAME for $os/$arch:"
	build_for_platform $os/$arch || exit $?
else
	echo "Building $BRNAME for default OS/ARCH:"
	for build in "${builds[@]}"; do
		echo "    $build"
		build_for_platform $build || exit $?
	done
fi

echo
echo Successfully built the following binaries:
echo
ls -lR ${outdir}/bin
echo

END_TIME=$(date +%s)
echo "Completed build.sh (brooklyn-client go build script) in $(( $END_TIME - START_TIME ))s"

exit 0
