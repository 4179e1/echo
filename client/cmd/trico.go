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
	"io"
	"log"
	"strings"

	pb "github.com/4179e1/echo/echopb"
	"github.com/spf13/cobra"
)

// tricoCmd represents the trico command
var tricoCmd = &cobra.Command{
	Use:   "trico",
	Short: "trico command",
	Long:  "trico command",
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

		stream, err := client.Trico(context.Background(), data)
		if err != nil {
			log.Fatalf(err.Error())
		}

		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			fmt.Println(msg)
		}

	},
}

func init() {
	rootCmd.AddCommand(tricoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tricoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tricoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
