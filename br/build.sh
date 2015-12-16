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
BRFILE="brooklyn.go"
BRDIR="brooklyn-cli/br"

#
# Compile options
#
export CGO_ENABLED=0

#
# Globals
#
os=""
arch=""
all=""

show_help() {
	echo "Usage:	$0"
	echo "	$0 (-o | --os) <OS> (-a | --arch) <ARCH>"
	echo "	$0 --all"
	echo "	$0 (-h | --help)"
	echo $OSVALUES | awk 'BEGIN{printf("OS:\n")};{for(i=1;i<=NF;i++){printf("\t%s\n",$i)}}'
	echo $ARCHVALUES | awk 'BEGIN{printf("ARCH:\n")};{for(i=1;i<=NF;i++){printf("\t%s\n",$i)}}'
	echo
}

while [ $# -gt 0 ]; do
	case $1 in 
	-h|--help|help)
		show_help
		exit 0
		;;
	-o|--os)
		if [ $# -lt 2 ]; then
			show_help
			echo "Value for OS must be provided"
			exit 1
		fi
		os="$2"
		shift 2
		;;
	-a|--arch)
		if [ $# -lt 2 ]; then
			show_help
			echo "Value for ARCH must be provided"
			exit 1
		fi
		arch="$2"
		shift 2
		;;
	--all)
		all="all"
		shift 1
		;;
	*)
		show_help
		echo "Unrecognised parameter: $1"
		exit 1
		;;
	esac
done

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

thisdir=`pwd`
validdir=`expr "$thisdir" : ".*${BRDIR}\$"`
if [ "$validdir" -eq 0 ]; then
	echo "Must be in CLI directory: $BRDIR"
	exit 2
fi
if [ ! -f "$BRFILE" ]; then
	echo "Directory must contain CLI file: $BRFILE"
	exit 2
fi

if [ -z "$os" -a -z "$all" ]; then
	echo "Building $NAME for native OS/ARCH"
	go build
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
	export GOOS="$os"
	export GOARCH="$arch"
	echo "Building $NAME for $os/$arch"
	go build -o "$BRNAME.$os.$arch"
else
	echo "Building $NAME for common OS/ARCH:"
	os="$OSVALUES"
	arch="$ARCHVALUES"
	for j in $arch; do
		printf "  "
		for i in $os; do
			export GOOS="$i"
			export GOARCH="$j"
			printf "$i/$j "
			go build -o "$BRNAME.$i.$j"
		done
		printf "\n"
	done
fi

exit 0
