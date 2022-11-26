#!/bin/bash

echo "chref binary building"

echo "[+] building linux amd64 ..."
env GOOS=linux GOARCH=amd64 go build -o pkg/linux_amd64/bin/chref -v -ldflags="-X 'chref/build.Version=$(cat VERSION)' -X 'chref/build.BuildUser=Team chref' -X 'chref/build.BuildTime=$(date)'" && cp -p {VERSION,LICENSE} pkg/linux_amd64/. && echo "[+] building linux arm64 ..."
env GOOS=linux GOARCH=arm64 go build -o pkg/linux_arm64/bin/chref -v -ldflags="-X 'chref/build.Version=$(cat VERSION)' -X 'chref/build.BuildUser=Team chref' -X 'chref/build.BuildTime=$(date)'" && cp -p {VERSION,LICENSE} pkg/linux_arm64/. && echo "[+] building freebsd amd64 ..."
env GOOS=freebsd GOARCH=amd64 go build -o pkg/freebsd_amd64/bin/chref -v -ldflags="-X 'chref/build.Version=$(cat VERSION)' -X 'chref/build.BuildUser=Team chref' -X 'chref/build.BuildTime=$(date)'" && cp -p {VERSION,LICENSE} pkg/freebsd_amd64/. && echo "[+] building darwin (MacOS) amd64 ..."
env GOOS=darwin GOARCH=amd64 go build -o pkg/darwin_amd64/bin/chref -v -ldflags="-X 'chref/build.Version=$(cat VERSION)' -X 'chref/build.BuildUser=Team chref' -X 'chref/build.BuildTime=$(date)'" && cp -p {VERSION,LICENSE} pkg/darwin_amd64
