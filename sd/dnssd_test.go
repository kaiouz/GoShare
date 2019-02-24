package sd

import (
	"context"
	"testing"
)

func TestStartSD(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	StartSD(ctx, 8080)
}
