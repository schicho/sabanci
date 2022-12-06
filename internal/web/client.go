package web

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path"

	pcj "github.com/orirawlings/persistent-cookiejar"
)

type client struct {
	*http.Client
}

const suJarName = "su-jar.json"

var (
	pathname = path.Join(os.TempDir(), suJarName)
	options  = pcj.Options{Filename: pathname, PersistSessionCookies: true}
)

// DefaultSUClient is the default SUClient implementation.
var DefaultSUClient = &client{Client: &http.Client{}}

// Init initializes the SUClient with a cookie jar.
func (c *client) Init() {
	var jar http.CookieJar

	jar, err := pcj.New(&options)
	if err != nil {
		log.Println("could not create persistent cookie jar:", err)
		// use un-persisted cookie jar
		jar, err = cookiejar.New(nil)
		if err != nil {
			panic(err)
		}
	}
	c.Client.Jar = jar
}

// SaveCookies saves the cookies to a temporary file in the system's
// temporary directory.
//
// These cookies can be reloaded by calling Init() and will be used
// to sign in the user automatically, without the need to enter the
// user's credentials.
//
// When the creation of a persistent cookie jar failed in Init,
// this method does nothing.
func (c *client) SaveCookies() error {
	pJar, ok := c.Client.Jar.(*pcj.Jar)
	if !ok {
		return nil
	}
	err := pJar.Save()
	if err != nil {
		log.Println("could not save cookies:", err)
	}
	return err
}
