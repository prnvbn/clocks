#!/bin/bash
set -euo pipefail

# Check for required dependencies
for cmd in curl grep cut tr xargs; do
    if ! command -v $cmd &>/dev/null; then
        echo "Error: '$cmd' is required but not installed. Please install it and try again."
        exit 1
    fi
done

# Determine platform
platform=$(uname -ms)
case $platform in
    'Darwin x86_64') target_platform=darwin-amd64 ;;
    'Darwin arm64') target_platform=darwin-arm64 ;;
    'Linux aarch64' | 'Linux arm64') target_platform=linux-arm64 ;;
    'Linux x86_64') target_platform=linux-amd64 ;;
    *)
        echo "Unsupported platform: ${platform}"
        echo "Please open an issue: https://github.com/prnvbn/clocks/issues/new"
        exit 1
    ;;
esac

echo "Detected platform: $platform -> Target binary: $target_platform"
echo "Downloading the latest Clocks binary..."

# Fetch the latest release URL
download_url=$(curl -s https://api.github.com/repos/prnvbn/clocks/releases/latest | 
               grep "browser_download_url" | 
               cut -d '"' -f 4 | 
               grep "$target_platform" || true)

if [[ -z "$download_url" ]]; then
    echo "Error: Could not find a compatible binary for $target_platform."
    exit 1
fi

# Download and make executable
curl -sL "$download_url" -o clocks
chmod +x clocks

echo "-------------------------------------------------------------------"
echo "✅ Clocks has been downloaded to: $(pwd)/clocks"
echo ""

# Ask where to move the binary (loop until valid input)
while true; do
    echo "Where would you like to install Clocks?"
    echo "1) /usr/local/bin (system-wide, requires sudo)"
    echo "2) ~/.local/bin (user only, no sudo needed)"
    echo "3) Keep it in the current directory"
    read -p "Enter choice (1/2/3): " choice

    case "$choice" in
        1)
            sudo mv clocks /usr/local/bin/
            echo "✅ Clocks has been installed globally! You can now run 'clocks' from anywhere."
            break
        ;;
        2)
            mkdir -p "$HOME/.local/bin"
            mv clocks "$HOME/.local/bin/"
            echo "✅ Clocks has been installed to ~/.local/bin."
            
            # Check if ~/.local/bin is in PATH
            if [[ ! "$PATH" =~ (^|:)"$HOME/.local/bin"(:|$) ]]; then
                echo "ℹ️  ~/.local/bin is not in your PATH."
                echo "   Add this to your shell profile (e.g., ~/.bashrc or ~/.zshrc):"
                echo '   export PATH="$HOME/.local/bin:$PATH"'
            fi
            break
        ;;
        3)
            echo "ℹ️  Clocks will remain in the current directory."
            echo "   You can move it manually later if needed."
            break
        ;;
        *)
            echo "❌ Invalid choice. Please enter 1, 2, or 3."
        ;;
    esac
done

echo ""
echo "ℹ️  To enable auto-completion, visit:"
echo "   https://github.com/prnvbn/clocks/tree/main?tab=readme-ov-file#enabling-command-autocompletion"
echo "-------------------------------------------------------------------"
