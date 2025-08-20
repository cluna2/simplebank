DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

pull_postgres:
	docker pull postgres:17-alpine

pull_redis:
	docker pull redis:latest

network: 
	docker network create bank-network

postgres:
	docker run --name postgres17 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb: 
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank


migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

db_docs:
	dbdocs build doc/db.dbml

db_schema: 
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml
	
sqlc:
	sqlc generate

test: 
	go test -v -cover -short ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/cluna2/simplebank/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/cluna2/simplebank/worker TaskDistributor


proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simplebank \
    proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans: 
	evans --host localhost --port 9090 -r repl

redis: 
	docker run --name redis -p 6379:6379 -d redis:7-alpine

clean_postgres:
	docker stop postgres17
	docker remove postgres17
	docker rmi postgres:17-alpine

clean_redis:
	docker stop redis
	docker remove redis
	docker rmi redis
	

build_local:
	make pull_postgres
	make postgres
	sleep 1
	make createdb
	make migrateup
	make sqlc
	make mock
	make proto
	make pull_redis
	make redis
	
.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans clean_postgres clean_redis build
