package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"upcrm.com/interesados/internal/controller/interesados"
	httpHandler "upcrm.com/interesados/internal/handler/http"
	"upcrm.com/interesados/internal/repository/memory"
	"upcrm.com/pkg/discovery/consul"
	discovery "upcrm.com/pkg/registry"
)

const servicename = "interesados"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("Starting rating service on port %d", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateINstanceID(servicename)
	if err := registry.Register(ctx, instanceID, servicename, fmt.Sprintf("localhost:%d", port)); err != nil {
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
	ctrl := interesados.New(repo)
	h := httpHandler.New(ctrl)
	http.Handle("/interesados", http.HandlerFunc(h.Handler))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
