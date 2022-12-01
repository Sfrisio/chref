#!/bin/bash

BRED='\033[1;31m'
GREEN='\033[0;32m'
NC='\033[0m'

PKG_VERSION=$(cat VERSION)
PKG_NAME="chref-${PKG_VERSION}-"

allSupportedOS=("linux" "freebsd" "darwin")

echo -e "\n${BRED}**** chref binary building ****${NC}\n"

for os in ${allSupportedOS[@]}; do 
    env GOOS=linux GOARCH=amd64 go build -o pkg/${PKG_NAME}${os}_amd64/bin/chref -v -ldflags="-X 'github.com/Sfrisio/chref/build.Version=$(cat VERSION)' -X 'github.com/Sfrisio/chref/build.BuildUser=Team chref' -X 'github.com/Sfrisio/chref/build.BuildTime=$(date)'" && cp -p {VERSION,LICENSE} pkg/${PKG_NAME}${os}_amd64/.
    if [[ ${os} != "freebsd" ]]; then
        echo -e "${GREEN}[+] building ${os} arm64 ...${NC}"
        env GOOS=linux GOARCH=arm64 go build -o pkg/${PKG_NAME}${os}_arm64/bin/chref -v -ldflags="-X 'github.com/Sfrisio/chref/build.Version=$(cat VERSION)' -X 'github.com/Sfrisio/chref/build.BuildUser=Team chref' -X 'github.com/Sfrisio/chref/build.BuildTime=$(date)'" && cp -p {VERSION,LICENSE} pkg/${PKG_NAME}${os}_arm64/.
    fi
done

cd pkg/

for d in */ ; do
    basedirname=$(basename "${d}")
    echo -e "\n[+] making tar.gz from: ${GREEN} ${basedirname} ...${NC}" && tar -zcf ${basedirname}.tar.gz ${d}
done
