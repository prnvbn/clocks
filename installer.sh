#!/bin/bash

SUPPORTED_PLATFORMS=("linux-arm64" "linux-amd64" "darwin")

# check if platform is set
if [ -z "$PLATFORM" ]; then
    echo "Error: PLATFORM not set."
    exit 1
fi

# Check if the platform is supported
supported=false
for supported_platform in "${SUPPORTED_PLATFORMS[@]}"; do
    if [ "$PLATFORM" == "$supported_platform" ]; then
        supported=true
    fi
done

if [ "$supported" == "false" ]; then
    echo "Error: Platform '$PLATFORM' not supported."
    echo "Supported platforms:" "${SUPPORTED_PLATFORMS[@]}"
    exit 1
fi

echo "Downloading the clocks binary for $PLATFORM"
echo "This may take a few seconds..."


curl -s https://api.github.com/repos/prnvbn/clocks/releases/latest \
| grep "browser_download_url" \
| cut -d : -f 2,3 \
| tr -d \" \
| grep darwin \
| xargs -I{} curl -o clocks -sL {}

chmod +x clocks
