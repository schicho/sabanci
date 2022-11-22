package web

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func TestInitialization(t *testing.T) {
	client := &client{&http.Client{}}

	if client.Client == nil {
		t.Error("client.Client is nil")
	}
	if client.Jar != nil {
		t.Error("client.Jar should be nil")
	}

	client.Init()
	if client.Jar == nil {
		t.Error("client.Jar should not be nil")
	}
}

func TestLoginAndRetrieval(t *testing.T) {
	client := &client{&http.Client{}}
	client.Init()

	username := os.Getenv("TEST_USERNAME")
	password := os.Getenv("TEST_PASSWORD")

	if username == "" || password == "" {
		t.Error("TEST_USERNAME and TEST_PASSWORD must be set")
		t.SkipNow()
	}

	err := client.Login(username, password)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	r, err := client.RetrieveWallet()
	if err != nil {
		t.Fatalf("RetrieveWallet failed: %v", err)
	}
	defer r.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("ReadAll failed: %v", err)
	}

	if len(b) == 0 {
		t.Error("Wallet is empty")
	}

	if b[0] != '{' {
		t.Error("did not receive JSON")
		t.Error("got:\n", string(b))
	}
}

func TestSuconnectorLoginAndRetrieval(t *testing.T) {
	suconnector := SUConnector(&client{&http.Client{}})
	suconnector.Init()

	username := os.Getenv("TEST_USERNAME")
	password := os.Getenv("TEST_PASSWORD")

	if username == "" || password == "" {
		t.Error("TEST_USERNAME and TEST_PASSWORD must be set")
		t.SkipNow()
	}

	err := suconnector.Login(username, password)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	r, err := suconnector.RetrieveWallet()
	if err != nil {
		t.Fatalf("RetrieveWallet failed: %v", err)
	}
	defer r.Close()

	b, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("ReadAll failed: %v", err)
	}

	if len(b) == 0 {
		t.Error("Wallet is empty")
	}

	if b[0] != '{' {
		t.Error("did not receive JSON")
		t.Error("got:\n", string(b))
	}
}
