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
	scribble "github.com/utsavanand2/scribble-diary/api/go"
	"google.golang.org/grpc"
)

type textToImageServer struct{}

func main() {
	// flag.Int() returns a pointer to an interger{here, port}
	port := flag.Int("p", 80, "port to listen on")
	flag.Parse()

	logrus.Infof("Server listening on port %d", *port)

	// net = Standard Library Package(predefined) ; Listen = Function defined in net package.
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not create a listener on port %d: %v", *port, err)
	}

	server := grpc.NewServer()

	// RegisterTextToImageServer is used to attach the 'server' to 'textToImageServer' to operate on 'textToImageServer'
	scribble.RegisterTextToImageServer(server, textToImageServer{})

	// server.Serve() expects a net.Listener type to start listening(listen to requests) on the listerner
	// here 'server' is our grpc server
	err = server.Serve(listener)
	if err != nil {
		logrus.Fatalf("could not server: %v", err)
	}
}

func (textToImageServer) Convert(ctx context.Context, imgspc *scribble.ImageSpec) (*scribble.Image, error) {

	file, err := tempfile.TempFile("", "", ".png")
	if err != nil {
		// Since no image is created return nil for image and err for error (where err is returned as a concatenated message with fmt.Errorf())
		return nil, fmt.Errorf("could not create temp file: %v", err)
	}
	defer file.Close()

	defer func() {
		err := os.Remove(file.Name())
		if err != nil {
			logrus.Warnf("could not remove file %s: %v", file.Name(), err)
		}
	}()

	// create a command to be executed later with cmd.Run() that will create the image from our given text
	cmd := exec.Command("convert", "-size", fmt.Sprintf("%dx", imgspc.GetImgsize()),
		"-background", "white", "-font", "fonts/Pacifico/Pacifico-Regular.ttf",
		"-pointsize", "70", "-gravity", "Center",
		fmt.Sprintf("caption:%s", imgspc.GetText()),
		file.Name())

	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("convert failed: %s", err)
	}

	// data is the array of bytes that holds the file contents
	data, err := ioutil.ReadFile(file.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read tmp file: %v", err)
	}

	logrus.Infof("created file %s", file.Name())
	return &scribble.Image{Image: data}, nil

}
