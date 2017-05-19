package collins

import (
	"testing"
)

func TestIPAMService_IpmiPools(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/ipmi/pools", "../tests/ipam/ipmi_pools.json", "application/json;", t)

	pools, _, err := client.IPAM.IpmiPools()

	if err != nil {
		t.Errorf("IPAMService.IpmiPools returned error: %v", err)
	}

	if len(pools) != 1 {
		t.Errorf("IPAMService.IpmiPools returned %d pools, want 5.", len(pools))
	}
}
