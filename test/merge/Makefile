regenerate:
	go install github.com/temporalio/gogo-protobuf/protoc-min-version
	go install github.com/temporalio/gogo-protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=. \
	--proto_path=../../../../../:../../protobuf/:. merge.proto
