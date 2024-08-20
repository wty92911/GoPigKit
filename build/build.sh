#!/usr/bin/env bash

# shellcheck disable=SC2164
cd "$(cd "$(dirname "$0")";pwd)"
cd ../

env=$1
unit_test=$2
lint=$3
bin=pigkit

set -e
export GO111MODULE=on

version=`cat CHANGELOG.md|grep "^Version "| head -1|awk -F' ' '{print $2}'`
fourthVersion=`date "+%Y%m%d%H%M"`
version=${version}_${fourthVersion}

if [[ `echo ${env} |grep -c win` -eq 1 ]]; then
    export GOARCH=amd64 GOOS=windows
fi

if [[ `echo ${env} |grep -c linux` -eq 1 ]]; then
    export GOARCH=amd64 GOOS=linux
fi

if [[ `echo ${env} |grep -c mac` -eq 1 ]]; then
    export GOARCH=amd64 GOOS=darwin
fi

if [[ "$lint" == lint ]]; then
    golangci-lint run
fi

function go_test() {
    mkdir -p coverage
    go test -coverprofile=coverage/cover.out ./...
    if [[ $? -ne 0 ]]; then
        echo "unit test failed!!!"
        exit 1
    fi
    go tool cover -html=coverage/cover.out -o coverage/coverage.html
}

echo ["Unit Test"]

if [[ "$unit_test" == test ]];
then
    go_test
else
    echo -e "\033[31mUnit tests are not performed!\033[0m"
fi

echo
echo [version]
go version
echo version ${version}

echo
echo [branch]
echo -n ${bin}": "
# shellcheck disable=SC2006
branch=`git branch |  awk  '$1 == "*"{print $2}'`
echo -e "\033[32m${branch}\033[0m"

echo
echo [build]
go build -ldflags "-s -w -X main.version=${version}" -o ./bin/${bin} ./cmd/pigkit

rm -rf ./dist ./*.tar.gz
mkdir -p ./dist/${bin}/log ./dist/${bin}/conf ./dist/${bin}/static

cp -r ./bin ./dist/${bin}/
cp -r ./cmd/${bin}/server.sh ./dist/${bin}/bin/

cd dist

tarFileName=${bin}_${version}.tar.gz
tar czf "${tarFileName}" ${bin}/*
mv ./*.tar.gz ../
cd ..

# shellcheck disable=SC2006
if [[ `uname` = Darwin ]]; then
    md5 "${tarFileName}"
elif [[ `uname` = Linux ]]; then
    md5sum "${tarFileName}"
else
    md5sum "${tarFileName}"
fi

echo
echo -e "\033[32m编译成功\033[0m"
echo
