#!/bin/bash

./build_proto.sh

VERSION=$(date +%Y%m%d-%H%M%S)
docker build -t registry.npf.dk/steam-token-handler .
docker tag registry.npf.dk/steam-token-handler registry.npf.dk/steam-token-handler:$VERSION
docker push registry.npf.dk/steam-token-handler:$VERSION
docker push registry.npf.dk/steam-token-handler

cd manifests && kustomize edit set image registry.npf.dk/steam-token-handler=registry.npf.dk/steam-token-handler:$VERSION
