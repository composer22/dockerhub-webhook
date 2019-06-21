package server

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

const (
	testPayloadJSON = `{"namespace":"development",
	"dockerhub":{
	  "callback_url": "https://registry.hub.docker.com/u/foo/testhook/hook/abced/",
	  "push_data": {
	    "tag": "1.2.3"
	  },
	  "repository": {
	    "name": "testhook",
	    "repo_name": "foo/testhook"
	  }
	}}`
)

var (
	testSrvr *Server
)

func TestServerStartup(t *testing.T) {
	opts := OptionsNew(func(o *Options) {
		o.Hostname = "localhost"
		o.ValidTokens = []string{"token101", "token102", "token103"}
		o.AlivePath = "/alivepath"
		o.NotifyPath = "/notifypath"
		o.StatusPath = "/statuspath"
		o.TargetHost = "myjenkins.test.com"
		o.TargetPath = "/testtargetpath/"
		o.TargetToken = "atokentester"
	})
	testSrvr = New(opts, func(s *Server) {})
	go func() { testSrvr.Start() }()
}

func TestMethods(t *testing.T) {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", "http://localhost:8080/alivepath", nil)
	resp, _ := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	body := strings.TrimSuffix(string(b), "\n")
	if body != InvalidMethod {
		t.Errorf("/alive body should return method error.")
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("/alive returned invalid method status code %d", resp.StatusCode)
	}

	req, _ = http.NewRequest("GET", "http://localhost:8080/notifypath", nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("token", "token999")
	req.URL.RawQuery = q.Encode()

	resp, _ = client.Do(req)
	b, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	body = strings.TrimSuffix(string(b), "\n")
	if body != InvalidMethod {
		t.Errorf("/notify body should return method error.")
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("/notify returned invalid method status code %d", resp.StatusCode)
	}

	req, _ = http.NewRequest("POST", "http://localhost:8080/statuspath", nil)
	resp, _ = client.Do(req)
	b, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	body = strings.TrimSuffix(string(b), "\n")
	if body != InvalidMethod {
		t.Errorf("/status body should return method error.")
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("/status returned invalid method status code %d", resp.StatusCode)
	}

}

func TestRoutes(t *testing.T) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "http://localhost:8080/alivepath", nil)
	resp, _ := client.Do(req)
	b, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	body := strings.TrimSuffix(string(b), "\n")
	if body != "" {
		t.Errorf("/alive body should be empty.")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("/alive returned invalid status code %d", resp.StatusCode)
	}

	req, _ = http.NewRequest("POST", "http://localhost:8080/notifypath",
		strings.NewReader(testPayloadJSON))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	q.Add("token", "token999")
	req.URL.RawQuery = q.Encode()

	resp, _ = client.Do(req)
	b, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	body = strings.TrimSuffix(string(b), "\n")
	if body != InvalidAuthorization {
		t.Errorf("/notify body should return authorization error.")
	}
	if resp.StatusCode != http.StatusUnauthorized {
		t.Errorf("/notify returned invalid authorization status code %d", resp.StatusCode)
	}

	// This doesnt test the whole flow as a backend server is needed. It should return the
	// correct error.
	req, _ = http.NewRequest("POST", "http://localhost:8080/notifypath",
		strings.NewReader(testPayloadJSON))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Content-Type", "application/json")
	q = req.URL.Query()
	q.Add("token", "token103")
	req.URL.RawQuery = q.Encode()

	resp, _ = client.Do(req)
	b, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	body = strings.TrimSuffix(string(b), "\n")
	if body != CouldNotProcess {
		t.Errorf("/notify body should return could not process error.")
	}
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("/notify returned invalid could not process status code %d", resp.StatusCode)
	}

	req, _ = http.NewRequest("GET", "http://localhost:8080/statuspath", nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	resp, _ = client.Do(req)
	b, _ = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	body = strings.TrimSuffix(string(b), "\n")
	if body == "" {
		t.Errorf("/status body should not be empty.")
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("/status returned invalid status code %d", resp.StatusCode)
	}

}

func TestServerPrintVersion(t *testing.T) {
	t.Parallel()
	t.Skip("Exit cannot be covered.")
}

func TestServerTakeDown(t *testing.T) {
	time.Sleep(2 * time.Second) // Coverage of timeout in Throttle.
	testSrvr.Shutdown()
	testSrvr.Shutdown() // Coverage of isRunning test in Shutdown().
	if testSrvr.isRunning() {
		t.Errorf("Server should have shut down.")
	}
	testSrvr = nil
}
