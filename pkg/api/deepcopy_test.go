package api

import (
	"reflect"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	copy := testContainerService.DeepCopy()
	if !reflect.DeepEqual(testContainerService, copy) {
		t.Error("OpenShiftManagedCluster differed after DeepCopy")
	}
	copy.Tags["test"] = "update"
	if reflect.DeepEqual(testContainerService, copy) {
		t.Error("copy should differ from testContainerService after mutation")
	}
}
