package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
	"github.com/tink-ab/tempfile"
	scribble "github.com/utsavanand2/scribble-diary/api"
	"google.golang.org/grpc"
)

type textToImageServer struct{}

func main() {
	port := flag.Int("p", 80, "port to listen on")
	flag.Parse()

	logrus.Infof("Server listening on port %d", *port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not create a listener on port %d: %v", *port, err)
	}

	server := grpc.NewServer()
	scribble.RegisterTextToImageServer(server, textToImageServer{})
	err = server.Serve(listener)
	if err != nil {
		logrus.Fatalf("could not server: %v", err)
	}
}

func (textToImageServer) Convert(ctx context.Context, imgspc *scribble.ImageSpec) (*scribble.Image, error) {

	file, err := tempfile.TempFile("", "", ".png")
	if err != nil {
		return nil, fmt.Errorf("could not create temp file: %v", err)
	}
	defer file.Close()

	defer func() {
		if err := os.Remove(file.Name()); err != nil {
			logrus.Warnf("could not remove file %s: %v", file.Name(), err)
		}
	}()

	cmd := exec.Command("convert", "-size", fmt.Sprintf("%dx%d", imgspc.GetImgsize(), imgspc.GetImgsize()),
		"-background", "white", "-font", "fonts/Pacifico/Pacifico-Regular.ttf",
		"-pointsize", "70", "-gravity", "Center",
		fmt.Sprintf("caption:%s", imgspc.GetText()),
		file.Name())

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("convert failed: %s", err)
	}

	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read tmp file: %v", err)
	}

	logrus.Infof("created file %s", file.Name())
	return &scribble.Image{Image: data}, nil

}
