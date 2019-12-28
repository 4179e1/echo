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
	"strings"

	pb "github.com/4179e1/echo/echopb"
	"github.com/spf13/cobra"
)

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo <msg>",
	Short: "echo command",
	Long:  "echo command",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		conn := getClientConn()
		defer conn.Close()
		client := pb.NewEchoServiceClient(conn)

		separator := " "
		msg := strings.Join(args, separator)

		data := &pb.EchoRequest{
			Index: 1,
			Msg:   msg,
		}

		sugar.Debug("Sending Requset ", data)

		reply, err := client.Echo(context.Background(), data)
		if err != nil {
			// TODO https://jiajunhuang.com/articles/2019_09_02-go_grpc_handshake.md.html
			log.Fatalf(err.Error())
		}

		fmt.Println(reply.Msg)
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
