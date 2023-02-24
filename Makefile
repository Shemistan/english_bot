generate-proto:
	mkdir -p pkg
	protoc \
		--proto_path=api/proto \
		--go_out=pkg \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pkg \
		--grpc-gateway_opt=logtostderr=true,paths=source_relative,generate_unbound_methods=true \
		--validate_out lang=go:pkg \
		api/proto/admin/v1/*.proto
	mv pkg/github.com/Shemistan/english_bot/api/api/proto/admin/v1/* pkg/admin/v1/
	rm -rf pkg/github.com

lint:
	golangci-lint run

migrate-up:
	migrate -path ./migrations -database "cockroachdb://root:@localhost:26257/english?sslmode=disable" -verbose up


migrate-down:
	migrate -path ./migrations -database "cockroachdb://root:@localhost:26258/cherrity?sslmode=disable" -verbose down

