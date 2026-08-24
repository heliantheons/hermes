package hermes

import (
	"encoding/base64"
	"errors"
	"fmt"

	pasetokit "github.com/heliantheon/aegis-go/utilities/paseto"
)

var (
	ErrApplicationSeedNotFound      = errors.New("application seed not found")
	ErrUnsupportedApplicationSecret = errors.New("unsupported application secret type")
)

type ApplicationSecretType string

const ApplicationSecretTypeClientSecret ApplicationSecretType = "client-secret"

func ParseApplicationSecretType(value string) (ApplicationSecretType, error) {
	secretType := ApplicationSecretType(value)
	if secretType != ApplicationSecretTypeClientSecret {
		return "", ErrUnsupportedApplicationSecret
	}
	return secretType, nil
}

func deriveApplicationSecret(seed []byte, secretType ApplicationSecretType) (string, error) {
	if _, err := ParseApplicationSecretType(string(secretType)); err != nil {
		return "", err
	}
	if len(seed) != 48 {
		return "", fmt.Errorf("invalid application seed length: got %d, want 48", len(seed))
	}

	parsed, err := pasetokit.ParseSeed(seed)
	if err != nil {
		return "", fmt.Errorf("parse application seed: %w", err)
	}
	derived := parsed.Derive(string(secretType))
	return base64.RawURLEncoding.EncodeToString(derived), nil
}
