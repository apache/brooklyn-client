#!/bin/sh
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


#
# Constants
#
OSVALUES="darwin freebsd linux netbsd openbsd windows"
ARCHVALUES="386 amd64"
BRNAME="br"
GOPACKAGE="github.com/apache/brooklyn-client/${BRNAME}"
EXECUTABLE_DIR="$GOPATH/src/$GOPACKAGE"
GOBIN=go
GODEP=godep

#
# Test if go and godep are available
#
command -v $GOBIN >/dev/null 2>&1 || { echo "Command for compiling Go not found: $GOBIN" 1>&2 ; exit 1; }

if [ -z "$GOPATH" -o ! -d "$GOPATH" ]; then
	echo "Environment variable GOPATH must be set to a valid directory"
	exit 1
fi

if [ ! -x "$GOPATH/bin/$GODEP" ]; then
	echo "Command for resolving dependencies ($GODEP) not found in GOPATH: $GOPATH"
	exit 1
fi

#
# Compile options
#

# Disable use of C code modules (causes problems with cross-compiling)
export CGO_ENABLED=0

#
# Globals
#
os=""
arch=""
all=""
dir="."
label=""
timestamp=""

show_help() {
# 
# -A  Build for all OS/ARCH combinations
# -a  Set ARCH to build for
# -d  Set output directory
# -h  Show help
# -l  Set label text for including in filename
# -o  Set OS to build for
# -t  Set timestamp for including in filename
#
	echo "Usage:	$0 [-d <DIRECTORY>] [-l <LABEL>] [-t]"
	echo "	$0 -o <OS> -a <ARCH> [-d <DIRECTORY>] [-l <LABEL>] [-t]"
	echo "	$0 -A [-d <DIRECTORY>] [-l <LABEL>] [-t]"
	echo "	$0 -h"
	echo $OSVALUES | awk 'BEGIN{printf("OS:\n")};{for(i=1;i<=NF;i++){printf("\t%s\n",$i)}}'
	echo $ARCHVALUES | awk 'BEGIN{printf("ARCH:\n")};{for(i=1;i<=NF;i++){printf("\t%s\n",$i)}}'
	echo
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
			echo "Value for DIRECTORY must be provided"
			exit 1
		fi
		dir="$2"
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

if [ -n "$dir" -a ! -d "$dir" ]; then
	show_help
	echo "No such directory: $dir"
	exit 1
fi

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


if [ -d ${EXECUTABLE_DIR} ]; then
    cd ${EXECUTABLE_DIR}
else
	echo "Directory not found: ${EXECUTABLE_DIR}"
	exit 2
fi

if [ -z "$os" -a -z "$all" ]; then
	echo "Building $BRNAME for native OS/ARCH"
	$GODEP $GOBIN build -ldflags "-s" -o "${dir}/${BRNAME}${label}${timestamp}" $GOPACKAGE
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
	echo "Building $BRNAME for $os/$arch"
	GOOS="$os" GOARCH="$arch" $GODEP $GOBIN build -ldflags "-s" -o "${dir}/${BRNAME}${label}${timestamp}.$os.$arch" $GOPACKAGE
else
	echo "Building $BRNAME for all OS/ARCH:"
	os="$OSVALUES"
	arch="$ARCHVALUES"
	for j in $arch; do
		for i in $os; do
			echo "    $i/$j"
			GOOS="$i" GOARCH="$j" $GODEP $GOBIN build -ldflags "-s" -o "${dir}/${BRNAME}${label}${timestamp}.$i.$j" $GOPACKAGE
		done
	done
fi

exit 0
