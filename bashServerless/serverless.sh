#!/bin/bash
RESPONSE=$(echo "{\"statusCode\":200, \"body\":$(date '+%s')}" | jq '.')
echo $RESPONSE | jq -c '.body |= tostring'