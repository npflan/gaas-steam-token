#!/bin/bash

./build_proto.sh

VERSION=$(date +%Y%m%d-%H%M%S)
echo $VERSION
# Local registry
echo "Pushing images to local registry"
docker build -t registry.npf.dk/steam-token-handler .
docker tag registry.npf.dk/steam-token-handler registry.npf.dk/steam-token-handler:$VERSION
docker push registry.npf.dk/steam-token-handler:$VERSION
docker push registry.npf.dk/steam-token-handler:latest

# Docker Hub
echo "Pushing images to Docker hub"
docker tag registry.npf.dk/steam-token-handler npflan/gaas-steam-token:$VERSION
docker tag registry.npf.dk/steam-token-handler npflan/gaas-steam-token
docker push npflan/gaas-steam-token:$VERSION
docker push npflan/gaas-steam-token:latest

cd manifests && kustomize edit set image registry.npf.dk/steam-token-handler=registry.npf.dk/steam-token-handler:$VERSION
