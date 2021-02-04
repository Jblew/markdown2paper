package main
import "testing"

func TestGetCitationKeysFromText(t *testing.T) {
	got := getCitationKeysFromText("# H1[@key1]\n\nC1[@key2]\n\n#H2\n\nC2[@key1][@key3]")
	if len(got) != 3 {
			t.Errorf("Should find 3 citation keys, got %d, %+v", len(got), got)
			return
	}

	if got[0] != "key1" {
		t.Errorf("citation[0] should be %s, got %s", "key1", got[0])
	}

	if got[1] != "key2" {
		t.Errorf("citation[1] should be %s, got %s", "key2", got[1])
	}

	if got[2] != "key3" {
		t.Errorf("citation[2] should be %s, got %s", "key3", got[2])
	}
}
