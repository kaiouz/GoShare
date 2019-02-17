package res

import (
	"github.com/kaiouz/GoShare/config"
	"testing"
)

func TestCreateResources(t *testing.T) {
	if err := config.InitConfig(); err == nil {
		resources := CreateResources(config.Config.Dir)
		t.Logf("len: %d, %v", len(resources), resources)
	} else {
		t.Error(err)
	}
}
