package goid

import "testing"
import "regexp"

func TestGeneratedV4UUIDisRFC4122Compliant(t *testing.T) {
	actualResult := NewV4UUID()
	t.Logf("Generated: %s", actualResult)
	if ismatch, err := regexp.MatchString(V4UUIDRegex, actualResult.String()); !ismatch {
		if err != nil {
			t.Fatalf(err.Error())
		} else {
			t.Fatalf("Generated: %s which does not match regex of: %s", actualResult, V4UUIDRegex)
		}
	}
}

func TestGetUUIDFromString(t *testing.T) {
	testUUID := "40db20bd-e9ea-41d9-bed4-473b5e147582"
	resultingUUID, err := GetUUIDFromString(testUUID)
	if err != nil {
		t.Fatal(err.Error())
	} else if resultingUUID == nil {
		t.Fatal("Failed to GetUUIDFromString")
	} else {
		if resultingUUID.String() != testUUID {
			t.Fatalf("Input: %s does not match output of: %s", testUUID, resultingUUID.String())
		}
	}
}

func TestGetUUIDFromByteSlice(t *testing.T) {
	testUUID := "40db20bd-e9ea-41d9-bed4-473b5e147582"
	expectedUUID, _ := GetUUIDFromString(testUUID)
	actualUUID, _ := GetUUIDFromByteSlice(expectedUUID[:])
	if *actualUUID != *expectedUUID {
		t.Fatalf("Expected %s, got %s", expectedUUID, actualUUID)
	}
}

func TestGetVersion(t *testing.T) {
	testUUID := "40db20bd-e9ea-41d9-bed4-473b5e147582"
	resultingUUID, _ := GetUUIDFromString(testUUID)
	expectedVersion := "4"
	actualVersion := resultingUUID.GetVersion()
	if actualVersion != expectedVersion {
		t.Fatalf("Got %s, expected %s", actualVersion, expectedVersion)
	}
}

func TestEquals(t *testing.T) {
	uuidString1 := "40db20bd-e9ea-41d9-bed4-473b5e147582"
	uuidString2 := "1f2614c6-742a-4a78-a3ff-4332b09ca99b"

	uuid1, _ := GetUUIDFromString(uuidString1)
	uuid2, _ := GetUUIDFromString(uuidString1)
	if !uuid1.Equals(uuid2) {
		t.Fatalf("UUIDS %s & %s should be equal", uuid1, uuid2)
	}

	uuid3, _ := GetUUIDFromString(uuidString2)
	if uuid1.Equals(uuid3) {
		t.Fatalf("UUIDS %s & %s should not be equal", uuid1, uuid2)
	}
}
