package suite

import "net/http"

const port = "8080"

type Suite struct {
	Client *http.Client
	URL    string
}

func New() *Suite {
	return &Suite{
		Client: &http.Client{},
		URL:    "http://localhost:" + port,
	}
}
