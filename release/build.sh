#!/bin/sh

#
# TODO
#
# Checks on availability of go coompiler
# Use godep for resolving dependencies
# Add some useful comments
#

#
# Constants
#
OSVALUES="darwin freebsd linux netbsd openbsd windows"
ARCHVALUES="386 amd64"
BRNAME="br"
PROJECT="github.com/apache/brooklyn-client"
CLI_PACKAGE="${PROJECT}/${BRNAME}"
GOBIN=go
GODEP=godep

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
outdir="."
sourcedir="."
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
# -s  Source directory
	echo "Usage:	$0 [-d <OUTPUTDIR>] [-l <LABEL>] [-t] -s <SOURCEDIR>"
	echo "	$0 -o <OS> -a <ARCH> [-d <DIRECTORY>] [-l <LABEL>] [-t] -s <SOURCEDIR>"
	echo "	$0 -A [-d <OUTPUTDIR>] [-l <LABEL>] [-t] -s <SOURCEDIR>"
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


#
# Test if go is available
#
command -v $GOBIN >/dev/null 2>&1 || { echo "Command for compiling Go not found: $GOBIN" 1>&2 ; exit 1; }


if [ -n "$outdir" -a ! -d "$outdir" ]; then
	show_help
	echo "No such directory: $outdir"
	exit 1
fi

# Set GOPATH to $outdir and link to source code.
export GOPATH=${outdir}
mkdir -p ${GOPATH}/src/${PROJECT%/*}
[ -e ${GOPATH}/src/${PROJECT} ] || ln -s ${sourcedir} ${GOPATH}/src/${PROJECT}
PATH=${GOPATH}/bin:${PATH}

command -v $GODEP >/dev/null 2>&1 || {
	go get github.com/tools/godep || { echo failed installing dodep ; exit 1; }
}

command -v $GODEP >/dev/null 2>&1 || {
	echo "Command for resolving dependencies ($GODEP) not found and could not be installed in $GOPATH"
	exit 1
}

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

EXECUTABLE_DIR="$GOPATH/src/$CLI_PACKAGE"
if [ -d ${EXECUTABLE_DIR} ]; then
    cd ${EXECUTABLE_DIR}
else
	echo "Directory not found: ${EXECUTABLE_DIR}"
	exit 2
fi

mkdir -p ${GOPATH}/bin

if [ -z "$os" -a -z "$all" ]; then
	echo "Building $BRNAME for native OS/ARCH"
	$GODEP $GOBIN build -ldflags "-s" -o "${GOPATH}/bin/${BRNAME}${label}${timestamp}" $CLI_PACKAGE
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
	mkdir -p ${GOPATH}/bin/$os.$arch
	GOOS="$os" GOARCH="$arch" $GODEP $GOBIN build -ldflags "-s" -o "${GOPATH}/bin/$os.$arch/${BRNAME}${label}" $CLI_PACKAGE
else
	echo "Building $BRNAME for all OS/ARCH:"
	os="$OSVALUES"
	arch="$ARCHVALUES"
	for archv in $arch; do
		for osv in $os; do
			echo "    $osv/$archv"
			mkdir -p ${GOPATH}/bin/$osv.$archv
			GOOS="$osv" GOARCH="$archv" $GODEP $GOBIN build -ldflags "-s" -o "${GOPATH}/bin/$osv.$archv/${BRNAME}${label}" $CLI_PACKAGE
		done
	done
fi

exit 0
