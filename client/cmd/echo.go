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
	"log"

	pb "github.com/4179e1/echo/echopb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("echo called")
		opts := []grpc.DialOption{
			grpc.WithInsecure(),
		}
		//creds := credentials.NewClientTLSFromCert(demoCertPool, "localhost:10000")
		//opts = append(opts, grpc.WithTransportCredentials(creds))
		hostPort := fmt.Sprintf("%s:%s", viper.GetString("Server.Host"), viper.GetString("Server.Port"))
		sugar.Debug("Dialing %s...", hostPort)
		conn, err := grpc.Dial(hostPort, opts...)
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()
		client := pb.NewEchoServiceClient(conn)

		reply, err := client.Echo(context.Background(), &pb.EchoRequest{Index: 0, Msg: "Hello"})
		if err != nil {
			// TODO https://jiajunhuang.com/articles/2019_09_02-go_grpc_handshake.md.html
			log.Fatalf(err.Error())
		}

		fmt.Println(reply)
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// echoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// echoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
