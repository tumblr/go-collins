package collins

import (
	"testing"
)

func TestErrorResponseAsPlainText(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(401, "/api/asset/tag36", "../tests/invalid_something.txt", "text/plain;", t)
	_, _, err := client.Assets.Get("tag36")

	if err == nil {
		t.Errorf("Did not receive an error for plain text error response.")
	}
}

func TestResponseAsPlainText(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(201, "/api/ping", "../tests/ping_status.txt", "text/plain;", t)
	req, _ := client.NewRequest("GET", "/api/ping")

	_, err := client.Do(req, nil)

	if err == nil {
		t.Errorf("Collins Health Check did not return a Content-Type error")
	}
}
