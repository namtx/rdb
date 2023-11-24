package server

import (
	"context"
	"sync"

	pb "github.com/namtx/rdb/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	mu   sync.Mutex
	data map[string][]byte
	pb.UnimplementedRdbServer
}

func NewServer() *Server {
	s := Server{}
	s.data = make(map[string][]byte)

	return &s
}

func (s *Server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	validateKey(in.Key)

	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.data[in.Key]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Could not find the key %q", in.Key)
	}

	return &pb.GetResponse{Key: in.Key, Value: value}, nil
}

func (s *Server) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetResponse, error) {
	validateKey(in.Key)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[in.Key] = in.Value

	return &pb.SetResponse{}, nil
}

func validateKey(k string) error {
	if k == "" {
		return status.Errorf(codes.InvalidArgument, "Key cannot be empty!")
	}

	return nil
}
