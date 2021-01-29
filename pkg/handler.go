package handler

import (
	"context"
	"github.com/Cytram/csgo-token/pkg/steam"
	pb "github.com/Cytram/csgo-token/proto"
	"log"
)

// Server Struct
type Server struct {
	pb.UnimplementedTokenserviceServer
}

func (s *Server) CreateToken (ctx context.Context, req *pb.CreateTokenRequest) (*pb.CreateTokenReply, error) {
	log.Printf("creating token for %s for game %d", req.Memo, req.AppId)
	game, err := steam.CreateAccount(req.AppId, req.Memo)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Token %s created for %s", game.LoginToken, req.Memo)
	return &pb.CreateTokenReply{ServerToken: game.LoginToken}, nil
}

func (s *Server) DeleteToken (ctx context.Context, req *pb.DeleteTokenRequest) (*pb.DeleteTokenReply, error){
	log.Printf("deleting token %s", req.ServerToken)
	err := steam.DeleteAccount(req.ServerToken)
	if err != nil{
		log.Println(err)
	}
	// return 1 for because no empty reply
	return &pb.DeleteTokenReply{Id: "1"}  , nil
}

