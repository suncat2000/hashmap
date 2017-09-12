package hashmap

import (
	"testing"
)
type HashKey struct {
	key interface{}
}

func TestIntHash16(t *testing.T) {
	hashKey, bucketIdx := hashFunc(16, 123112122)

	if hashKey != 249811310353814468 {
		t.Errorf("Expected hash key 249811310353814468, got %d", hashKey)
	}

	if bucketIdx != 4 {
		t.Errorf("Expected bucket idx 4, got %d", bucketIdx)
	}
}

func TestStringHash16(t *testing.T) {
	hashKey, bucketIdx := hashFunc(16, "test_value")

	if hashKey != 11696809291008937669 {
		t.Errorf("Expected hash key 11696809291008937669, got %d", hashKey)
	}

	if bucketIdx != 5 {
		t.Errorf("Expected bucket idx 5, got %d", bucketIdx)
	}
}

func TestSliceHash16(t *testing.T) {
	hashKey, bucketIdx := hashFunc(16, []string{"test_value"})

	if hashKey != 13905799525442369900 {
		t.Errorf("Expected hash key 13905799525442369900, got %d", hashKey)
	}

	if bucketIdx != 12 {
		t.Errorf("Expected bucket idx 12, got %d", bucketIdx)
	}
}
func TestMapHash16(t *testing.T) {
	hashKey, bucketIdx := hashFunc(16, map[string]string{"key":"value"})

	if hashKey != 6432338246103250761 {
		t.Errorf("Expected hash key 6432338246103250761, got %d", hashKey)
	}

	if bucketIdx != 9 {
		t.Errorf("Expected bucket idx 9, got %d", bucketIdx)
	}
}
func TestStructHash16(t *testing.T) {
	hashKey, bucketIdx := hashFunc(16, HashKey{"key"})

	if hashKey != 2074006382534894026 {
		t.Errorf("Expected hash key 2074006382534894026, got %d", hashKey)
	}

	if bucketIdx != 10 {
		t.Errorf("Expected bucket idx 10, got %d", bucketIdx)
	}
}