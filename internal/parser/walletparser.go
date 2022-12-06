package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/schicho/sabanci/data"
)

var ErrParseWallet = errors.New("could not parse wallet")

// walletResponse is the JSON response from the wallet endpoint.
// It contains a lot of metadata about the user's wallet.
// We only care about the sum of the shuttle, meal, and print.
type walletResponse struct {
	Shuttle struct {
		Errcode int     `json:"errcode"`
		Sum     float64 `json:"sum"`
	} `json:"shuttle"`
	Meal struct {
		Errcode int     `json:"errcode"`
		Sum     float64 `json:"sum"`
	} `json:"meal"`
	Print struct {
		Errcode int     `json:"errcode"`
		Sum     float64 `json:"sum"`
	} `json:"print"`
	UID string `json:"uid"`
	OU  string `json:"ou"`
}

// ParseWallet parses the JSON response from mySU.
// The response is parsed into a Wallet struct.
func ParseWallet(r io.Reader) (*data.Wallet, error) {
	walletResponse, err := parseWalletResponse(r)
	if err != nil {
		log.Println(err)
		// remove context for the frontend
		return nil, ErrParseSchedule
	}

	wallet := &data.Wallet{
		Shuttle: walletResponse.Shuttle.Sum,
		Meal:    walletResponse.Meal.Sum,
		Print:   walletResponse.Print.Sum,
	}
	return wallet, nil
}

// parseWalletResponse parses the JSON response from the wallet endpoint.
// It fully parses the JSON response into a walletResponse struct.
func parseWalletResponse(r io.Reader) (*walletResponse, error) {
	var walletResponse walletResponse
	err := json.NewDecoder(r).Decode(&walletResponse)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParseWallet, err)
	}
	return &walletResponse, nil
}
