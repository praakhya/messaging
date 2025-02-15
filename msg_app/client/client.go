package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/examples/data"

	pb "github.com/messaging/msg_app/msg"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:3000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func makeID(id int) *pb.ID {
	idnew := int32(id)
	return &pb.ID{Id: &idnew}
}

func showTopics(client pb.MessengerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	f := false
	fil := ""
	stream, err := client.GetTopics(ctx, &pb.Filter{Enabled: &f, Filter: &fil})
	if err != nil {
		log.Printf("Failure in getting topics: %v", err)
	}
	log.Printf("--- Topics(s) ---\n")
	for {
		topic, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Failure in getting topics: %v", err)
		}
		log.Printf("%d. %s", *topic.Id, *topic.Name)
	}

}
func readMessages(client pb.MessengerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	topicID := 0
	fmt.Printf("Enter topic id: ")
	fmt.Scanf("%d", &topicID)
	topic, err := client.GetTopic(ctx, makeID(topicID))
	if topic != nil && err == nil {
		stream, err := client.GetMsgs(ctx, topic)
		if err != nil {
			log.Printf("Failure in reading messages: %v", err)
		} else {
			log.Printf("--- Topic: %s ---\n%s\t\t%s\n", *topic.Name, "ID", "Message")
			for {
				msg, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Printf("Failure in reading messages: %v", err)
				}
				log.SetFlags(0)
				log.Printf("%d\t\t%s\n", *msg.Id, *msg.Content)
			}
		}
	} else {
		if err != nil {
			log.Printf("Topic retrieval unsuccessful: %v", err)
		}
	}
}
func addMessage(client pb.MessengerClient) {
	fmt.Println("Adding message")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	topicID := 0
	fmt.Printf("Enter topic id: ")
	fmt.Scanf("%d", &topicID)
	topic, err := client.GetTopic(ctx, makeID(topicID))
	if topic != nil && err == nil {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter message: ")
		content, _ := reader.ReadString('\n')
		//fmt.Scanf("%*[ ]%[^\n]", &content)
		id := int32(-1)
		msg, err := client.AddMsg(ctx, &pb.Msg{Id: &id, Topic: topic.Id, Content: &content})
		if err == nil {
			log.Println("Message added: ", msg)
		} else {
			log.Printf("Message could not be added: %v", err)
		}
	} else {
		log.Printf("Topic could not be fetched: %v", err)
	}

}
func addTopic(client pb.MessengerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//topicName := ""
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter topic name: ")
	topicName, _ := reader.ReadString('\n')
	//fmt.Scanf("%s", &topicName)
	id := int32(-1)
	topic, err := client.AddTopic(ctx, &pb.Topic{Id: &id, Name: &topicName})
	if err == nil {
		log.Println("Topic added: ", topic)
	} else {
		log.Printf("Topic could not be added: %v", err)
	}

}

func deleteMessage(client pb.MessengerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	messageID := 0
	fmt.Printf("Enter message id: ")
	fmt.Scanf("%d", &messageID)
	message, err := client.GetMsg(ctx, makeID(messageID))
	if message != nil && err == nil {
		msg, err := client.DelMsg(ctx, message)
		if err == nil {
			log.Println("Message deleted: ", msg)
		} else {
			log.Printf("Message could not be deleted: %v", err)
		}
	} else {
		log.Printf("Message could not be fetched: %v", err)
	}

}
func deleteTopic(client pb.MessengerClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	topicID := 0
	fmt.Printf("Enter topic id: ")
	fmt.Scanf("%d", &topicID)
	topic, err := client.GetTopic(ctx, makeID(topicID))
	if topic != nil && err == nil {
		msg, err := client.DelTopic(ctx, topic)
		if err == nil {
			log.Println("Topic deleted: ", msg)
		} else {
			log.Printf("Topic could not be deleted: %v", err)
		}
	} else {
		log.Printf("Topic could not be fetched: %v", err)
	}

}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Printf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Printf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewMessengerClient(conn)

	fmt.Println("Welcome to messenger")
	options := "1. See Topics\n2. Read Messages\n3. Add Message\n4. Add Topic\n5. Delete Message\n6. Delete Topic\n<q to quit>\n>> "
	fmt.Print(options)
	ch := "q"
	fmt.Scanf("%s", &ch)
	fmt.Println("Option: ", ch)
	for ch != "q" {
		switch ch {
		case "1":
			showTopics(client)
		case "2":
			readMessages(client)
		case "3":
			addMessage(client)
		case "4":
			addTopic(client)
		case "5":
			deleteMessage(client)
		case "6":
			deleteTopic(client)
		case "q":
			break
		default:
			fmt.Printf("The option '%s' is invalid. Try again.\n", ch)
		}
		fmt.Print(options)
		fmt.Scanf("%s", &ch)
	}
}
