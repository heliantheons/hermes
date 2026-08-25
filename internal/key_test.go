package hermes

import (
	"bytes"
	"encoding/base64"
	"testing"

	cryptoutil "github.com/heliantheon/common/crypto"
)

func TestEncryptApplicationSeedRoundTrip(t *testing.T) {
	t.Parallel()

	dbEncryptionKey := bytes.Repeat([]byte{0x21}, 32)
	seed := bytes.Repeat([]byte{0x42}, 48)
	const ownerID = "grafana"

	encoded, err := encryptApplicationSeed(dbEncryptionKey, seed, ownerID)
	if err != nil {
		t.Fatalf("encryptApplicationSeed() error = %v", err)
	}
	ciphertext, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		t.Fatalf("decode encrypted seed: %v", err)
	}
	decrypted, err := cryptoutil.DecryptAESGCM(dbEncryptionKey, ciphertext, ownerID)
	if err != nil {
		t.Fatalf("decrypt encrypted seed: %v", err)
	}
	if !bytes.Equal(decrypted, seed) {
		t.Fatal("decrypted application seed does not match the original")
	}
}

func TestEncryptApplicationSeedRejectsInvalidDatabaseKey(t *testing.T) {
	t.Parallel()

	if _, err := encryptApplicationSeed(make([]byte, 48), make([]byte, 48), "grafana"); err == nil {
		t.Fatal("encryptApplicationSeed() accepted a non-AES-256 database key")
	}
}
