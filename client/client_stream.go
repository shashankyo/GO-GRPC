package main

import (
	"context"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("client streaming started")
	stream, err := client.callSayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names:%v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		log.Printf("sent the request with name:%s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("client streaming finished ")
	if err != nil {
		log.Fatalf("error while recing %v", err)
	}

	log.Printf("%v", res.Messages)
}
