package registry

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Registry interface {
	Register(ctx context.Context, instanceID string, servicename string, hostPort string) error
	Deregister(ctx context.Context, instanceId string, servicename string) error
	ServiceAddress(ctx context.Context, serviceID string) ([]string, error)
	ReportHealthyState(instanceID string, servicename string) error
}

var ErrNotFound = errors.New("No service addresses found")

func GenerateINstanceID(servicename string) string {
	return fmt.Sprintf("%s-%d", servicename, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
