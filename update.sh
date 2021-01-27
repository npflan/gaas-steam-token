#!/bin/bash

./build_proto.sh

REGISTRY=$1
VERSION=$(date +%Y%m%d-%H%M%S)
echo $VERSION
# Local registry
echo "Pushing images to local registry"
docker build -t $REGISTRY/steam-token-handler .
docker tag $REGISTRY/steam-token-handler $REGISTRY/steam-token-handler:$VERSION
docker push $REGISTRY/steam-token-handler:$VERSION
docker push $REGISTRY/steam-token-handler:latest

# Docker Hub
#echo "Pushing images to Docker hub"
#docker tag $REGISTRY/steam-token-handler npflan/gaas-steam-token:$VERSION
#docker tag $REGISTRY/steam-token-handler npflan/gaas-steam-token
#docker push npflan/gaas-steam-token:$VERSION
#docker push npflan/gaas-steam-token:latest

cd manifests && kustomize edit set image $REGISTRY/steam-token-handler=$REGISTRY/steam-token-handler:$VERSION
