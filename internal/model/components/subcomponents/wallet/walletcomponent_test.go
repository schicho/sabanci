package wallet

import (
	"fmt"
	"strings"
	"testing"

	"github.com/schicho/sabanci/data"
)

func TestWalletComponentSize(t *testing.T) {
	wallet := &data.Wallet{Meal: 12.3, Print: 45, Shuttle: 6.7}
	m := Model{wallet: wallet}

	s := m.View()

	if strings.Count(s, "\n") != 2 {
		t.Errorf("Expected component to be 3 lines tall, but got %v newlines", strings.Count(s, "\n"))
	}
}

func TestWalletComponentError(t *testing.T) {
	m := Model{err: fmt.Errorf("test error")}

	s := m.View()

	if !strings.Contains(s, "Error: test error") {
		t.Errorf("Expected error message to be rendered, but got %v", s)
	}
}
