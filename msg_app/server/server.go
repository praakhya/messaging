package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"regexp"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"

	pb "github.com/messaging/msg_app/msg"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 3000, "The server port")
)

type messengerServer struct {
	pb.UnimplementedMessengerServer
	savedTopics   []*pb.Topic // read-only after initialized
	savedMessages []*pb.Msg
	currTopicId   int32
	currMsgId     int32
	mu            sync.Mutex // protects routeNotes

}

func removeTopic(s []*pb.Topic, i int32, server *messengerServer) []*pb.Topic {
	server.mu.Lock()
	s[i] = s[len(s)-1]
	*s[i].Id = i
	server.currTopicId--
	server.mu.Unlock()
	return s[:server.currTopicId]
}

func removeMsg(s []*pb.Msg, i int32, server *messengerServer) []*pb.Msg {
	server.mu.Lock()
	s[i] = s[len(s)-1]
	*s[i].Id = i
	server.currMsgId--
	server.mu.Unlock()
	return s[:server.currMsgId]
}

func (s *messengerServer) GetTopic(ctx context.Context, id *pb.ID) (*pb.Topic, error) {
	for _, topic := range s.savedTopics {
		if *topic.Id == *id.Id {
			return topic, nil
		}
	}
	return nil, nil
}

func (s *messengerServer) AddTopic(ctx context.Context, topic *pb.Topic) (*pb.Topic, error) {
	s.mu.Lock()
	*topic.Id = s.currTopicId
	fmt.Println("topic id incremented: ", s.currTopicId)
	s.currTopicId += 1
	s.savedTopics = append(s.savedTopics, topic)
	s.mu.Unlock()
	return topic, nil
}

func (s *messengerServer) AddMsg(ctx context.Context, message *pb.Msg) (*pb.Msg, error) {
	s.mu.Lock()
	*message.Id = s.currMsgId
	s.currMsgId += 1
	fmt.Println("msg id incremented: ", s.currMsgId)
	s.savedMessages = append(s.savedMessages, message)
	s.mu.Unlock()
	return message, nil
}

func (s *messengerServer) GetTopics(filter *pb.Filter, stream pb.Messenger_GetTopicsServer) error {
	for _, topic := range s.savedTopics {
		if *filter.Enabled && *filter.Filter != "" {
			if match, _ := regexp.MatchString(*filter.Filter, *topic.Name); match == true {
				if err := stream.Send(topic); err != nil {
					return err
				}
			}
		} else {
			if err := stream.Send(topic); err != nil {
				return err
			}
		}

	}
	return nil
}

func (s *messengerServer) GetMsg(ctx context.Context, id *pb.ID) (*pb.Msg, error) {
	for _, message := range s.savedMessages {
		if *message.Id == *id.Id {
			return message, nil
		}
	}
	return nil, nil
}

func (s *messengerServer) GetMsgs(topic *pb.Topic, stream pb.Messenger_GetMsgsServer) error {
	for _, msg := range s.savedMessages {
		if *msg.Topic == *topic.Id {
			if err := stream.Send(msg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *messengerServer) DelTopic(ctx context.Context, topic *pb.Topic) (*pb.Topic, error) {
	s.mu.Lock()
	for i, savedTopic := range s.savedTopics {
		if *topic.Id == *savedTopic.Id {
			deletedTopic := savedTopic
			s.savedTopics = removeTopic(s.savedTopics, int32(i), s)
			return deletedTopic, nil
		}
	}
	s.mu.Unlock()
	return nil, nil
}
func (s *messengerServer) DelMsg(ctx context.Context, msg *pb.Msg) (*pb.Msg, error) {
	s.mu.Lock()
	for i, savedMessage := range s.savedMessages {
		if *msg.Id == *savedMessage.Id {
			deletedMsg := savedMessage
			s.savedMessages = removeMsg(s.savedMessages, int32(i), s)
			return deletedMsg, nil
		}
	}
	s.mu.Unlock()
	return nil, nil
}

func newServer() *messengerServer {
	s := &messengerServer{savedTopics: make([]*pb.Topic, 0), savedMessages: make([]*pb.Msg, 0), currTopicId: 0, currMsgId: 0}
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = data.Path("x509/server_cert.pem")
		}
		if *keyFile == "" {
			*keyFile = data.Path("x509/server_key.pem")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials: %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterMessengerServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
