package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"runtime"
	"strings"

	"github.com/digitaldata-cz/htmltopdf"
	pb "github.com/digitaldata-cz/pdfgen/proto/go"

	"google.golang.org/grpc"
)

var run = make(chan func())

type tGrpcServer struct {
	// TODO: Nastudovat k cemu je "UnimplementedPrinterServer"
	pb.UnimplementedPdfGenServer
}

func init() {
	// Set main function to run on the main thread.
	runtime.LockOSThread()

	// Initialize library.
	if err := htmltopdf.Init(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer htmltopdf.Destroy()

	go startServer()

	// Listen for functions that need to run on the main thread.
	var quit = make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	for {
		select {
		case f := <-run:
			f()
		case <-quit:
			log.Println("shutting down")
			return
		}
	}
}

// callFunc calls the provided function on the main thread.
func callFunc(f func() error) error {
	err := make(chan error)
	run <- func() {
		err <- f()
	}
	return <-err
}

func startServer() {
	var (
		ipAddress = "0.0.0.0"
		port      = "50051"
	)

	if s := os.Getenv("PS_IP"); s != "" {
		ipAddress = s
	}

	if s := os.Getenv("PS_PORT"); s != "" {
		port = s
	}

	listener, err := net.Listen("tcp", net.JoinHostPort(ipAddress, port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterPdfGenServer(s, &tGrpcServer{})
	log.Printf("server listening at %s", listener.Addr().String())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}

func (s *tGrpcServer) Generate(ctx context.Context, in *pb.GenerateRequest) (*pb.GenerateResponse, error) {
	out := bytes.NewBuffer(nil)
	if err := callFunc(func() error {
		tmpl, err := htmltopdf.NewObjectFromReader(strings.NewReader(in.GetHtmlBody()))
		if err != nil {
			return err
		}
		converter, err := htmltopdf.NewConverter()
		if err != nil {
			log.Fatal(err)
		}
		defer converter.Destroy()
		converter.Add(tmpl)

		colorMode := "Color"
		if in.GetGrayscale() {
			colorMode = "Grayscale"
		}
		tmpl.Zoom = in.GetZoom()
		converter.DPI = in.GetDpi()
		converter.PaperSize = htmltopdf.PaperSize(in.GetPageSize())
		converter.Orientation = htmltopdf.Orientation(in.GetOrientation())
		converter.Colorspace = htmltopdf.Colorspace(colorMode)
		converter.MarginLeft = in.GetMarginLeft()
		converter.MarginRight = in.GetMarginRight()
		converter.MarginTop = in.GetMarginTop()
		converter.MarginBottom = in.GetMarginBottom()
		converter.UseCompression = true
		return converter.Run(out)
	}); err != nil {
		return &pb.GenerateResponse{Report: nil, Error: err.Error()}, nil
	}
	fmt.Println("gRPC JEDEEEEEE")
	return &pb.GenerateResponse{Report: out.Bytes()}, nil
}
