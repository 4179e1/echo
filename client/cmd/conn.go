package cmd

import (
	"fmt"

	"github.com/4179e1/echo/common"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

func getClientConn() *grpc.ClientConn {

	hostPort := fmt.Sprintf("%s:%d", viper.GetString("Server.Host"), viper.GetInt("Server.Port"))
	_, certPool := common.GetCerts(viper.GetString("Server.CertFile"), viper.GetString("Server.KeyFile"))
	creds := credentials.NewClientTLSFromCert(certPool, hostPort)
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	//creds := credentials.NewClientTLSFromCert(demoCertPool, "localhost:10000")
	//opts = append(opts, grpc.WithTransportCredentials(creds))
	sugar.Debug(fmt.Sprintf("Dialing %s...", hostPort))
	conn, err := grpc.Dial(hostPort, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	return conn
}
