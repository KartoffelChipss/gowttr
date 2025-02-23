#!/bin/bash

set -e

LATEST_RELEASE=$(curl -s https://api.github.com/repos/KartoffelChipss/gowttr/releases/latest | grep "browser_download_url" | cut -d '"' -f 4)

if [ -z "$LATEST_RELEASE" ]; then
    echo "Error: Could not find the latest release."
    exit 1
fi

echo "Downloading latest release from: $LATEST_RELEASE"
curl -L "$LATEST_RELEASE" -o gowttr

chmod +x gowttr
sudo mv gowttr /usr/local/bin/

echo "Installation complete! Run 'gowttr' to get started."