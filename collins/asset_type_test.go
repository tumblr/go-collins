package collins

import (
	"testing"
)

func TestAssetTypeService_Create(t *testing.T) {
	setup()
	defer teardown()

	SetupPUT(201, "/api/assettype/SERVICE", "../tests/asset_types/create_success.json", "application/json;", t)

	_, _, err := client.AssetTypes.Create("SERVICE", "Service Asset Type")
	if err != nil {
		t.Errorf("AssetType.Create returned error: %v", err)
	}
}

func TestAssetTypeService_Create_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPUT(409, "/api/assettype/SERVICE", "../tests/asset_types/create_error.json", "application/json;", t)

	_, _, err := client.AssetTypes.Create("SERVICE", "Service Asset Type")
	if err == nil {
		t.Errorf("AssetType.Create did not return error.")
	}
}

func TestAssetTypeService_Update(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(200, "/api/assettype/SERVICE", "../tests/asset_types/update_success.json", "application/json;", t)

	_, _, err := client.AssetTypes.Update("SERVICE", "SERVICE_NEW", "Updated Service Asset Type")
	if err != nil {
		t.Errorf("AssetType.Update returned error: %v", err)
	}
}

func TestAssetTypeService_Update_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(404, "/api/assettype/SERVICE", "../tests/asset_types/update_error.json", "application/json;", t)

	_, _, err := client.AssetTypes.Update("SERVICE", "SERVICE_NEW", "Updated Service Asset Type")
	if err == nil {
		t.Errorf("AssetType.Update did not return error.")
	}
}

func TestAssetTypeService_Get(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/assettype/SERVER_NODE", "../tests/asset_types/server_node.json", "application/json;", t)

	assetType, _, err := client.AssetTypes.Get("SERVER_NODE")
	if err != nil {
		t.Errorf("AssetTypes.Get returned error: %v", err)
	}

	if assetType.Label != "Server Node" {
		t.Errorf("Asset type label: %v, want \"Server Node\".", assetType.Label)
	}
}

func TestAssetTypeService_Get_error(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(404, "/api/assettype/FOO_TYPE", "../tests/asset_types/get_error.json", "application/json;", t)

	_, _, err := client.AssetTypes.Get("FOO_TYPE")
	if err == nil {
		t.Errorf("AssetTypes.Get did not return error.")
	}
}

func TestAssetTypeService_List(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/assettypes", "../tests/asset_types/all.json", "application/json;", t)

	assetTypes, _, err := client.AssetTypes.List()
	if err != nil {
		t.Errorf("AssetTypes.List returned error: %v", err)
	}

	if len(*assetTypes) != 9 {
		t.Errorf("Number of asset types: %d, want 9", len(*assetTypes))
	}
}

func TestAssetTypeService_Delete(t *testing.T) {
	setup()
	defer teardown()

	SetupDELETE(200, "/api/assettype/SERVER_NODE", "../tests/asset_types/delete_success.json", "application/json;", t)

	_, err := client.AssetTypes.Delete("SERVER_NODE")
	if err != nil {
		t.Errorf("AssetTypes.Delete returned error: %v", err)
	}
}

func TestAssetTypeService_Delete_error(t *testing.T) {
	setup()
	defer teardown()

	SetupDELETE(404, "/api/assettype/SERVER_NODE", "../tests/asset_types/delete_error.json", "application/json;", t)

	_, err := client.AssetTypes.Delete("SERVER_NODE")
	if err == nil {
		t.Errorf("AssetTypes.Delete did not return error.")
	}
}
