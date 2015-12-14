#!/usr/bin/env bash

declare -a FAILS

function title() {
    echo "*************************************************************************"
    echo "                 ${FUNCNAME[1]}"
    echo "*************************************************************************"
}

function fail() {
    local code=$?
    local message="$*"
    FAILS[${#FAILS[*]}]="${FUNCNAME[1]}: (${code}) ${message}"
}

function report() {
    if [ 0 -ne ${#FAILS[*]} ] ; then
        echo " ${#FAILS[*]} Test failures" 1>&2
    else
        echo All tests succeeded
    fi
    local n=0
    while [ $n -lt ${#FAILS[*]} ] ; do
        echo ${FAILS[$N]} 1>&2
        n=$(( $n + 1 ))
    done
}


function isAppStatus() {
    local appname=$1
    local status=$2
    br apps | grep ${appname} | grep ${status} > /dev/null
}

function isEntityStatus() {
    local appname=$1
    local entity=$2
    local status=$3

    br app "${appname}" ent "${entity}" | grep ${status} > /dev/null
}

function waitForCommand() {
    local N=300 # 5 minutes
    while [ $N -gt 0 ] && ! "$@" ; do
        sleep 1
        N=$(($N - 1))
    done
    if [ $N -eq 0 ] ; then
        return 1
    fi
    return 0
}



function shouldHaveBr() {
    title
    type br > /dev/null 2>&1 || fail No br
}

function shouldLogin() {
    title
    br login $1 $2 $3  || fail
}



function shouldDeployTomcat() {
    title
    local filename=$1
    local appname=$2

    br deploy $1

    waitForCommand isAppStatus "${appname}" RUNNING || fail
}

function shouldRenameApp() {
    title
    local appname=$1
    local rename=$2

    br app "${appname}" rename "${rename}"

    br apps | grep "${rename}" | grep RUNNING || fail
}

function shouldGetAppConfig() {
    title
    local appname=$1
    br app "${appname}" config | grep brooklyn.wrapper_app | grep true || fail
}

function shouldGetTomcatServerEntity() {
    title
    local appname=$1

    br app "${appname}" entity | grep TomcatServer || fail
}


function shouldRenameTomcatServerEntity() {
    title
    local appname=$1
    local rename=$2

    br app "${appname}" entity "Tomcat Server" rename "${rename}"
    br app "${appname}" entity "${rename}" | grep TomcatServer || fail
}

function shouldStopEntity() {
    title
    local appname=$1
    local entityname=$2

    br app "${appname}" ent "${entityname}" stop

    waitForCommand isEntityStatus "${appname}" "${entityname}" STOPPED || fail; return
}

function shouldRestartEntity() {
    title
    local appname=$1
    local entityname=$2

    br app "${appname}" restart "${entityname}"

    waitForCommand isEntityStatus "${appname}" "${entityname}" RUNNING || fail; return
}

function shouldStartEntity() {
    title
    local appname=$1
    local entityname=$2

    br app "${appname}" start "${entityname}"

    waitForCommand isEntityStatus "${appname}" "${entityname}" RUNNING || fail; return
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
    shouldRenameTomcatServerEntity mytest myserver
    shouldStopEntity mytest myserver
    shouldStartEntity mytest myserver
    shouldRestartEntity mytest myserver


    #... TODO add more tests here

    # If there are test failures we leave the application running, in case it helps determine what failed.
    if [ 0 -eq ${#FAILS[*]} ] ; then
      echo Stopping test application
      br app mytest stop
    fi

    report

}

main $@
exit ${#FAILS[*]}