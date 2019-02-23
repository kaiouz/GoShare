package sd

import (
	"context"
	"fmt"
	"github.com/brutella/dnssd"
)

const serviceName = "mifinder"

func StartSD(ctx context.Context, port int) {

	if resp, err := dnssd.NewResponder(); err != nil {
		fmt.Println(err)
	} else {
		cfg := dnssd.Config{
			Name: serviceName,
			Type: "_http._tcp",
			Port: port,
		}

		srv, err := dnssd.NewService(cfg)

		if err != nil {
			fmt.Println(err)
			return
		}

		handle, err := resp.Add(srv)

		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("service %s: Name now registered and active\n", handle.Service().ServiceInstanceName())
		}

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		err = resp.Respond(ctx)

		if err != nil {
			fmt.Println(err)
		}
	}
}
