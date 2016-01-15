package collins

import (
	"testing"
)

func TestStateService_Create(t *testing.T) {
	setup()
	defer teardown()

	SetupPUT(201, "/api/state/TEST_STATE", "../tests/states/generic_success.json", "application/json;", t)

	_, err := client.States.Create("TEST_STATE", "Test state", "State used in test suite", "")
	if err != nil {
		t.Errorf("State.Create returned error: %v", err)
	}
}

func TestStateService_Create_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPUT(409, "/api/state/TEST_STATE", "../tests/states/generic_error.json", "application/json;", t)

	_, err := client.States.Create("TEST_STATE", "Test state", "State used in test suite", "")
	if err == nil {
		t.Errorf("State.Create did not return error.")
	}
}

func TestStateService_Update(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(200, "/api/state/TEST_STATE", "../tests/states/generic_success.json", "application/json;", t)

	opts := StateUpdateOpts{Label: "My new label"}
	_, err := client.States.Update("TEST_STATE", opts)
	if err != nil {
		t.Errorf("State.Update returned error: %v", err)
	}
}

func TestStateService_Update_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(404, "/api/state/NO_SUCH_STATE", "../tests/states/generic_error.json", "application/json;", t)

	opts := StateUpdateOpts{Label: "My new label"}
	_, err := client.States.Update("NO_SUCH_STATE", opts)
	if err == nil {
		t.Errorf("State.Update did not return error.")
	}
}

func TestStateService_Delete(t *testing.T) {
	setup()
	defer teardown()

	SetupDELETE(200, "/api/state/TEST_STATE", "../tests/states/delete_success.json", "application/json;", t)

	_, err := client.States.Delete("TEST_STATE")
	if err != nil {
		t.Errorf("State.Delete returned error: %v", err)
	}
}

func TestStateService_Delete_error(t *testing.T) {
	setup()
	defer teardown()

	SetupDELETE(404, "/api/state/NO_SUCH_STATE", "../tests/states/delete_error.json", "application/json;", t)

	_, err := client.States.Delete("NO_SUCH_STATE")
	if err == nil {
		t.Errorf("State.Delete did not return error.")
	}
}

func TestStateService_Get(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/state/HW_ON_FIRE", "../tests/states/hw_on_fire.json", "application/json;", t)

	state, _, err := client.States.Get("HW_ON_FIRE")
	if err != nil {
		t.Errorf("State.Get returned error: %v", err)
	}

	if state.Label != "Hardware on Fire!" {
		t.Errorf("State label: %v, want \"Hardware on Fire!\".", state.Label)
	}
}

func TestStateService_Get_error(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(404, "/api/state/NO_SUCH_STATE", "../tests/states/get_error.json", "application/json;", t)

	_, _, err := client.States.Get("NO_SUCH_STATE")
	if err == nil {
		t.Errorf("State.Get did not return error.")
	}
}

func TestStateService_List(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/states", "../tests/states/all.json", "application/json;", t)

	states, _, err := client.States.List()
	if err != nil {
		t.Errorf("States.List returned error: %v", err)
	}

	if len(states) != 14 {
		t.Errorf("Number of states: %d, want 14", len(states))
	}
}
