postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest
createdb:
	winpty docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	winpty docker exec -it postgres dropdb simple_bank

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc