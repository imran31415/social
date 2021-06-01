package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "social/app/grpc/protos"
	"social/repo"
)

type Server struct {
	r repo.Iface
}

func NewServer(r repo.Iface) *Server {
	return &Server{r: r}
}

func Run(grpcPort string, serv *Server) error {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		return fmt.Errorf("failed to listen, err: %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterSocialServer(s, serv)
	log.Println("Serving GRPC...")
	if err = s.Serve(lis); err != nil {
		log.Println("Err running GRPC server: ", err)
		return err
	}
	return nil
}
