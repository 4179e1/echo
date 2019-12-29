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
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	pb "github.com/4179e1/echo/echopb"
	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn := getClientConn()
		defer conn.Close()

		client := pb.NewEchoServiceClient(conn)
		stream, err := client.Chat(context.Background())
		if err != nil {
			panic(err)
		}

		waitc := make(chan struct{})
		go func() {
			for {
				msg, err := stream.Recv()
				if err == io.EOF {
					// read done.
					close(waitc)
					return
				}
				if err != nil {
					panic(err)
				}

				fmt.Printf(fmt.Sprintf("\r< %s\n> ", msg.Msg))
			}
		}()

		// Read Stdin
		reader := bufio.NewReader(os.Stdin)
		for i := int32(1); ; i++ {
			fmt.Printf("> ")
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}

			data := &pb.EchoRequest{
				Index: i,
				Msg:   strings.TrimSuffix(line, "\n"),
			}

			if err := stream.Send(data); err != nil {
				fmt.Println(err)
				panic(err)
			}

		}

		stream.CloseSend()
		<-waitc

	},
}

func init() {
	rootCmd.AddCommand(chatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
