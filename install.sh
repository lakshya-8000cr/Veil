#!/bin/bash
set -e

BOLD="\033[1m"
DIM="\033[2m"
GREEN="\033[32m"
CYAN="\033[36m"
WHITE="\033[97m"
RESET="\033[0m"

URL="https://github.com/lakshya-8000cr/veil/releases/download/v1.0.0/veil"

echo ""
echo -e "  ${BOLD}${WHITE}veil${RESET}  ${DIM}›  installing${RESET}"
echo ""

echo -e "  ${DIM}·  fetching binary${RESET}"
sudo curl -fsSL "$URL" -o /usr/local/bin/veil 2>/dev/null
echo -e "  ${GREEN}✔${RESET}  binary downloaded"

echo -e "  ${DIM}·  setting permissions${RESET}"
sudo chmod +x /usr/local/bin/veil
echo -e "  ${GREEN}✔${RESET}  permissions set"

echo -e "  ${GREEN}✔${RESET}  veil ready ${DIM}→${RESET} ${WHITE}/usr/local/bin/veil${RESET}"

echo ""
echo -e "  ${DIM}spawn    ${RESET}${CYAN}veil spawn <name>${RESET}"
echo -e "  ${DIM}help     ${RESET}${CYAN}veil --help${RESET}"
echo ""