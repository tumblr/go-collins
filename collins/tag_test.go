package collins

import (
	"testing"
)

func TestTagService_List(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/tags", "../tests/tags/all.json", "application/json;", t)

	tags, _, err := client.Tags.List()
	if err != nil {
		t.Errorf("Tags.List returned error: %v", err)
	}

	if len(tags) != 35 {
		t.Errorf("Number of tags: %d, want 35", len(tags))
	}
}

func TestTagService_Values(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(200, "/api/tag/CPU_CORES", "../tests/tags/values_success.json", "application/json;", t)

	vals, _, err := client.Tags.Values("CPU_CORES")
	if err != nil {
		t.Errorf("Tags.Values returned error: %v", err)
	}

	want := []string{"12", "8", "4", "10", "6", "1"}

	if len(vals) != len(want) {
		t.Errorf("Number of values: %d, want %d", len(vals), len(want))
	}

	eq := true
	for i, v := range want {
		if vals[i] != v {
			eq = false
		}
	}

	if !eq {
		t.Errorf("Wrong values returned: %v, want %v.", vals, want)
	}
}

func TestTagService_Values_error(t *testing.T) {
	setup()
	defer teardown()

	SetupGET(404, "/api/tag/DOESNOTEXIST", "../tests/tags/values_error.json", "application/json;", t)

	_, _, err := client.Tags.Values("DOESNOTEXIST")
	if err == nil {
		t.Errorf("Tags.Values did not return error.")
	}
}
