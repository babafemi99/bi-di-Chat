package main

import (
	"chat/chatpb/chatpb"
	"chat/entity"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v\n", err)
	}
	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatalf("Error closing connection: %v\n", err)
		}
	}(cc)

	c := chatpb.NewMsgServiceClient(cc)
	doBidiStreaming(c)

}

func doBidiStreaming(c chatpb.MsgServiceClient) {
	println("Starting Bidi Streaming")
	stream, err := c.ChatService(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v\n", err)
	}

	// implement communication with grpc server
	ch := entity.ClientHandle{Stream: stream}
	ch.ClientConfig()
	go ch.SendMessage()
	go ch.ReceiveMessage()

	//blocker

	bl := make(chan bool)
	<-bl
}
