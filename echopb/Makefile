PROTO_FILE="echo.proto"
GRPC_GATEWAY_PATH=$(shell go list -m -f "{{.Dir}}" github.com/grpc-ecosystem/grpc-gateway)

all:
	# go output: echo.pb.go
	protoc -I/usr/local/include -I. \
		-I$(GRPC_GATEWAY_PATH)/third_party/googleapis \
		--go_out=plugins=grpc:. \
		$(PROTO_FILE)

	# gw output: echo.pb.gw.go
	protoc -I/usr/local/include -I. \
		-I$(GRPC_GATEWAY_PATH)/third_party/googleapis \
	    --grpc-gateway_out=logtostderr=true:. \
	    $(PROTO_FILE)

	# swagger output: echo.swagger.json
	protoc -I/usr/local/include -I. \
		-I$(GRPC_GATEWAY_PATH)/third_party/googleapis \
	    --swagger_out=logtostderr=true:. \
	    $(PROTO_FILE)

	# generate swagger.pb.go from swagger.go using scripts/includetxt.go
	go generate .



deps:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go

clean:
	rm -rf *.go *.json