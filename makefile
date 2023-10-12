# protobuf 文件生成
PROTOC_DIR := $(shell pwd)/protobuf/script/

go:
	@$(PROTOC_DIR)/protoc --plugin=protoc-gen-bic=$(PROTOC_DIR)/protoc-gen-bic --plugin=protoc-gen-go=$(PROTOC_DIR)/protoc-gen-go --go_out=. --bic_out=. protobuf/http/*.proto --proto_path=protobuf
	@$(PROTOC_DIR)gomodifytags -file pkg/gen/api/user.pb.go -all -add-tags  form -w -quiet
	@$(PROTOC_DIR)protoc-go-inject-tag -input="./pkg/gen/api/*.pb.go"
	@$(PROTOC_DIR)swag fmt
	@$(PROTOC_DIR)swag init

# swag 文档生成
swag:
	@$(PROTOC_DIR)swag init

# test 执行单元测试
test:
	@go test -v ./...

# 编译打包
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gin-bic main.go
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gin-bic main.go
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags " -X 'main.goVersion=$(go version)' -X 'main.gitTag=v1.88.88' -X 'main.gitHash=$(git show -s --format=%H)'" -o gin-bic main.go

ts:
	@$(PROTOC_DIR)/protoc --plugin=protoc-gen-bic=$(PROTOC_DIR)/protoc-gen-bic --plugin=protoc-gen-go=$(PROTOC_DIR)/protoc-gen-go --bic_out=ts_dir=../vue-admin-bic/src/api:. protobuf/http/*.proto --proto_path=protobuf

js:
	@$(PROTOC_DIR)/protoc --plugin=protoc-gen-bic=$(PROTOC_DIR)/protoc-gen-bic --plugin=protoc-gen-go=$(PROTOC_DIR)/protoc-gen-go --bic_out=js_dir=../vue-admin-bic/src/api:. protobuf/http/*.proto --proto_path=protobuf