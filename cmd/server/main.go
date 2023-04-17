package main

import (
	"log"

	"github.com/sdivyansh59/go-grpc-service/internal/db"
	"github.com/sdivyansh59/go-grpc-service/internal/rocket"
)

//pkgm

func Run() error {
	// responsible for initializing and starting
	// out gRPC server
	rocketStore, err := db.New()
	if err != nil {
		return err
	}

	err = rocketStore.Migrate()
	if err != nil {
		log.Println("Failed to run migrations")
		return err
	}
	
	_ = rocket.New(rocketStore)

	return nil 
}
 


func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
	
}