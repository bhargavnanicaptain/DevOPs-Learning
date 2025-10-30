#!/bin/bash
echo "Enter a site to check the connectivity issues:"
read site

echo -e "\nPinging site..."
ping -c 3 "$site"             # use -c to limit pings

echo -e "\nChecking whether DNS is resolving or not..."
nslookup "$site"

echo -e "\nIdentifying slow or failing network..."
if command -v traceroute >/dev/null 2>&1; then
  tracepath "$site"
else
  echo "tracepath command not found (common on Windows Git Bash)"
fi

echo -e "\nFetching headers using curl..."
curl -I "$site"
