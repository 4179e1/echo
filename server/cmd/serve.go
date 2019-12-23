/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/4179e1/echo/common"
	pb "github.com/4179e1/echo/echopb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Echo Server",
	Long:  `Start the Echo Server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		fmt.Println("args: ", args)
		fmt.Println("PidFile:", viper.GetString("Global.PidFile"))
		fmt.Printf("Listen %s:%s\n", viper.GetString("Server.Host"), viper.GetString("Server.Port"))
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	serveCmd.Flags().StringP("server.host", "H", "0.0.0.0", "bind address")
	serveCmd.Flags().IntP("server.port", "P", 8080, "bind port")
	serveCmd.Flags().String("server.certfile", "/etc/echo/server.pem", "cert file")
	serveCmd.Flags().String("server.keyfile", "/etc/echo/server.key", "key file")

	viper.BindPFlag("Server.Host", serveCmd.Flags().Lookup("server.host"))
	viper.BindPFlag("Server.Port", serveCmd.Flags().Lookup("server.port"))
	viper.BindPFlag("Server.CertFile", serveCmd.Flags().Lookup("server.certfile"))
	viper.BindPFlag("Server.KeyFile", serveCmd.Flags().Lookup("server.keyfile"))
}

// implemented pb.EchoServiceServer interface
// its methods were copied from echo.pb.go with modification
type echoService struct {
}

// curl -k -X POST https://localhost:8080/echo/api/v1/echo
func (*echoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoReply, error) {
	//return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
	reply := pb.EchoReply{
		Index: 1,
		Msg:   "hello",
	}

	return &reply, nil
}
func (*echoService) Trico(req *pb.EchoRequest, srv pb.EchoService_TricoServer) error {
	return status.Errorf(codes.Unimplemented, "method Trico not implemented")
}
func (*echoService) Sink(srv pb.EchoService_SinkServer) error {
	return status.Errorf(codes.Unimplemented, "method Sink not implemented")
}

func (*echoService) Chat(srv pb.EchoService_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}

func newEchoServer() *echoService {
	return new(echoService)
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func serve() {
	// TODO: certs
	host := viper.GetString("Server.Host")
	port := viper.GetInt("Server.Port")
	serverAddr := fmt.Sprintf("%s:%d", host, port)
	keyPair, certPool := common.GetCerts(viper.GetString("Server.CertFile"), viper.GetString("Server.KeyFile"))

	serverNameOverride := fmt.Sprintf("localhost:%d", port)
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(certPool, serverNameOverride)),
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEchoServiceServer(grpcServer, newEchoServer())

	// TODO: dial opts with certs
	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: serverAddr,
		RootCAs:    certPool,
	})
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}
	mux := http.NewServeMux()

	ctx := context.Background()
	gwmux := runtime.NewServeMux()
	err := pb.RegisterEchoServiceHandlerFromEndpoint(ctx, gwmux, serverAddr, dopts)
	if err != nil {
		panic(err)
	}

	mux.Handle("/", gwmux)

	//hostPort := fmt.Sprintf("%s:%s", viper.GetString("Server.Host"), viper.GetString("Server.Port"))
	conn, err := net.Listen("tcp", serverAddr)
	if err != nil {
		panic(err)
	}

	/*
		if err := grpcServer.Serve(conn); err != nil {
			panic(err)
		}
	*/

	srv := &http.Server{
		Addr:    serverAddr,
		Handler: grpcHandlerFunc(grpcServer, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*keyPair},
			NextProtos:   []string{"h2"},
		},
	}

	err = srv.Serve(tls.NewListener(conn, srv.TLSConfig))
	if err != nil {
		panic(err)
	}

	return
}
