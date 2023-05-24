network:
	docker network create bank-network
postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1qaz2wsx -d postgres:alpine3.17
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
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
sqlc:
	sqlc generate
test:
	go test -v -cover -short ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/kys20548/simple_bank/db/sqlc Store
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto
evans:
	evans --host localhost --port 9090 -r repl

.PHONY:postgres createDB dropDB migratedown migrateup migratedown1 migrateup1 sqlc test server mock network new_migration proto evans