#! /bin/bash
set -e

mkdir -p /usr/local/go && curl -Ls https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz | tar xvzf - -C /usr/local/go --strip-components=1


