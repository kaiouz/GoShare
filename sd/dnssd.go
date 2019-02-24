package sd

import (
	"context"
	"fmt"
	"github.com/grandcat/zeroconf"
)

const serviceName = "mifinder"

func StartSD(ctx context.Context, port int) {

	server, err := zeroconf.Register(serviceName, "_http._tcp", "local.", port, []string{}, nil)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()

	fmt.Printf("service %s: Name now registered and active\n", fmt.Sprintf("%s._http._tcp.local.", serviceName))

	select {
	case <-ctx.Done():
	}
	fmt.Sprint("service stop")
}
