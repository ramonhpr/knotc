#!/usr/bin/env bash

package_name=knotc

targetOSes=("linux" "windows" "darwin")
platforms=("amd64")
version=`cat VERSION`

mkdir -p build
for targetOS in "${targetOSes[@]}"
do
    for platform in "${platforms[@]}"
    do
        output_name=$package_name'-'$targetOS'-'$platform'-'v$version
        if [ $targetOS = "windows" ]; then
            output_name+='.exe'
        fi

        env GOOS=${targetOS} GOARCH=${platform} go build -ldflags="-X main.version=$version" -o build/$output_name cmd/main.go
        if [ $? -ne 0 ]; then
            echo "An error has occurred on $targetOS/$platform! Aborting the script execution..."
            exit 1
        fi
    done
done
