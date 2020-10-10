#!/bin/bash

./build_proto.sh

VERSION=$(date +%Y%m%d-%H%M%S)
# Local registry
docker build -t registry.npf.dk/steam-token-handler .
docker tag registry.npf.dk/steam-token-handler registry.npf.dk/steam-token-handler:$version
docker push registry.npf.dk/steam-token-handler:$version
docker push registry.npf.dk/steam-token-handler

# Docker Hub
docker tag registry.npf.dk/steam-token-handler npflan/gaas-steam-token:$version
docker tag registry.npf.dk/steam-token-handler npflan/gaas-steam-token
docker push npflan/gaas-steam-token:$version
docker push npflan/gaas-steam-token

cd manifests && kustomize edit set image registry.npf.dk/steam-token-handler=registry.npf.dk/steam-token-handler:$VERSION
