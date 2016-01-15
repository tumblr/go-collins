package collins

import (
	"testing"
)

func TestLogService_Create(t *testing.T) {
	setup()
	defer teardown()

	SetupPUT(201, "/api/asset/tag1/log", "../tests/logs/log_create_success.json", "application/json;", t)

	logOpts := LogCreateOpts{Message: "Hello World", Type: "DEBUG"}
	log, _, err := client.Logs.Create("tag1", &logOpts)
	if err != nil {
		t.Errorf("Logs.Create returned error: %v", err)
	}

	if log.Message != "Hello World" {
		t.Errorf("Log message: %s, want \"Hello World\".", log.Message)
	}
}

func TestLogService_Create_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPUT(404, "/api/asset/tag1/log", "../tests/logs/log_create_error.json", "application/json;", t)

	logOpts := LogCreateOpts{Message: "Hello World", Type: "DEBUG"}
	_, _, err := client.Logs.Create("tag1", &logOpts)
	if err == nil {
		t.Errorf("Logs.Create did not return error.")
	}
}

func TestLogService_Get(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/asset/tag1/logs", "../tests/logs/logs.json", "application/json;", t)

	logs, _, err := client.Logs.Get("tag1", nil)
	if err != nil {
		t.Errorf("Logs.Get returned error: %v", err)
	}

	if len(logs) != 3 {
		t.Errorf("Number of logs: %d, want 3", len(logs))
	}
}

func TestLogService_Get_error(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(404, "/api/asset/tag1/logs", "../tests/logs/logs_error.json", "application/json;", t)

	_, _, err := client.Logs.Get("tag1", nil)
	if err == nil {
		t.Errorf("Logs.Get did not return error.")
	}
}

func TestLogService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/assets/logs", "../tests/logs/logs.json", "application/json;", t)

	logs, _, err := client.Logs.GetAll(nil)
	if err != nil {
		t.Errorf("Logs.GetAll returned error: %v", err)
	}

	if len(logs) != 3 {
		t.Errorf("Number of logs: %d, want 3", len(logs))
	}
}
