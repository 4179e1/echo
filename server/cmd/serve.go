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
	"fmt"

	pb "github.com/4179e1/echo/echopb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
		fmt.Println("PidFile:", viper.GetString("Server.PidFile"))
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
	serveCmd.Flags().IntP("server.port", "P", 8081, "bind port")

	viper.BindPFlag("Server.Host", serveCmd.Flags().Lookup("server.host"))
	viper.BindPFlag("Server.Port", serveCmd.Flags().Lookup("server.port"))
}

// implemented pb.EchoServiceServer interface
// its methods were copied from echo.pb.go with modification
type echoService struct {
}

func (*echoService) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
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

func serve() {
	// TODO: certs
	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEchoServiceServer(grpcServer, newEchoServer())

}
