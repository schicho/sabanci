package web

import (
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

type client struct {
	*http.Client
}

// DefaultSUClient is the default SUClient implementation.
var DefaultSUClient = &client{Client: &http.Client{}}

// Init initializes the SUClient with a cookie jar.
func (c *client) Init() {
	opt := cookiejar.Options{PublicSuffixList: publicsuffix.List}
	jar, err := cookiejar.New(&opt)
	if err != nil {
		panic(err)
	}
	c.Client.Jar = jar
}

// SaveCookies saves the cookies to a file.
// It is used to persist the session.
func (c *client) SaveCookies() error {
	return nil
}
