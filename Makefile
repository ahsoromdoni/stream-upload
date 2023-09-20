protogen:
	@echo "Loading ..."
	@go get google.golang.org/grpc@v1.40.0
	@go install github.com/golang/protobuf/protoc-gen-go@v1.5.2
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v1.16.0
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v1.16.0
	@go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators@v0.3.2

	@protoc --go_out=plugins=grpc:grpc/. grpc/upload.proto
	@echo "Done"