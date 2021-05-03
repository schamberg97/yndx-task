#!/usr/bin/env bash
while { echo -en "HTTP/1.1 200 OK\r\nConnection: close\r\n\r\n$(date '+%s')"; } | nc -l "${1:-9000}" > /dev/null; do
  :
done
