package service

import (
	"os"
	"testing"

	"github.com/schicho/sabanci/internal/web"
)

func TestService(t *testing.T) {
	// using the default su client
	s := &service{suConn: web.DefaultSUClient}
	s.suConn.Init()

	username := os.Getenv("TEST_USERNAME")
	password := os.Getenv("TEST_PASSWORD")

	if username == "" || password == "" {
		t.Error("TEST_USERNAME and TEST_PASSWORD must be set")
		t.SkipNow()
	}

	err := s.Login(username, password)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	t.Run("GetWallet", func(t *testing.T) {
		w, err := s.GetWallet()
		if err != nil {
			t.Fatalf("GetWallet failed: %v", err)
		}
		t.Logf("Wallet: %v", w)
	})

	t.Run("GetCafeteria", func(t *testing.T) {
		c, err := s.GetCafeteria()
		if err != nil {
			t.Fatalf("GetCafeteria failed: %v", err)
		}
		t.Logf("Cafeteria: %v", c)
	})

	t.Run("GetSchedule", func(t *testing.T) {
		s, err := s.GetSchedule()
		if err != nil {
			t.Fatalf("GetSchedule failed: %v", err)
		}
		t.Logf("Schedule: %v", s)
	})
}
