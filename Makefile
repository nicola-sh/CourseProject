postgres:
	docker run --name hike-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:latest

createdb:
	docker exec -it hike-postgres createdb --username=root --owner=root dohike

dropdb:
	docker exec -it hike-postgres dropdb dohike

migrateup:
	migrate -path hike/db/migration -database "postgresql://root:root@localhost:5432/dohike?sslmode=disable" -verbose up

migratedown:
	migrate -path hike/db/migration -database "postgresql://root:root@localhost:5432/dohike?sslmode=disable" -verbose down
sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc