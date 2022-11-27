#!/bin/bash

BRED='\033[1;31m'
GREEN='\033[0;32m'
NC='\033[0m'

echo -e "\n${BRED}**** chref binary building ****${NC}\n"

echo -e "${GREEN}[+] building linux amd64 ...${NC}"
env GOOS=linux GOARCH=amd64 go build -o pkg/linux_amd64/bin/chref -v -ldflags="-X 'chref/build.Version=$(cat VERSION)' -X 'chref/build.BuildUser=Team chref' -X 'chref/build.BuildTime=$(date)'" && cp -p {VERSION,LICENSE} pkg/linux_amd64/. && echo -e "${GREEN}[+] building linux arm64 ...${NC}"
env GOOS=linux GOARCH=arm64 go build -o pkg/linux_arm64/bin/chref -v -ldflags="-X 'chref/build.Version=$(cat VERSION)' -X 'chref/build.BuildUser=Team chref' -X 'chref/build.BuildTime=$(date)'" && cp -p {VERSION,LICENSE} pkg/linux_arm64/. && echo -e "${GREEN}[+] building freebsd amd64 ...${NC}"
env GOOS=freebsd GOARCH=amd64 go build -o pkg/freebsd_amd64/bin/chref -v -ldflags="-X 'chref/build.Version=$(cat VERSION)' -X 'chref/build.BuildUser=Team chref' -X 'chref/build.BuildTime=$(date)'" && cp -p {VERSION,LICENSE} pkg/freebsd_amd64/. && echo -e "${GREEN}[+] building darwin (MacOS) amd64 ...${NC}"
env GOOS=darwin GOARCH=amd64 go build -o pkg/darwin_amd64/bin/chref -v -ldflags="-X 'chref/build.Version=$(cat VERSION)' -X 'chref/build.BuildUser=Team chref' -X 'chref/build.BuildTime=$(date)'" && cp -p {VERSION,LICENSE} pkg/darwin_amd64
