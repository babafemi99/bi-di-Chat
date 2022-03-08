package serverChat

import (
	"chat/chatpb/chatpb"
	"chat/entity"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)

var messageHandleObject = entity.MessageHandle{}

type ChatServer struct {
	chatpb.UnimplementedMsgServiceServer
}

func (*ChatServer) ChatService(stream chatpb.MsgService_ChatServiceServer) error {
	CUQ := rand.Intn(1e6)
	errch := make(chan error)

	// recieve messages - init a go routine
	go recieveFromStream(stream, CUQ, errch)
	// send messages - init a go routine
	go sendToStream(stream, CUQ, errch)

	return <-errch
}

// recieve message function

func recieveFromStream(stream chatpb.MsgService_ChatServiceServer, CQU int, errch chan error) {
	// implement a loop
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Chat ended", err)
			break
			errch <- err
		}
		if err != nil {
			fmt.Println("something went wrong", err)
			errch <- err
		}
		messageHandleObject.MU.Lock()
		messageHandleObject.MQue = append(messageHandleObject.MQue, entity.MessageUint{
			ClientName:        msg.Name,
			MessageBody:       msg.Body,
			MessageUniqueCode: rand.Intn(1e6),
			ClientUniqueCode:  CQU,
		})

		messageHandleObject.MU.Unlock()
		log.Printf("%v", messageHandleObject.MQue[len(messageHandleObject.MQue)-1])
	}
}

// Send Message Function

func sendToStream(stream chatpb.MsgService_ChatServiceServer, CQU int, errch chan error) {
	for {
		// loop through messages in Mque
		for {
			time.Sleep(500 * time.Millisecond)
			messageHandleObject.MU.Lock()
			if len(messageHandleObject.MQue) == 0 {
				messageHandleObject.MU.Unlock()
				break
			}
			senderUniqueCode := messageHandleObject.MQue[0].ClientUniqueCode
			senderName := messageHandleObject.MQue[0].ClientName
			message := messageHandleObject.MQue[0].MessageBody
			//  CQU := messageHandleObject.MQue[0].ClientUniqueCode

			messageHandleObject.MU.Unlock()
			// send message to designated client (do not send to same client)
			if senderUniqueCode != CQU {
				err := stream.Send(&chatpb.FromServer{Name: senderName, Body: message})
				if err != nil {
					errch <- err
				}

				messageHandleObject.MU.Lock()
				if len(messageHandleObject.MQue) > 1 {
					messageHandleObject.MQue = messageHandleObject.MQue[1:]
				} else {
					messageHandleObject.MQue = []entity.MessageUint{}
				}
				messageHandleObject.MU.Unlock()
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
