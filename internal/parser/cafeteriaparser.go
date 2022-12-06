package parser

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/schicho/sabanci/data"
)

var ErrParseCafeteria = errors.New("could not parse cafeteria")

// cafeteriaResponse represents the HTML table containing the cafeteria menu.
type cafeteriaResponse struct {
	Rows []struct {
		Cells []struct {
			Text string `xml:",chardata"`
		} `xml:"td"`
	} `xml:"tr"`
}

// ParseCafeteria parses the HTML response from mySu.
// The reader must be an io.Reader of the HTML response.
// It returns a Cafeteria struct containing the cafeteria menu.
func ParseCafeteria(r io.Reader) (*data.Cafeteria, error) {
	food, err := parseCafeteriaResponse(r)
	if err != nil {
		log.Println(err)
		// remove context for the frontend
		return nil, ErrParseCafeteria
	}

	var menu []data.Food
	for _, row := range food.Rows {
		if len(row.Cells) != 2 {
			return nil, fmt.Errorf("%w: invalid number of cells: %d", ErrParseCafeteria, len(row.Cells))
		}
		menu = append(menu, data.Food{
			Name:     row.Cells[0].Text,
			Calories: row.Cells[1].Text,
		})
	}
	return &data.Cafeteria{Menu: menu}, nil
}

// parseCafeteriaResponse parses the HTML table containing the cafeteria menu.
// The table is parsed into a CafeteriaResponse struct.
// We use a XML decoder to parse the HTML.
func parseCafeteriaResponse(r io.Reader) (*cafeteriaResponse, error) {
	var cafeteriaResponse cafeteriaResponse

	d := xml.NewDecoder(r)
	d.Entity = xml.HTMLEntity
	d.AutoClose = xml.HTMLAutoClose
	d.Strict = false

	err := d.Decode(&cafeteriaResponse)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParseCafeteria, err)
	}
	return &cafeteriaResponse, nil
}
