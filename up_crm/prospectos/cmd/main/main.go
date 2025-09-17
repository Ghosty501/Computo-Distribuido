package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"upcrm.com/pkg/discovery/consul"
	discovery "upcrm.com/pkg/registry"
	"upcrm.com/prospectos/internal/controller/prospectos"
	httphandler "upcrm.com/prospectos/internal/handler/http"
	"upcrm.com/prospectos/internal/repository/memory"
)

const servicename = "prospectos"

func main() {
	var port int
	flag.IntVar(&port, "port", 8081, "API handler port")
	flag.Parse()
	log.Printf("Starting prospectos service in port %d", port)
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
	r := memory.New()
	c := prospectos.New(r)
	h := httphandler.New(c)

	http.Handle("/Prospecto", http.HandlerFunc(h.GetProspecto))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}

}
