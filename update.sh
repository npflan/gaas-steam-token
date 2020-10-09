#!/bin/bash
VERSION=$(date +%Y%m%d-%H%M%S)
docker build -t registry.npf.dk/csgo-token .
docker tag registry.npf.dk/csgo-token registry.npf.dk/csgo-token:$VERSION
docker push registry.npf.dk/csgo-token:$VERSION
docker push registry.npf.dk/csgo-token
