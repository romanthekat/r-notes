#!/usr/bin/env bash

basePath="github.com/romanthekat/r-notes/cmd"

for cmdPath in cmd/*/; do
  cmdName=$(basename $cmdPath)
  echo path - $cmdPath
  echo name - $cmdName
  echo ""

  CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o "$cmdName"_linux "$basePath/$cmdName"
  CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o "$cmdName"_mac "$basePath/$cmdName"
  CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o "$cmdName"_mac_arm "$basePath/$cmdName"
  CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o "$cmdName"_win "$basePath/$cmdName"
done
