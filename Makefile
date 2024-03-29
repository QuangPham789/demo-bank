DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name postgresdb -p 5432:5432  -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgresdb createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresdb dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go
mock:
	mockgen -destination db/mock/store.go  github.com/QuangPham789/demo-bank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
