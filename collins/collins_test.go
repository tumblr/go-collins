package collins

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client, _ = NewClient("test", "test", "test")
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func SetupMethod(code int, method, url, file string, contentType string, t *testing.T) {
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			t.Errorf("Request method: %v, want %s", r.Method, method)
		}
		resp, err := ioutil.ReadFile(file)
		if err != nil {
			t.Errorf("Could not read %s\n", file)
		}
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(code)
		fmt.Fprintf(w, "%s", resp)
	})
}

func SetupFirehose(file, contentType string, t *testing.T) {
	mux.HandleFunc("/api/firehose", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Request method: %v, want GET", r.Method)
		}
		resp, err := ioutil.ReadFile(file)
		if err != nil {
			t.Errorf("Could not read %s\n", file)
		}
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", resp)
	})
}

func SetupGET(code int, url, file string, contentType string, t *testing.T) {
	SetupMethod(code, "GET", url, file, contentType, t)
}

func SetupPUT(code int, url, file string, contentType string, t *testing.T) {
	SetupMethod(code, "PUT", url, file, contentType, t)
}

func SetupPOST(code int, url, file string, contentType string, t *testing.T) {
	SetupMethod(code, "POST", url, file, contentType, t)
}

func SetupDELETE(code int, url, file string, contentType string, t *testing.T) {
	SetupMethod(code, "DELETE", url, file, contentType, t)
}

func TestNewClient(t *testing.T) {
	client, err := NewClient("testuser", "testpassword", "https://collins.example.net")
	if err != nil {
		t.Errorf("Failed to create client: %s", err)
	}

	if client.User != "testuser" || client.Password != "testpassword" {
		t.Errorf("Failed to parse user or password")
	}
}

func TestNewClientFromYaml(t *testing.T) {
	err := os.Setenv("COLLINS_CLIENT_CONFIG", "../tests/test_config.yml")
	if err != nil {
		t.Errorf("Failed to set COLLINS_CLIENT_CONFIG environment variable.")
	}

	client, err := NewClientFromYaml()
	if err != nil {
		t.Errorf("Failed to create client: %s", err)
	}

	if client.User != "testuser" || client.Password != "testpassword" {
		t.Errorf("Failed to parse user or password")
	}
}

func TestNewClientFromYaml_error(t *testing.T) {
	err := os.Setenv("COLLINS_CLIENT_CONFIG", "../tests/test_config_non_existent.yml")
	if err != nil {
		t.Errorf("Failed to set COLLINS_CLIENT_CONFIG environment variable.")
	}

	err = os.Setenv("HOME", "/dev/null")
	if err != nil {
		t.Errorf("Failed to set HOME environment variable.")
	}

	_, err = NewClientFromYaml()
	if err == nil {
		t.Errorf("Did not throw error with non-existent config file.")
	}
}
