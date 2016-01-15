package collins

import (
	"testing"
)

func TestAssetService_Find(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(201, "/api/assets", "../tests/assets/find_success.json", "application/json;", t)

	opts := AssetFindOpts{Type: "SERVER_NODE"}
	assets, _, err := client.Assets.Find(&opts)
	if err != nil {
		t.Errorf("Asset.Find returned error: %v", err)
	}

	if len(assets) != 2 {
		t.Fatalf("Number of assets returned: %d, want 2", len(assets))
	}

	if assets[0].Metadata.Tag != "tag30" {
		t.Errorf("First asset tag: %s, want tag30", assets[0].Metadata.Tag)
	}
}

func TestAssetService_FindSimilar(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/asset/tag1/similar", "../tests/assets/similar.json", "application/json;", t)

	assets, _, err := client.Assets.FindSimilar("tag1", nil)
	if err != nil {
		t.Errorf("Asset.FindSimilar returned error: %v", err)
	}

	if len(assets) != 3 {
		t.Errorf("Asset.FindSimilar returned %d assets, want 3.", len(assets))
	}
}
