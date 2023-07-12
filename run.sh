#!/bin/sh
# run linters and tests
go fmt ./... && go vet ./... && go test ./...
# start env
docker-compose up -d --build 
# wait for mysql to warm up
sleep 20
# up migrations
docker run --rm -v ${PWD}/services/profile/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database  'postgres://postgres:example@localhost:5432/postgres?sslmode=disable' up 
# wait for migrations to finish
sleep 10
# run tests
docker run -it --rm --network host -v ${PWD}:/etc/newman -t postman/newman:alpine run otus-social-network.postman_collection.json --env-var "host=localhost"
# wait for tests to finish
sleep 10
# stop env
# docker-compose down -v