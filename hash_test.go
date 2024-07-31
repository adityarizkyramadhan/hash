// file: hash/hash_test.go
package hash

import (
	"testing"
)

func TestHashID(t *testing.T) {
	secret := "examplekey123456" // Secret key for the hash
	id := 1234                   // Sample ID for hashing

	hashedID, err := HashID(id, secret)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if hashedID == "" {
		t.Fatalf("Expected non-empty hash, got empty string")
	}

	// Additional test: check that different IDs result in different hashes
	id2 := 5678
	hashedID2, err := HashID(id2, secret)
	if err != nil {
		t.Fatalf("Expected no error for second ID, got %v", err)
	}

	if hashedID == hashedID2 {
		t.Fatalf("Expected different hashes for different IDs, got same hash: %v", hashedID)
	}
}

func TestDecryptID(t *testing.T) {
	id := 1234
	key := "example"
	secret, err := HashID(id, key)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	decryptedID, err := DecryptID(secret, key)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if decryptedID != id {
		t.Fatalf("Expected %v, got %v", id, decryptedID)
	}
}
