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

# For Brooklyn Client, we use a debian distribution instead of alpine as there are some libgcc incompatibilities with GO
FROM --platform=linux/amd64 maven:3.9.0-amazoncorretto-8

# Install necessary binaries to build brooklyn-client
RUN yum -y update && yum install -y git-core

# Download Go 1.15 and verify checksum against value from https://golang.org/dl/
# then install to /usr/local
RUN cd /tmp \
&& curl -O https://dl.google.com/go/go1.22.5.linux-amd64.tar.gz \
&& CKSUM=$(sha256sum go1.22.5.linux-amd64.tar.gz | awk '{print $1}') \
&& [ ${CKSUM} = "904b924d435eaea086515bc63235b192ea441bd8c9b198c507e85009e6e4c7f0" ] \
&& tar xf go1.22.5.linux-amd64.tar.gz \
&& rm go1.22.5.linux-amd64.tar.gz \
&& chown -R root:root ./go \
&& mv go /usr/local

ENV PATH="${PATH}:/usr/local/go/bin"

RUN mkdir -p /var/maven/.m2/ && chmod -R 777 /var/maven/
ENV MAVEN_CONFIG=/var/maven/.m2
