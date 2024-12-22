createdb:
	docker exec -it postgres-amole createdb --username=root --owner=root amole_db
dropdb:
	docker exec -it postgres-amole dropdb amole_db
migrate-cart-table:
	migrate create -ext sql -dir ./script/migrations -seq tb_amole_cart
migrate-cart-item-table:
	migrate create -ext sql -dir ./script/migrations -seq tb_amole_cart_items
migrateup:
	migrate -path script/migrations -database "postgresql://root:secret@localhost:5432/amole_db?sslmode=disable" -verbose up
migratedown:
	migrate -path script/migrations -database "postgresql://root:secret@localhost:5432/amole_db?sslmode=disable" -verbose down
proto:
	protoc \
      --go_out=paths=import:./api/grpc/protos/ \
      --go-grpc_out=paths=import:./api/grpc/protos/ \
      --proto_path=./api/grpc/protos/ \
      ./api/grpc/protos/*.proto
.PHONY: createdb dropdb migrate migrateup migratedown proto