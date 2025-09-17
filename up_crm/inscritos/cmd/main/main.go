package main

import(
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"upcrm.com/pkg/discovery/consul"
	discovery "upcrm.com/pkg/registry"
	"upcrm.com/inscritos/internal/controller/inscritos"
	httpHandler "upcrm.com/inscritos/internal/handler/http"
	"upcrm.com/inscritos/internal/repository/memory"
)

const servicename = 