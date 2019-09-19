package license

import "testing"

func TestUUID(t *testing.T) {
	uuid, err := NewUUID()
	if err != nil {
		t.Fatalf("Shouldn't have an error! %v", err)
	}
	if uuid == BlankUUID {
		t.Fatal("UUID was blank")
	}
}
