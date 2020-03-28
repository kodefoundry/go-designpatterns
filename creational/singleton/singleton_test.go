package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetUnsafeInstance()

	if &counter1 == nil {
		t.Error("Expected pointer to an UnsafeSingleton, got nil.")
	}
	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("Expected initial value is 1, but received %d\n", currentCount)
	}

	counter2 := GetUnsafeInstance()
	if counter2 != expectedCounter {
		t.Error("Unique instance expected on repeated calls, but received different instances")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("Inconsistant counter, expected 2 but received %d.\n", currentCount)
	}
}
