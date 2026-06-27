#!/bin/bash

set -e

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo " VEIL Installer"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo ""

URL="https://github.com/lakshya-8000cr/Veil/releases/download/v1.0.0/Veil"

echo "Downloading Veil..."
sudo curl -fsSL "$URL" -o /usr/local/bin/veil

echo "Setting permissions..."
sudo chmod +x /usr/local/bin/veil

echo ""
echo "Veil installed successfully."
echo ""
echo "Run:"
echo "  veil --help"
echo "  veil spawn"
echo ""