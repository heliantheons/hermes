package hermes

import (
	"bytes"
	"testing"
)

func TestDeriveApplicationClientSecret(t *testing.T) {
	t.Parallel()

	seed := bytes.Repeat([]byte{0x42}, 48)
	first, err := deriveApplicationClientSecret(seed)
	if err != nil {
		t.Fatalf("deriveApplicationClientSecret() error = %v", err)
	}
	second, err := deriveApplicationClientSecret(seed)
	if err != nil {
		t.Fatalf("deriveApplicationClientSecret() second error = %v", err)
	}
	if first != second {
		t.Fatal("deriveApplicationClientSecret() is not deterministic")
	}
	if len(first) != 43 {
		t.Fatalf("derived secret length = %d, want 43", len(first))
	}
	if want := "G24kZfyp4aI6VIrms1ghmjMMuMA0vvEWkR5pEAU_UUY"; first != want {
		t.Fatalf("derived secret = %q, want protocol test vector %q", first, want)
	}

}

func TestDeriveApplicationClientSecretRejectsInvalidSeed(t *testing.T) {
	t.Parallel()

	if _, err := deriveApplicationClientSecret(make([]byte, 32)); err == nil {
		t.Fatal("deriveApplicationClientSecret() accepted an invalid seed length")
	}
}
