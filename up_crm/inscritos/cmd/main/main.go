package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"upcrm.com/inscritos/internal/controller/inscritos"
	httpHandler "upcrm.com/inscritos/internal/handler/http"
	"upcrm.com/inscritos/internal/repository/memory"
	"upcrm.com/pkg/discovery/consul"
	discovery "upcrm.com/pkg/registry"
)

const servicename = "inscritos"

func main() {
	var port int
	flag.IntVar(&port, "port", 8083, "API handler port")
	flag.Parse()
	log.Printf("Starting rating service on port %d", port)
	registry, err := consul.NewRegistry(os.Getenv("CONSUL_HTTP_ADDR"))
	if err != nil {
		log.Fatalf("Error creating Consul registry: %v", err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateINstanceID(servicename)
	if err := registry.Register(ctx, instanceID, servicename, fmt.Sprintf("inscritos:%d", port)); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, servicename); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()

	defer registry.Deregister(ctx, instanceID, servicename)
	repo := memory.New()
	ctrl := inscritos.New(repo)
	h := httpHandler.New(ctrl)
	http.Handle("/inscritos", http.HandlerFunc(h.Handler))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
