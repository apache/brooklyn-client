#!/usr/bin/env bash

declare -a FAILS

function fail() {
    local code=$?
    local message="$*"
    FAILS[${#FAILS[*]}]="${FUNCNAME[1]}: (${code}) ${message}"
}

function report() {
    if [ 0 -ne ${#FAILS[*]} ] ; then
        echo "Test failures" 1>&2
    else
        echo All tests succeeded
    fi
    local n=0
    while [ $n -lt ${#FAILS[*]} ] ; do
        echo ${FAILS[$N]} 1>&2
        n=$(( $n + 1 ))
    done
}

function shouldHaveBr() {
    type br > /dev/null 2>&1 || fail No br
}

function shouldLogin() {
    br login $1 $2 $3  || fail
}

function shouldDeployTomcat() {
    local filename=$1
    local appname=$2

    br deploy $1

    local N=300 # 5 minutes
    while [ $N -gt 0 ] && ! br apps | grep ${appname} | grep RUNNING > /dev/null ; do
        sleep 1
        N=$(($N - 1))
    done
    if [ $N -eq 0 ] ; then
        fail Timeout waiting for start
        return
    fi
    br apps | grep ${appname} | grep RUNNING || fail Not running
}

function shouldRenameApp() {
    local appname=$1
    local rename=$2

    br app "${appname}" rename "${rename}"

    br apps | grep "${rename}" | grep RUNNING || fail
}

function shouldGetAppConfig() {
    local appname=$1
    br app "${appname}" config | grep brooklyn.wrapper_app | grep true || fail
}

function shouldGetTomcatServerEntity() {
    local appname=$1

    br app "${appname}" entity | grep TomcatServer || fail
}


function shouldRenameServerEntity() {
    local appname=$1
    local rename=$2

    br app "${appname}" entity "Tomcat Server" rename "${rename}"
    br app "${appname}" entity "${rename}" | grep TomcatServer || fail
}


function main() {

    local application_yaml=test_app.yaml
    local appname=$(grep name: ${application_yaml} | sed 's/.*: *//')

    while getopts "a:" opt; do
      case $opt in
        a)
          application_yaml=$OPTARG
          shift; shift;
          ;;
        \?)
          echo "Invalid option: -$OPTARG" >&2
          ;;
      esac
    done

    local brooklyn_url=${1:-localhost}
    local user=${2}
    local password=${3}

    shouldHaveBr
    shouldLogin ${brooklyn_url} ${user} ${password}
    shouldDeployTomcat ${application_yaml} ${appname}
    shouldRenameApp "${appname}" mytest
    shouldGetAppConfig mytest
    shouldGetTomcatServerEntity mytest
    shouldRenameServerEntity mytest myserver

    report

}

main $@
exit ${#FAILS[*]}