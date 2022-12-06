// This files contains the functions to download the schedule,
// cafeteria menu, and wallet information from mysu.sabanciuniv.edu.
package web

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	walletURL = `https://mysu.sabanciuniv.edu/en/ajax/getSUCard`

	// example: https://mysu.sabanciuniv.edu/en/ajax/getMeal?_=1668094026473
	cafeteriaURL = `https://mysu.sabanciuniv.edu/en/ajax/getMeal`

	// example: https://mysu.sabanciuniv.edu/en/ajax/getCourseSchedule?termcode=202201&_=1668097502097
	scheduleURL = `https://mysu.sabanciuniv.edu/en/ajax/getCourseSchedule`
)

var ErrRequestData = errors.New("could not download data")

func (c *client) getRequestBodyReader(URL string) (io.ReadCloser, error) {
	// can use c.Get() here later if I am done debugging.
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Println(fmt.Errorf("%w: %v", ErrRequestData, err))
		return nil, ErrRequestData
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Println(fmt.Errorf("%w: %v", ErrRequestData, err))
		return nil, ErrRequestData
	}
	if resp.StatusCode != 200 {
		log.Println(fmt.Errorf("%w: got status code %v", ErrRequestData, resp.Status))
		return nil, ErrRequestData
	}

	return resp.Body, nil
}

// RetrieveSchedule downloads the schedule from mysu.sabanciuniv.edu.
// The data is unparsed and returned as a reader.
// The ReadCloser must be closed by the caller.
//
// Currently (2022-11-07) the schedule is a HTML document.
func (c *client) RetrieveSchedule(time time.Time) (io.ReadCloser, error) {
	unix := time.Unix()
	year := time.Year()

	// It seems that the termcode is the current year and 01 for fall?
	// However, in my limited tests, any value seems to have the same effect.
	return c.getRequestBodyReader(scheduleURL +
		"?termcode=" + strconv.Itoa(year) + "01" +
		"&_=" + strconv.FormatInt(unix, 10))
}

// RetrieveCafeteria downloads the cafeteria menu from mysu.sabanciuniv.edu.
// The data is unparsed and returned as a reader.
// The ReadCloser must be closed by the caller.
//
// Currently (2022-11-07) the cafeteria menu is a HTML document.
func (c *client) RetrieveCafeteria(time time.Time) (io.ReadCloser, error) {
	unix := time.Unix()
	return c.getRequestBodyReader(cafeteriaURL + "?_=" + strconv.FormatInt(unix, 10))
}

// RetrieveWallet downloads the wallet information from mysu.sabanciuniv.edu.
// The data is unparsed and returned as a reader.
// The ReadCloser must be closed by the caller.
//
// Currently (2022-11-07) the wallet information is a JSON document.
func (c *client) RetrieveWallet() (io.ReadCloser, error) {
	unixnow := time.Now().Unix()
	return c.getRequestBodyReader(walletURL + "?_=" + strconv.FormatInt(unixnow, 10))
}
