package service

import (
	"time"

	"github.com/schicho/sabanci/data"
	"github.com/schicho/sabanci/internal/parser"
	"github.com/schicho/sabanci/internal/web"
)

type service struct {
	suConn web.SUConnector
}

func init() {
	defaultService.suConn.Init()
}

var defaultService = &service{suConn: web.DefaultSUClient}

// Login performs the login to the SU website via the default SU client.
func Login(username, password string) error {
	return defaultService.Login(username, password)
}

func (s *service) Login(username, password string) error {
	return s.suConn.Login(username, password)
}

// SaveCookies saves the cookies of the default service and default SU client.
func SaveCookies() error {
	return defaultService.SaveCookies()
}

func (s *service) SaveCookies() error {
	return s.suConn.SaveCookies()
}

// GetCafeteria returns the cafeteria menu for the current day.
// It performs the web request to the cafeteria endpoint.
func GetCafeteria() (*data.Cafeteria, error) {
	return defaultService.GetCafeteria()
}

func (s *service) GetCafeteria() (*data.Cafeteria, error) {
	r, err := s.suConn.RetrieveCafeteria(time.Now())
	if err != nil {
		return nil, err
	}
	defer r.Close()

	cafeteria, err := parser.ParseCafeteria(r)
	if err != nil {
		return nil, err
	}
	return cafeteria, nil
}

// GetSchedule returns the schedule of the logged in user.
func GetSchedule() (*data.Schedule, error) {
	return defaultService.GetSchedule()
}

func (s *service) GetSchedule() (*data.Schedule, error) {
	r, err := s.suConn.RetrieveSchedule(time.Now())
	if err != nil {
		return nil, err
	}
	defer r.Close()

	schedule, err := parser.ParseSchedule(r)
	if err != nil {
		return nil, err
	}
	return schedule, nil
}

// GetWallet returns the user's wallet balances.
// It performs the web request to the wallet endpoint.
func GetWallet() (*data.Wallet, error) {
	return defaultService.GetWallet()
}

func (s *service) GetWallet() (*data.Wallet, error) {
	r, err := s.suConn.RetrieveWallet()
	if err != nil {
		return nil, err
	}
	defer r.Close()

	wallet, err := parser.ParseWallet(r)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}
