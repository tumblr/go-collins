package collins

import (
	"fmt"
	"strings"
	"testing"
)

func TestAssetService_Create(t *testing.T) {
	setup()
	defer teardown()

	SetupPUT(201, "/api/asset/tag30", "../tests/assets/create_success.json", "application/json;", t)

	asset, _, err := client.Assets.Create("tag30", nil)
	if err != nil {
		t.Errorf("Asset.Create returned error: %v", err)
	}
	if asset.Tag != "tag30" {
		t.Errorf("Asset tag %s, want tag30.", asset.Tag)
	}
}

func TestAssetService_Create_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPUT(409, "/api/asset/tag30", "../tests/assets/create_error.json", "application/json;", t)

	asset, _, err := client.Assets.Create("tag30", nil)
	if err == nil {
		t.Errorf("Asset.Create did not return error.")
	}
	if asset != nil {
		t.Errorf("Asset is %T, want nil.", asset)
	}
}

func TestAssetService_Create_error_ipmi_pool(t *testing.T) {
	setup()
	defer teardown()

	SetupPUT(400, "/api/asset/tag30", "../tests/assets/create_error_ipmi.json", "application/json;", t)

	asset, _, err := client.Assets.Create("tag30", nil)
	if err == nil {
		t.Errorf("Asset.Create did not return error.")
	}
	if asset != nil {
		t.Errorf("Asset is %T, want nil.", asset)
	}
}

func TestAssetService_Update(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(200, "/api/asset/tag30", "../tests/assets/generic_success.json", "application/json;", t)

	_, err := client.Assets.Update("tag30", nil)
	if err != nil {
		t.Errorf("Asset.Update returned error: %v", err)
	}
}

func TestAssetService_Update_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(404, "/api/asset/tag30", "../tests/assets/invalid_asset.json", "application/json;", t)

	_, err := client.Assets.Update("tag30", nil)
	if err == nil {
		t.Errorf("Asset.Update did not return error on duplicate asset.")
	}
}

func TestAssetService_AssetUpdateOpts_PowerConfig(t *testing.T) {
	opts := AssetUpdateOpts{
		PowerConfig: make(PowerConfig),
	}

	opts.PowerConfig["POWER_PORT_A"] = "port_a"
	opts.PowerConfig["POWER_PORT_B"] = "port_b"
	opts.PowerConfig["POWER_OUTLET_A"] = "outlet_a"
	opts.PowerConfig["POWER_OUTLET_B"] = "outlet_b"

	ustr, err := addOptions("api/test", &opts)

	if err != nil {
		t.Errorf("Error adding options: %s", err)
	}

	for k, v := range opts.PowerConfig {
		expected := fmt.Sprintf("%s=%s", k, v)
		if !strings.Contains(ustr, expected) {
			t.Errorf("query string contains %s, want %s", ustr, expected)
		}
	}
}

func TestAssetService_UpdateStatus(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(200, "/api/asset/tag30/status", "../tests/assets/generic_success.json", "application/json;", t)

	updateOpts := AssetUpdateStatusOpts{Status: "Maintenance", State: "Test", Reason: "Testing"}
	_, err := client.Assets.UpdateStatus("tag30", &updateOpts)

	if err != nil {
		t.Errorf("Asset.UpdateStatus return error: %v", err)
	}
}

func TestAssetService_UpdateStatus_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(400, "/api/asset/tag30/status", "../tests/assets/updatestatus_error.json", "application/json;", t)

	updateOpts := AssetUpdateStatusOpts{Reason: "Testing"}
	_, err := client.Assets.UpdateStatus("tag30", &updateOpts)

	if err == nil {
		t.Errorf("Asset.UpdateStatus did not return error when state and status not specified.")
	}
}

func TestAssetService_Get(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/asset/tag30", "../tests/assets/get_success.json", "application/json;", t)

	asset, _, err := client.Assets.Get("tag30")

	if err != nil {
		t.Fatalf("Asset.Get return error: %v", err)
	}
	if asset.Tag != "tag30" {
		t.Errorf("Asset tag: %v, want tag30", asset.Tag)
	}
}

func TestAssetService_Get_error(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(404, "/api/asset/tag30", "../tests/assets/invalid_asset.json", "application/json;", t)

	asset, _, err := client.Assets.Get("tag30")

	if err == nil {
		t.Errorf("Asset.Get did not return error.")
	}
	if asset != nil {
		t.Errorf("Asset is %T, want nil.", asset)
	}
}

func TestAssetService_Delete(t *testing.T) {
	setup()
	defer teardown()

	SetupDELETE(200, "/api/asset/tag30", "../tests/assets/generic_success.json", "application/json;", t)

	_, err := client.Assets.Delete("tag30", "Testing")

	if err != nil {
		t.Errorf("Asset.Delete return error: %v", err)
	}
}

func TestAssetService_Delete_error(t *testing.T) {
	setup()
	defer teardown()

	SetupDELETE(404, "/api/asset/tag30", "../tests/assets/invalid_asset.json", "application/json;", t)

	_, err := client.Assets.Delete("tag30", "Testing")

	if err == nil {
		t.Errorf("Asset.Delete did not return error for unknown asset.")
	}
}

func TestAssetService_UpdateIpmi(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(200, "/api/asset/tag30/ipmi", "../tests/assets/generic_success.json", "application/json;", t)

	_, err := client.Assets.UpdateIpmi("tag30", nil)

	if err != nil {
		t.Errorf("Asset.UpdateIpmi return error: %v", err)
	}
}

func TestAssetService_UpdateIpmi_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(404, "/api/asset/tag30/ipmi", "../tests/assets/invalid_asset.json", "application/json;", t)

	_, err := client.Assets.UpdateIpmi("tag30", nil)

	if err == nil {
		t.Errorf("Asset.UpdateIpmi did not return error for unknown asset.")
	}
}
