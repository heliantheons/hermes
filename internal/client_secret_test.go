package hermes

import (
	"bytes"
	"testing"
)

func TestDeriveApplicationSecret(t *testing.T) {
	t.Parallel()

	seed := bytes.Repeat([]byte{0x42}, 48)
	first, err := deriveApplicationSecret(seed, ApplicationSecretTypeBasic)
	if err != nil {
		t.Fatalf("deriveApplicationSecret() error = %v", err)
	}
	second, err := deriveApplicationSecret(seed, ApplicationSecretTypeBasic)
	if err != nil {
		t.Fatalf("deriveApplicationSecret() second error = %v", err)
	}
	if first != second {
		t.Fatal("deriveApplicationSecret() is not deterministic")
	}
	if len(first) != 43 {
		t.Fatalf("derived secret length = %d, want 43", len(first))
	}
	if want := "G24kZfyp4aI6VIrms1ghmjMMuMA0vvEWkR5pEAU_UUY"; first != want {
		t.Fatalf("derived secret = %q, want protocol test vector %q", first, want)
	}

}

func TestDeriveApplicationSecretRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	if _, err := deriveApplicationSecret(make([]byte, 32), ApplicationSecretTypeBasic); err == nil {
		t.Fatal("deriveApplicationSecret() accepted an invalid seed length")
	}
	if _, err := deriveApplicationSecret(make([]byte, 48), ApplicationSecretType("sign")); err == nil {
		t.Fatal("deriveApplicationSecret() accepted an unsupported secret type")
	}
}

func TestParseApplicationSecretType(t *testing.T) {
	t.Parallel()

	got, err := ParseApplicationSecretType("basic")
	if err != nil || got != ApplicationSecretTypeBasic {
		t.Fatalf("ParseApplicationSecretType(basic) = %q, %v", got, err)
	}
	if _, err := ParseApplicationSecretType("encrypt"); err == nil {
		t.Fatal("ParseApplicationSecretType() accepted an internal key purpose")
	}
}
