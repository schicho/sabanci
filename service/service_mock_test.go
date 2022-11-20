// This file implements a mocked SUConnector.
// It always returns success and works of testdata.
package service

import (
	"io"
	"os"
	"testing"
	"time"
)

// successTestSUConn is a mocked SUConnector implementation for testing.
// It always returns success and testdata.
type successTestSUConn struct{}

func (t successTestSUConn) Init() {}

func (t successTestSUConn) SaveCookies() error {
	return nil
}

func (t successTestSUConn) Login(username, password string) error {
	return nil
}

func (t successTestSUConn) RetrieveCafeteria(time.Time) (io.ReadCloser, error) {
	f, err := os.Open("../internal/parser/testdata/cafeteria.html")
	if err != nil {
		panic(err)
	}
	return f, nil
}

func (t successTestSUConn) RetrieveSchedule(time.Time) (io.ReadCloser, error) {
	f, err := os.Open("../internal/parser/testdata/schedule.html")
	if err != nil {
		panic(err)
	}
	return f, nil
}

func (t successTestSUConn) RetrieveWallet() (io.ReadCloser, error) {
	f, err := os.Open("../internal/parser/testdata/wallet.json")
	if err != nil {
		panic(err)
	}
	return f, nil
}

func createSuccessTestService() service {
	return service{
		suConn: successTestSUConn{},
	}
}

func TestCafeteriaService(t *testing.T) {
	ts := createSuccessTestService()
	_, err := ts.GetCafeteria()
	if err != nil {
		t.Error(err)
	}
}

func TestScheduleService(t *testing.T) {
	ts := createSuccessTestService()
	_, err := ts.GetSchedule()
	if err != nil {
		t.Error(err)
	}
}

func TestWalletService(t *testing.T) {
	ts := createSuccessTestService()
	_, err := ts.GetWallet()
	if err != nil {
		t.Error(err)
	}
}
