package main

import (
	serverChat "chat/chat_server"
	"chat/chatpb/chatpb"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	// wuzuppb.RegisterMsgServiceServer(s)
	chatpb.RegisterMsgServiceServer(s, &serverChat.ChatServer{})
	reflection.Register(s)

	go func() {
		fmt.Println("Starting Server ...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to server: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	// Block till end
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Stopping the listener")
	lis.Close()
	fmt.Println("closing postgres connection")
	fmt.Println("End of program")

}
