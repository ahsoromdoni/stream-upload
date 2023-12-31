package main

import (
	"context"
	"flag"
	"log"

	upload "github.com/ahsoromdoni/stream-upload/service"

	"google.golang.org/grpc"
)

func main() {
	// Catch user input.
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatalln("Missing file path")
	}

	// Initialise gRPC connection.
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// Start uploading the file. Error if failed, otherwise echo download URL.
	client := upload.NewClient(conn)
	name, err := client.Upload(context.Background(), flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(name)
}
