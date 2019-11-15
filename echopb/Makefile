PROTO_FILE="echo.proto"

all:
	# go output: echo.pb.go
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		$(PROTO_FILE)

	# gw output: echo.pb.gw.go
	protoc -I/usr/local/include -I. \
	    -I${GOPATH}/src \
	    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	    --grpc-gateway_out=logtostderr=true:. \
	    $(PROTO_FILE)

	# swagger output: echo.swagger.json
	protoc -I/usr/local/include -I. \
	    -I${GOPATH}/src \
	    -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	    --swagger_out=logtostderr=true:. \
	    $(PROTO_FILE)