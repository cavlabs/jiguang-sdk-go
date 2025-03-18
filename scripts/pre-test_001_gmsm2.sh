#!/bin/bash

sm2B64PubKey=$(awk -F'"' '/sm2B64PubKey/ {print $2}' jiguang/gmsm2_test.go)
sm2B64PrivKey=$(awk -F'"' '/sm2B64PrivKey/ {print $2}' jiguang/gmsm2_test.go)

if [ "$(uname)" = "Darwin" ]; then
  sed -i '' "s|sm2B64PubKey  = \"[^\"]*\"|sm2B64PubKey  = \"$sm2B64PubKey\"|" "jiguang/gmsm2.go"
  sed -i '' "s|sm2B64PrivKey = \"[^\"]*\"|sm2B64PrivKey = \"$sm2B64PrivKey\"|" "jiguang/gmsm2.go"
else # Linux
  sed -i "s|sm2B64PubKey  = \"[^\"]*\"|sm2B64PubKey  = \"$sm2B64PubKey\"|" "jiguang/gmsm2.go"
  sed -i "s|sm2B64PrivKey = \"[^\"]*\"|sm2B64PrivKey = \"$sm2B64PrivKey\"|" "jiguang/gmsm2.go"
fi
