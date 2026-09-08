postgres:
	docker run --name postgres18 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:18-alpine
postgres-terminal:
	docker exec -it postgres18 psql -U root apex_bank
migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/apex_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/apex_bank?sslmode=disable" -verbose down
createdb:
	docker exec -it postgres18 createdb --username=root --owner=root apex_bank
sqlc:
	sqlc generate
dropdb:
	docker exec -it postgres18 dropdb apex_bank
.PHONY: createdb dropdb postgres postgres-terminal migrateup migratedown