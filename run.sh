#!/bin/sh
# run linters and tests
go fmt ./... && go vet ./... && go test ./...
# start mysql and profile service
docker compose up -d
# wait for mysql to warm up
sleep 30
# up migrations
docker run -v ${PWD}/services/profile/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database  'mysql://root:example@tcp(localhost:3306)/mysql' up 
# run tests
docker run -it --rm --network host -v ${PWD}:/etc/newman -t postman/newman:alpine run otus-social-network.postman_collection.json --env-var "host=localhost"





    