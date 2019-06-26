package main

import (
	"context"
	"flag"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	scribble "github.com/utsavanand2/scribble-diary/api"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("b", "localhost:8080", "Scribble backend")
	output := flag.String("o", "image.png", "Image output from server")
	text := flag.String("t", "My Scribble Diary", "Text to create image of")
	flag.Parse()

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("could not connect to %s: %v", *backend, err)
	}
	defer conn.Close()

	client := scribble.NewTextToImageClient(conn)

	imgspec := &scribble.ImageSpec{Text: *text, Fontsize: 70, Imgsize: 720}

	res, err := client.Convert(context.Background(), imgspec)
	if err != nil {
		logrus.Fatalf("could not convert text %s to image: %v", imgspec.Text, err)
	}

	if err := ioutil.WriteFile(*output, res.Image, 0666); err != nil {
		logrus.Fatalf("could not write to %s: %v", *output, err)
	}

}
