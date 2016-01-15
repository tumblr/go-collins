package collins

import (
	"testing"
)

func TestAssetService_GetAttribute(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/asset/tag30", "../tests/assets/get_success.json", "application/json;", t)

	value, err := client.Assets.GetAttribute("tag30", "NODECLASS")

	if err != nil {
		t.Errorf("AssetService.GetAttribute returned error.")
	}
	if value != "web" {
		t.Errorf("AssetService.GetAttribute returned %s, want web.", value)
	}
}

func TestAssetService_GetAttribute_error(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/asset/tag30", "../tests/assets/get_success.json", "application/json;", t)

	_, err := client.Assets.GetAttribute("tag30", "NOSUCHATTRIBUTE")

	if err == nil {
		t.Errorf("AssetService.GetAttribute did not return error.")
	}
}

func TestAssetService_SetAttribute(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(200, "/api/asset/tag30", "../tests/assets/generic_success.json", "application/json;", t)

	_, err := client.Assets.SetAttribute("tag30", "TEST_ATTRIBUTE", "foobar")
	if err != nil {
		t.Errorf("Asset.SetAttribute returned error: %v", err)
	}
}

func TestAssetService_SetAttribute_error(t *testing.T) {
	_, err := client.Assets.SetAttribute("", "TEST_ATTRIBUTE", "foobar")
	if err == nil {
		t.Errorf("Asset.SetAttribute did not return error with empty tag.")
	}
}

func TestAssetService_DeleteAttribute(t *testing.T) {
	setup()
	defer teardown()

	SetupDELETE(202, "/api/asset/tag30/attribute/NODECLASS", "../tests/assets/generic_success.json", "application/json;", t)

	_, err := client.Assets.DeleteAttribute("tag30", "NODECLASS")

	if err != nil {
		t.Errorf("AssetService.DeleteAttribute returned error: %s", err)
	}
}

func TestAssetService_DeleteAttribute_error(t *testing.T) {
	setup()
	defer teardown()

	SetupDELETE(202, "/api/asset/tag30/attribute/NODECLASS", "../tests/assets/generic_success.json", "application/json;", t)

	_, err := client.Assets.DeleteAttribute("tag30", "NODECLASS")

	if err != nil {
		t.Errorf("AssetService.DeleteAttribute returned error: %s", err)
	}
}
