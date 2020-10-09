package handler

import (
	"fmt"
	pb "github.com/Cytram/csgo-token/proto"
	steam "github.com/Cytram/csgo-token/pkg/steam"
	"context"

)

// Server Struct
type Server struct {
	pb.UnimplementedTokenserviceServer
}

func (s *Server) CreateToken (ctx context.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenReply, error) {
	game, err := steam.CreateAccount(req.AppId, req.Memo)
	if err != nil {
		fmt.Println(err)
	}
	return &pb.CreateTokenReply{Id: game.LoginToken}, nil
}
