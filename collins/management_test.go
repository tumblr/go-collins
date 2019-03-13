package collins

import (
	"testing"
)

func TestManagementService_powerAction(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(200, "/api/asset/tag30/power", "../tests/management/generic_success.json", "application/json;", t)

	_, err := client.Management.powerAction("tag30", "powerOff")
	if err != nil {
		t.Errorf("ManagementService.powerAction returned error: %v", err)
	}
}

func TestManagementService_powerAction_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(404, "/api/asset/tag30/power", "../tests/management/invalid_asset.json", "application/json;", t)

	_, err := client.Management.powerAction("tag30", "powerOff")
	if err == nil {
		t.Errorf("ManagementService.powerAction did not return error on invalid asset.")
	}
}

// Make sure functions for all power actions exists
func TestManagementService_PowerActions(t *testing.T) {
	setup()
	defer teardown()

	functions := make(map[string]func(string) (*Response, error))
	functions["powerOff"] = client.Management.PowerOff
	functions["powerOn"] = client.Management.PowerOn
	functions["powerSoft"] = client.Management.SoftPowerOff
	functions["rebootSoft"] = client.Management.SoftReboot
	functions["rebootHard"] = client.Management.HardReboot
	functions["identify"] = client.Management.Identify
	functions["verify"] = client.Management.Verify

	SetupPOST(200, "/api/asset/tag30/power", "../tests/management/generic_success.json", "application/json;", t)

	for action, f := range functions {
		_, err := f("tag30")
		if err != nil {
			t.Errorf("%s returned error: %s", action, err)
		}
	}
}

func TestManagementService_PowerStatus(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/asset/tag30/power", "../tests/management/power_status.json", "application/json;", t)

	msg, _, err := client.Management.PowerStatus("tag30")
	if err != nil {
		t.Errorf("ManagementService.PowerStatus returned error: %v", err)
	}

	if msg != "on" {
		t.Errorf("ManagementService.PowerStatus returned \"%s\", want \"on\".", msg)
	}
}

func TestManagementService_PowerStatus_error(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(404, "/api/asset/tag30/power", "../tests/management/invalid_asset.json", "application/json;", t)

	_, _, err := client.Management.PowerStatus("tag30")
	if err == nil {
		t.Errorf("ManagementService.PowerStatus did not return error on invalid asset.")
	}
}

func TestManagementService_Provision(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(200, "/api/provision/tag30", "../tests/management/generic_success.json", "application/json;", t)

	opts := ProvisionOpts{}
	_, err := client.Management.Provision("tag30", "testprofile", "testuser", opts)
	if err != nil {
		t.Errorf("ManagementService.Provision returned error: %v", err)
	}
}

func TestManagementService_Provision_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(404, "/api/provision/tag30", "../tests/management/invalid_asset.json", "application/json;", t)

	opts := ProvisionOpts{}
	_, err := client.Management.Provision("tag30", "testprofile", "testuser", opts)
	if err == nil {
		t.Errorf("ManagementService.Provision did not return error on invalid asset.")
	}
}

func TestManagementService_GetProvisioningProfiles(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/provision/profiles", "../tests/management/provision_profiles.json", "application/json;", t)

	_, _, err := client.Management.GetProvisioningProfiles()
	if err != nil {
		t.Errorf("ManagementService.GetProvisionProfiles returned error: %v", err)
	}
}
