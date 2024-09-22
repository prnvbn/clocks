#!/bin/bash
set -eou pipefail

platform=$(uname -ms)

target_platform=""
case $platform in
    'Darwin x86_64')
        target_platform=darwin-amd64
    ;;
    'Darwin arm64')
        target_platform=darwin-arm64
    ;;
    'Linux aarch64' | 'Linux arm64')
        target_platform=linux-arm64
    ;;
    'Linux x86_64')
        target_platform=linux-amd64
    ;;
    *)
        echo "The installer script doesn't support ${platform}"
        echo "please open an issue - https://github.com/prnvbn/clocks/issues/new"
        exit 1
    ;;
esac

echo "Downloading the clocks binary for $target_platform"
echo "This may take a few seconds..."


curl -s https://api.github.com/repos/prnvbn/clocks/releases/latest \
| grep "browser_download_url" \
| cut -d : -f 2,3 \
| tr -d \" \
| grep "$target_platform" \
| xargs -I{} curl -o clocks -sL {}

chmod +x clocks

sep="-------------------------------------------------------------------"

echo $sep
echo "The clocks binary has been installed in in the current directory - $(pwd)"
echo "You can add the directory to your PATH"

echo "To enable command auto completion for clocks referr to - https://github.com/prnvbn/clocks/tree/main?tab=readme-ov-file#enabling-command-autocompletion"

echo $sep
