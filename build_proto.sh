#!/bin/sh

# Get latest version of steam token proto
curl https://raw.githubusercontent.com/npflan/proto/main/steam-token/token.proto -o proto/token.proto -s

protoc -I . proto/token.proto  --go_out=plugins=grpc:. --go_opt=paths=source_relative
