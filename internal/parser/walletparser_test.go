package parser

import (
	"bytes"
	"testing"
)

const testJSON = `{"shuttle":{"errcode":0,"sum":123},"meal":{"errcode":0,"sum":456.7},"print":{"errcode":0,"sum":8.9},"uid":"firstname.lastname","ou":"student"}`

func TestParseWallet(t *testing.T) {
	b := bytes.NewBuffer([]byte(testJSON))

	wallet, err := ParseWallet(b)
	if err != nil {
		t.Error(err)
	}
	if wallet.Shuttle != 123 {
		t.Errorf("wallet.Shuttle = %f; want 123", wallet.Shuttle)
	}
	if wallet.Meal != 456.7 {
		t.Errorf("wallet.Meal = %f; want 456.7", wallet.Meal)
	}
	if wallet.Print != 8.9 {
		t.Errorf("wallet.Print = %f; want 0", wallet.Print)
	}
}

func TestParseWalletResponse(t *testing.T) {
	b := bytes.NewBuffer([]byte(testJSON))

	walletResponse, err := parseWalletResponse(b)
	if err != nil {
		t.Error(err)
	}
	if walletResponse.Shuttle.Sum != 123 {
		t.Errorf("walletResponse.Shuttle.Sum = %f; want 123", walletResponse.Shuttle.Sum)
	}
	if walletResponse.Meal.Sum != 456.7 {
		t.Errorf("walletResponse.Meal.Sum = %f; want 456.7", walletResponse.Meal.Sum)
	}
	if walletResponse.Print.Sum != 8.9 {
		t.Errorf("walletResponse.Print.Sum = %f; want 0", walletResponse.Print.Sum)
	}
	if walletResponse.UID != "firstname.lastname" {
		t.Errorf("walletResponse.UID = %s; want firstname.lastname", walletResponse.UID)
	}
	if walletResponse.OU != "student" {
		t.Errorf("walletResponse.OU = %s; want student", walletResponse.OU)
	}
}
