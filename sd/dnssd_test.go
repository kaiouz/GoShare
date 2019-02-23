package sd

import "testing"

func TestStartSD(t *testing.T) {
	if err := StartSD(8080); err != nil {
		t.Error(err)
	}
}
