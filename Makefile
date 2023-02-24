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

generate-admin-openapi-spec:
	mkdir -p docs
	protoc \
		--proto_path=api/proto \
		--openapiv2_out=docs \
		--openapiv2_opt=use_go_templates=true,allow_merge=true,merge_file_name=admin,json_names_for_fields=true \
		api/proto/admin/v1/*.proto


lint:
	golangci-lint run

migrate-up:
	migrate -path ./migrations -database "cockroachdb://root:@localhost:26257/english?sslmode=disable" -verbose up


migrate-down:
	migrate -path ./migrations -database "cockroachdb://root:@localhost:26258/cherrity?sslmode=disable" -verbose down

