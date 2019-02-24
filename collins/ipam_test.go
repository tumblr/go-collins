package collins

import (
	"testing"
)

func TestIPAMService_Allocate(t *testing.T) {
	setup()
	defer teardown()

	SetupPUT(201, "/api/asset/tag1/address", "../tests/ipam/allocate_success.json", "application/json;", t)

	opts := AddressAllocateOpts{}
	addresses, _, err := client.IPAM.Allocate("tag1", opts)

	if err != nil {
		t.Errorf("IPAMService.Allocate returned error: %v", err)
	}

	if len(addresses) != 1 {
		t.Errorf("IPAMService.Allocate returned %d addresses, want 1.", len(addresses))
	}
}

func TestIPAMService_Allocate_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPUT(404, "/api/asset/nosuchtag/address", "../tests/ipam/invalid_asset.json", "application/json;", t)

	opts := AddressAllocateOpts{}
	_, _, err := client.IPAM.Allocate("nosuchtag", opts)

	if err == nil {
		t.Errorf("IPAMService.Allocate did not return error on non-existing tag.")
	}
}

func TestIPAMService_Update(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(200, "/api/asset/tag1/address", "../tests/ipam/generic_success.json", "application/json;", t)

	opts := AddressUpdateOpts{}
	_, err := client.IPAM.Update("tag1", opts)

	if err != nil {
		t.Errorf("IPAMService.Update returned error: %v", err)
	}
}

func TestIPAMService_Update_error(t *testing.T) {
	setup()
	defer teardown()

	SetupPOST(404, "/api/asset/nosuchtag/address", "../tests/ipam/invalid_asset.json", "application/json;", t)

	opts := AddressUpdateOpts{}
	_, err := client.IPAM.Update("nosuchtag", opts)

	if err == nil {
		t.Errorf("IPAMService.Update did not return error on non-existing tag.")
	}
}

func TestIPAMService_Delete(t *testing.T) {
	setup()
	defer teardown()

	SetupDELETE(200, "/api/asset/tag1/addresses", "../tests/ipam/delete_success.json", "application/json;", t)

	opts := AddressDeleteOpts{}
	num, _, err := client.IPAM.Delete("tag1", opts)

	if err != nil {
		t.Errorf("IPAMService.Delete returned error: %v", err)
	}

	if num != 1 {
		t.Errorf("IPAMService.Delete returned %d, want 1.", num)
	}
}

func TestIPAMService_Delete_error(t *testing.T) {
	setup()
	defer teardown()

	SetupDELETE(404, "/api/asset/tag1/addresses", "../tests/ipam/invalid_asset", "application/json;", t)

	opts := AddressDeleteOpts{}
	_, _, err := client.IPAM.Delete("nosuchtag", opts)

	if err == nil {
		t.Errorf("IPAMService.Delete did not return error on non-existing tag.")
	}
}

func TestIPAMService_Pools(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/address/pools", "../tests/ipam/pools.json", "application/json;", t)

	pools, _, err := client.IPAM.Pools()

	if err != nil {
		t.Errorf("IPAMService.Pools returned error: %v", err)
	}

	if len(pools) != 5 {
		t.Errorf("IPAMService.Pools returned %d pools, want 5.", len(pools))
	}
}

func TestIPAMService_IPMIPools(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/ipmi/pools", "../tests/ipam/ipmi_pools.json", "application/json;", t)

	pools, _, err := client.IPAM.IPMIPools()

	if err != nil {
		t.Errorf("IPAMService.IPMIPools returned error: %v", err)
	}

	if len(pools) != 1 {
		t.Errorf("IPAMService.IPMIPools returned %d pools, want 5.", len(pools))
	}
}

func TestIPAMService_Get(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/asset/tag1/addresses", "../tests/ipam/get_success.json", "application/json;", t)

	addresses, _, err := client.IPAM.Get("tag1")

	if err != nil {
		t.Errorf("IPAMService.Get returned error: %v", err)
	}

	if len(addresses) != 1 {
		t.Errorf("IPAMService.Get returned %d addresses, want 1.", len(addresses))
	}
}

func TestIPAMService_Get_error(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(404, "/api/asset/nosuchtag/addresses", "../tests/ipam/invalid_asset.json", "application/json;", t)

	_, _, err := client.IPAM.Get("nosuchtag")

	if err == nil {
		t.Errorf("IPAMService.Get did not return error on non-existing tag.")
	}
}

func TestIPAMService_AssetFromAddress(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/asset/with/address/172.16.5.4", "../tests/ipam/asset_with_address.json", "application/json;", t)

	asset, _, err := client.IPAM.AssetFromAddress("172.16.5.4")

	if err != nil {
		t.Errorf("IPAMService.AssetFromAddress returned error: %v", err)
	}

	if asset.Metadata.Tag != "tag1" {
		t.Errorf("IPAMService.AssetFromAddress return tag %s, want tag1.", asset.Metadata.Tag)
	}
}

func TestIPAMService_AssetFromAddress_error(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(404, "/api/asset/with/address/172.16.6.1", "../tests/ipam/asset_with_address_error.json", "application/json;", t)

	_, _, err := client.IPAM.AssetFromAddress("172.16.6.1")

	if err == nil {
		t.Errorf("IPAMService.AssetFromAddress did not return error on non-existing address.")
	}
}
