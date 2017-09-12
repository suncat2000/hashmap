package hashmap

import (
	"testing"
)

func TestIntHash16(t *testing.T) {
	hashKey, bucketIdx := hashFunc(16, 123112122)

	if hashKey != 4 {
		t.Errorf("Expected hash key 4, got %d", hashKey)
	}

	if bucketIdx != 4 {
		t.Errorf("Expected bucket idx 4, got %d", bucketIdx)
	}
}