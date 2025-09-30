#!/usr/bin/env bash
set -e
echo "This demo script will start docker-compose and run a placeholder HTLC demo."
docker-compose -f docker-compose.yml up -d
echo "Started demo containers. Follow logs for progress."
