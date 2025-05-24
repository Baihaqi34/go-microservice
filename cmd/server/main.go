package main

import (
    "context"
    "log"
    "net"
    "net/http"

    proto "microservice/proto/todo"            // alias untuk proto
    todo "microservice/internal/todo"    // alias untuk service

    "google.golang.org/grpc"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc/credentials/insecure"
)


func main() {
    grpcPort := ":50051"
    httpPort := ":8080"

    lis, err := net.Listen("tcp", grpcPort)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
   todoService := todo.NewTodoServiceServer() // dari internal/todo

    proto.RegisterTodoServiceServer(grpcServer, todoService)

    go func() {
        log.Println("Start gRPC server on", grpcPort)
        if err := grpcServer.Serve(lis); err != nil {
            log.Fatalf("failed to serve grpc: %v", err)
        }
    }()

    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()
    err = proto.RegisterTodoServiceHandlerFromEndpoint(ctx, mux, grpcPort, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
    if err != nil {
        log.Fatalf("failed to start HTTP gateway: %v", err)
    }

    log.Println("Start HTTP server on", httpPort)
    if err := http.ListenAndServe(httpPort, mux); err != nil {
        log.Fatalf("failed to serve http: %v", err)
    }
}
