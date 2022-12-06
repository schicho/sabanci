package parser

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/schicho/sabanci/data"
)

var ErrParseSchedule = errors.New("could not parse schedule")

// scheduleResponse is the full HTML response from the schedule endpoint.
// It is a huge HTML table containing the schedule.
// We assume the HTML to be XML, and parse it with an XML decoder.
// Example data can be seen in testdata/schedule.html.
type scheduleResponse struct {
	TableHead struct {
		Rows []struct {
			Headers []struct {
				A struct {
					Text string `xml:",chardata"`
				} `xml:"a"`
			} `xml:"th"`
		} `xml:"tr"`
	} `xml:"thead"`
	TableBody struct {
		Rows []struct {
			Cells []struct {
				DivSchedule struct {
					DivsDays []struct { // Starts with Monday, ends with Sunday
						Table struct {
							Rows []struct {
								Cells []struct {
									Text string `xml:",chardata"`
								} `xml:"td"`
							} `xml:"tr"`
						} `xml:"table"`
					} `xml:"div"`
				} `xml:"div"`
			} `xml:"td"`
		} `xml:"tr"`
	} `xml:"tbody"`
}

// ParseSchedule parses the HTML response from mySU.
// The reader must be an io.Reader of the HTML response.
// It returns a Schedule struct containing the schedule.
func ParseSchedule(r io.Reader) (*data.Schedule, error) {
	scheduleResponse, err := parseScheduleResponse(r)
	if err != nil {
		log.Println(err)
		// remove context for the frontend
		return nil, ErrParseSchedule
	}

	var schedule data.Schedule

	// We iterate over the DivsDays, which contain the schedule for each day.
	// Each DivsDay contains a table with the schedule for that day.
	// In each table, two consecutive rows represent one class.
	// The first row contains the name and start time of the class.
	// The second row contains the end time and building of the class.
	weekdays := scheduleResponse.TableBody.Rows[0].Cells[0].DivSchedule.DivsDays
	if len(weekdays) != 7 {
		log.Println(fmt.Errorf("%w: expected 7 days, got %d", ErrParseSchedule, len(weekdays)))
		return nil, ErrParseSchedule
	}
	for i, day := range weekdays {
		if len(day.Table.Rows)%2 != 0 {
			log.Println(fmt.Errorf("%w: expected even number of rows, got %d", ErrParseSchedule, len(day.Table.Rows)))
			return nil, ErrParseSchedule
		}
		for j, row := range day.Table.Rows {
			if j%2 == 0 {
				// On even rows a new class starts.
				// Even rows contain the name and start time of the class.
				schedule.Classes[i] = append(schedule.Classes[i], data.Class{
					Name: row.Cells[0].Text,
					// Cells[1] is empty.
					TimeStart: row.Cells[2].Text,
				})
			} else {
				// Odd rows contain the end time and building of the class.
				schedule.Classes[i][j/2].ClassCode = row.Cells[0].Text
				schedule.Classes[i][j/2].Building = row.Cells[1].Text
				schedule.Classes[i][j/2].TimeEnd = row.Cells[2].Text
			}
		}
	}

	return &schedule, nil
}

// parseScheduleResponse parses the HTML response from mySU.
// The reader must be an io.Reader of the HTML response.
// It returns a scheduleResponse struct containing the schedule.
func parseScheduleResponse(r io.Reader) (*scheduleResponse, error) {
	var scheduleResponse scheduleResponse

	d := xml.NewDecoder(r)
	d.Entity = xml.HTMLEntity
	d.AutoClose = xml.HTMLAutoClose
	d.Strict = false

	err := d.Decode(&scheduleResponse)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrParseSchedule, err)
	}
	return &scheduleResponse, nil
}
