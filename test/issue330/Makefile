regenerate:
	go install github.com/temporalio/gogo-protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/temporalio/gogo-protobuf/protoc-gen-gogo/descriptor:. --proto_path=../../../../../:../../protobuf/:. *.proto
