#!/bin/sh
docker compose up -d
sleep 30
#migrate -source file://path/to/migrations -database postgres://localhost:5432/database up
docker run -it --rm --network host -v ${PWD}:/etc/newman -t postman/newman:alpine run otus-social-network.postman_collection.json --env-var "host=localhost"
