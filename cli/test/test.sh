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

declare -a FAILS

export BRCLI_HOME=/tmp
trap cleanup EXIT


function usage() {
    echo "Usage: $0 <brooklyn_url> [ <user> <password> ]"
    exit 0
}


function cleanup() {
    rm -f $BRCLI_HOME/.brooklyn_cli
}

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
    br apps | grep "${appname}" | grep ${status} > /dev/null
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



function brIsDefined() {
    type br > /dev/null 2>&1
}


#
# TESTS
#

function shouldLogin() {
    title
    br login $1 $2 $3  || fail
}



function shouldDeployTomcat() {
    title

    br deploy test_app.yaml

    waitForCommand isAppStatus "Test Tomcat" RUNNING || fail
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

    br app "${appname}" entity "Tomcat 7 Server" rename "${rename}"
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


function runAllTests() {

    local brooklyn_url=${1}
    local user=${2}
    local password=${3}

    shouldLogin ${brooklyn_url} ${user} ${password}
    shouldDeployTomcat
    shouldRenameApp "Test Tomcat" mytest
    shouldGetAppConfig mytest
    shouldGetTomcatServerEntity mytest
    shouldRenameTomcatServerEntity mytest myserver
    shouldStopEntity mytest myserver
    shouldStartEntity mytest myserver
    shouldRestartEntity mytest myserver


    #... TODO add more tests here
}

#
# main function
#
function main() {

    [ $1 ] || usage

    local brooklyn_url=${1}
    local user=${2}
    local password=${3}

    brIsDefined || {
        >&2 echo br is not defined
        exit 1
    }

    runAllTests ${brooklyn_url} ${user} ${password}

    # If there are test failures we leave the application running, in case it helps determine what failed.
    if [ 0 -eq ${#FAILS[*]} ] ; then
      echo Stopping test application
      br app mytest stop
    fi

    report

}

main $@
exit ${#FAILS[*]}
