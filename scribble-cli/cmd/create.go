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
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	scribble "github.com/utsavanand2/scribble-diary/api"
	"google.golang.org/grpc"
)

var text, output *string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a captioned image file from a caption.",
	Long: `create generates a captioned image file by taking in a text caption string
and generating a captioned image file from it.`,
	Run: func(cmd *cobra.Command, args []string) {
		text, _ := cmd.Flags().GetString("caption")
		output, _ := cmd.Flags().GetString("output")
		server, _ := cmd.Flags().GetString("server")
		create(context.Background(), text, output, server)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringP("caption", "c", "My Scribble Diary", "input text to be used as caption for the image")
	createCmd.Flags().StringP("output", "o", "image.png", "path to the output file")
	createCmd.Flags().StringP("server", "s", "35.232.69.118:80", "IPv4 + port address of the scribble-server")

}

func create(ctx context.Context, text, output, server string) {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect to %s: %v", server, err)
	}
	defer conn.Close()

	client := scribble.NewTextToImageClient(conn)

	imgspec := &scribble.ImageSpec{Text: text, Fontsize: 70, Imgsize: 720}

	res, err := client.Convert(ctx, imgspec)
	if err != nil {
		logrus.Fatalf("could not convert text %s to image: %v", imgspec.Text, err)
	}

	if err := ioutil.WriteFile(output, res.Image, 0666); err != nil {
		logrus.Fatalf("could not write to %s: %v", output, err)
	}
}
