find proto -name "*.proto" | xargs protoc \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	--go_out=generated \
	--go-grpc_out=generated \
    # --proto_path=proto
