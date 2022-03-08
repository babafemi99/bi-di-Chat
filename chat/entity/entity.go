package entity

import (
	"bufio"
	"chat/chatpb/chatpb"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type MessageUint struct {
	ClientName        string
	MessageBody       string
	MessageUniqueCode int
	ClientUniqueCode  int
}

type MessageHandle struct {
	MQue []MessageUint
	MU   sync.Mutex
}

type ClientHandle struct {
	Stream     chatpb.MsgService_ChatServiceClient
	clientName string
}

func (ch *ClientHandle) ClientConfig() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read from console: %v\n", err)
	}
	ch.clientName = strings.Trim(name, "\r\n")
}

func (ch *ClientHandle) SendMessage() {
	for {
		reader := bufio.NewReader(os.Stdin)
		//fmt.Print("\n Enter your message: ")
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from console: %v\n", err)
		}
		clientMessage = strings.Trim(clientMessage, "\r\n")
		clientMessageBox := &chatpb.FromClient{
			Name: ch.clientName,
			Body: clientMessage,
		}
		if err := ch.Stream.Send(clientMessageBox); err != nil {
			log.Printf("Error while sending message to the server: %v", err)
		}
	}
}

// receive a message

func (ch *ClientHandle) ReceiveMessage() {
	for {
		msg, err := ch.Stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving message from server: %v", err)
		}
		fmt.Printf(" NEW MWSSAGE -- \n %s : %s\n", msg.Name, msg.Body)
	}
}
