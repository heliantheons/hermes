package hermes

import (
	"bytes"
	"testing"
)

func TestDeriveApplicationSecret(t *testing.T) {
	t.Parallel()

	seed := bytes.Repeat([]byte{0x42}, 48)
	first, err := deriveApplicationSecret(seed, ApplicationSecretTypeClientSecret)
	if err != nil {
		t.Fatalf("deriveApplicationSecret() error = %v", err)
	}
	second, err := deriveApplicationSecret(seed, ApplicationSecretTypeClientSecret)
	if err != nil {
		t.Fatalf("deriveApplicationSecret() second error = %v", err)
	}
	if first != second {
		t.Fatal("deriveApplicationSecret() is not deterministic")
	}
	if len(first) != 43 {
		t.Fatalf("derived secret length = %d, want 43", len(first))
	}
	if want := "5TSHOjfPvWkRKtg_IBoVtw0dq-8YKt6rVjUFxdo3J1k"; first != want {
		t.Fatalf("derived secret = %q, want protocol test vector %q", first, want)
	}

}

func TestDeriveApplicationSecretRejectsInvalidInput(t *testing.T) {
	t.Parallel()

	if _, err := deriveApplicationSecret(make([]byte, 32), ApplicationSecretTypeClientSecret); err == nil {
		t.Fatal("deriveApplicationSecret() accepted an invalid seed length")
	}
	if _, err := deriveApplicationSecret(make([]byte, 48), ApplicationSecretType("sign")); err == nil {
		t.Fatal("deriveApplicationSecret() accepted an unsupported secret type")
	}
}

func TestParseApplicationSecretType(t *testing.T) {
	t.Parallel()

	got, err := ParseApplicationSecretType("client-secret")
	if err != nil || got != ApplicationSecretTypeClientSecret {
		t.Fatalf("ParseApplicationSecretType(client-secret) = %q, %v", got, err)
	}
	if _, err := ParseApplicationSecretType("basic"); err == nil {
		t.Fatal("ParseApplicationSecretType() exposed the internal KDF purpose")
	}
	if _, err := ParseApplicationSecretType("encrypt"); err == nil {
		t.Fatal("ParseApplicationSecretType() accepted an internal key purpose")
	}
}
