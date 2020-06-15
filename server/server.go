package server

import (
	"context"
	"net"

	"github.com/pkg/errors"
	pb "github.com/webmalc/namecheap-dns/protos/changer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Server runs the server
type Server struct {
	logger  logger
	config  *config
	changer Changer
}

// Change changes the DNS records
func (s *Server) Change(
	ctx context.Context, in *pb.ChangeRequest,
) (*pb.ChangeReply, error) {
	s.changer.Change(in.GetIp())
	return &pb.ChangeReply{Result: true}, nil
}

// Run runs the server
func (s *Server) Run() {
	lis, err := net.Listen("tcp", s.config.address)
	if err != nil {
		s.logger.Error(errors.Wrap(err, "server"))
	}
	creds, err := credentials.NewServerTLSFromFile(
		s.config.crtPath, s.config.keyPath,
	)
	if err != nil {
		s.logger.Error(errors.Wrap(err, "server"))
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterChangerServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		s.logger.Error(errors.Wrap(err, "server"))
	}
}

// NewServer creates a new server instance
func NewServer(l logger, c Changer) *Server {
	return &Server{logger: l, config: newConfig(), changer: c}
}
