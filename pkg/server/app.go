package server

import (
	"context"
	"fmt"
	"log"

	"github.com/scottyw/go-app/generated/app"
)

type appServer struct {
}

func (*appServer) Hello(context context.Context, req *app.HelloRequest) (*app.HelloResponse, error) {
	log.Println("Saying hello ...")
	return &app.HelloResponse{Message: fmt.Sprintf("Hello %s", req.Name)}, nil
}
