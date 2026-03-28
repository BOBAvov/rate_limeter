package function

import (
	"net/http"
	"testing"

	"github.com/BOBAvov/rate_limeter/tets/function/suite"
)

func TestHealthCheck(t *testing.T) {
	s := suite.New()

	req, err := http.NewRequest("GET", s.URL+"/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := s.Client.Do(req)
	if err != nil {
		t.Fatalf("failed to execute request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %v", res.StatusCode)
	}
	t.Logf("response: %v", res)
}
