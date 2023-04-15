postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1qaz2wsx -d postgres:alpine3.17
createDB:
	docker exec -it postgres createdb --username=root --owner=root simple_bank
dropDB:
	docker exec -it postgres dropDB simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:1qaz2wsx@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:1qaz2wsx@localhost:5432/simple_bank?sslmode=disable" -verbose down
migrateup1:
	migrate -path db/migration -database "postgresql://root:1qaz2wsx@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown1:
	migrate -path db/migration -database "postgresql://root:1qaz2wsx@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover db/sqlc/*
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/kys20548/simpleBank/db/sqlc Store

.PHONY:postgres createDB dropDB migratedown migrateup migratedown1 migrateup1 sqlc test server mock