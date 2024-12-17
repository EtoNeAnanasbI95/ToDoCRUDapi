package storage

import (
	"github.com/EtoNeAnanasbI95/ToDoCRUD/internal/config"
	"os"
	"testing"
)

func TestMustInitDB(t *testing.T) {
	os.Setenv("CONFIG_PATH", "../../config/local.yaml")
	cfg := config.MustLoadConfig()
	db := MustInitDB(cfg.ConnectionString)
	if db == nil {
		t.Error("db init failed")
	}
}
