package web

import (
	"io"
	"time"
)

type SUConnector interface {
	Init()
	Login(username, password string) error
	SaveCookies() error
	SURetriever
}

type SURetriever interface {
	RetrieveCafeteria(time.Time) (io.ReadCloser, error)
	RetrieveSchedule(time.Time) (io.ReadCloser, error)
	RetrieveWallet() (io.ReadCloser, error)
}
