package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/scottyw/go-app/generated/app"

	"github.com/scottyw/go-app/pkg/swagger"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	swagger, _ := swagger.GeneratedAppSwaggerJsonBytes()
	w.Write(swagger)
}

// StartGRPC starts the GRPC server running
func StartGRPC() {
	lis, err := net.Listen("tcp", "localhost:5566")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	app.RegisterAppServer(grpcServer, &appServer{})
	log.Println("gRPC server ready...")
	grpcServer.Serve(lis)
}

// StartHTTP starts an HTTP server to server swagger
func StartHTTP() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	conn, err := grpc.Dial("localhost:5566", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	// Register grpc-gateway
	rmux := runtime.NewServeMux()
	client := app.NewAppClient(conn)
	err = app.RegisterAppHandlerClient(ctx, rmux, client)
	if err != nil {
		log.Fatal(err)
	}

	// Serve the swagger, swagger-ui and grpc-gateway REST bindings on 8080
	mux := http.NewServeMux()
	mux.HandleFunc("/swagger.json", serveSwagger)
	mux.Handle("/", rmux)
	fs := http.FileServer(http.Dir("swagger-ui"))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", fs))
	log.Println("Swagger served at: http://localhost:8080/swagger.json")
	log.Println("Swagger UI served at: http://localhost:8080/swagger-ui")
	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
