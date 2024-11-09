.PHONY: install-deps gen gen-auth-api vendor-proto

LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.1.0
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

gen-api:
	# Удаляем и создаем папку для error_v1
	rm -rf pkg/proto/error_v1
	mkdir -p pkg/proto/error_v1
	protoc --proto_path=api/v1 --proto_path=vendor.protogen \
		--go_out=pkg/proto/error_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=bin/protoc-gen-go \
		--go-grpc_out=pkg/proto/error_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		--validate_out=lang=go:pkg/proto/error_v1 --validate_opt=paths=source_relative \
		--plugin=protoc-gen-validate=bin/protoc-gen-validate \
		api/v1/error.proto

	# Удаляем и создаем папку для auth_v1
	rm -rf pkg/proto/auth_v1
	mkdir -p pkg/proto/auth_v1
	protoc --proto_path=api/v1 --proto_path=vendor.protogen \
		--go_out=pkg/proto/auth_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=bin/protoc-gen-go \
		--go-grpc_out=pkg/proto/auth_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		--validate_out=lang=go:pkg/proto/auth_v1 --validate_opt=paths=source_relative \
		--plugin=protoc-gen-validate=bin/protoc-gen-validate \
		api/v1/auth.proto

	# Удаляем и создаем папку для payment_v1
	rm -rf pkg/proto/payment_v1
	mkdir -p pkg/proto/payment_v1
	protoc --proto_path=api/v1 --proto_path=vendor.protogen \
		--go_out=pkg/proto/payment_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=bin/protoc-gen-go \
		--go-grpc_out=pkg/proto/payment_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
		--validate_out=lang=go:pkg/proto/payment_v1 --validate_opt=paths=source_relative \
		--plugin=protoc-gen-validate=bin/protoc-gen-validate \
		api/v1/payment.proto

gen:
	make install-deps
	make gen-api

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi

gen-proto:
	rm -rf pkg/proto/error_v1
	mkdir -p pkg/proto/error_v1
	protoc -I=api/v1 --go_out=pkg/proto/error_v1 --go-grpc_out=pkg/proto/error_v1 api/v1/error.proto
	rm -rf pkg/proto/auth_v1
	mkdir -p pkg/proto/auth_v1
	protoc -I=api/v1 --go_out=pkg/proto/auth_v1 --go-grpc_out=pkg/proto/auth_v1 api/v1/auth.proto
	rm -rf pkg/proto/payment_v1
	mkdir -p pkg/proto/payment_v1
	protoc -I=api/v1 --go_out=pkg/proto/payment_v1 --go-grpc_out=pkg/proto/payment_v1 api/v1/payment.proto
