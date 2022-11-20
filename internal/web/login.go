package web

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

const (
	loginPageURL = `https://login.sabanciuniv.edu/cas/login?service=https%3A%2F%2Fmysu.sabanciuniv.edu%2Fen%2Fnode`
)

var (
	ErrLoginFailed   = errors.New("login failed")
	ErrNoCredentials = errors.New("no credentials provided")
)

// Login logs in to mysu.sabanciuniv.edu using the given username and
// password. It returns an error if the login fails.
func (c *client) Login(username, password string) error {
	resp, err := c.loadLoginPage()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrLoginFailed, err)
	}
	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrLoginFailed, err)
	}

	// we can short-circuit the login process if the session can be restored
	// from the cookies.
	if sessionRestored(root) {
		return nil
	}

	// We only check for the credentials here, because we want to be able to
	// restore the session from the cookies if possible.
	if username == "" || password == "" {
		return ErrNoCredentials
	}

	// The execution random is a random string that is generated by the
	// server and is set as a value of a hidden form input field. We need
	// to find this value and send it back to the server when we post the
	// login form.
	executionRandom, err := findExecutionRandom(root)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrLoginFailed, err)
	}

	respLogin, err := c.postLogin(username, password, executionRandom)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrLoginFailed, err)
	}
	defer respLogin.Body.Close()

	// If the login is successful, the server redirects us to the
	// dashboard page. If the login is unsuccessful, the server
	// redirects us to the login page again.
	if respLogin.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("%w: %v", ErrLoginFailed, "unauthorized")
	}

	root, err = html.Parse(respLogin.Body)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrLoginFailed, err)
	}

	if !sessionRestored(root) {
		return fmt.Errorf("%w: %v", ErrLoginFailed, "session could not be established")
	}

	return nil
}

func (c *client) loadLoginPage() (*http.Response, error) {
	return c.Get(loginPageURL)
}

func (c *client) postLogin(username, password, executionRandom string) (*http.Response, error) {
	return c.PostForm(loginPageURL, url.Values{
		"username":    {username},
		"password":    {password},
		"execution":   {executionRandom},
		"_eventId":    {"submit"},
		"geolocation": {""},
	})
}

// findExecutionRandom iterates over the nodes of the HTML tree and returns
// the value of the execution random form input field.
func findExecutionRandom(node *html.Node) (string, error) {
	if node.Type == html.ElementNode && node.Data == "input" {
		for _, a := range node.Attr {
			if a.Key == "name" && a.Val == "execution" {
				for _, a := range node.Attr {
					if a.Key == "value" {
						return a.Val, nil
					}
				}
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if val, err := findExecutionRandom(c); err == nil {
			return val, nil
		}
	}

	return "", errors.New("execution random not found")
}

// sessionRestored checks if the session is restored by checking if
// we are redirected to the mysu dashboard page directly on the first
// request to the CAS.
// We do this by checking if the HTML title is <title>mySU</title>.
// This is a hacky way of doing it, but it works.
func sessionRestored(node *html.Node) bool {
	if node.Type == html.ElementNode && node.Data == "title" {
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				return c.Data == "mySU"
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if sessionRestored(c) {
			return true
		}
	}

	return false
}
