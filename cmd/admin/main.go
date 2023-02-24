package main

import (
	"context"
	"fmt"
	"github.com/Shemistan/english_bot/internal/api"
	"github.com/Shemistan/english_bot/internal/service"
	"github.com/Shemistan/english_bot/internal/storage"
	v1 "github.com/Shemistan/english_bot/pkg/admin/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"sync"
)

const (
	portGrpc = ":9001"
	portHttp = ":8001"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		log.Fatal(startGRPC())
	}()

	go func() {
		defer wg.Done()
		log.Fatal(startHTTP())
	}()

	wg.Wait()

}

func startGRPC() error {

	sqlConnectionString := fmt.Sprintf("user=root password=root host=localhost port=26257 dbname=english sslmode=disable pool_max_conns=10")

	pgxPoolConfig, err := pgxpool.ParseConfig(sqlConnectionString)
	if err != nil {
		return errors.Wrap(err, "failed to parse pgx config")
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), pgxPoolConfig)
	if err != nil {
		return errors.Wrap(err, "failed to create pgxpool")
	}

	newStorage := storage.New(dbPool)
	newService := service.New(newStorage)

	listener, err := net.Listen("tcp", portGrpc)
	if err != nil {
		return errors.Wrap(err, "failed to listen tcp")
	}

	s := grpc.NewServer()

	v1.RegisterBotServiceServer(s, &api.Bot{BotService: newService})

	return s.Serve(listener)
}

func startHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := v1.RegisterBotServiceHandlerFromEndpoint(ctx, mux, portGrpc, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(portHttp, mux)
}
