package hermes

import (
	"encoding/base64"
	"errors"
	"fmt"

	pasetokit "github.com/heliantheon/aegis-go/utilities/paseto"
)

var ErrApplicationSeedNotFound = errors.New("application seed not found")

const applicationClientSecretPurpose = "basic"

func deriveApplicationClientSecret(seed []byte) (string, error) {
	if len(seed) != 48 {
		return "", fmt.Errorf("invalid application seed length: got %d, want 48", len(seed))
	}

	parsed, err := pasetokit.ParseSeed(seed)
	if err != nil {
		return "", fmt.Errorf("parse application seed: %w", err)
	}
	derived := parsed.Derive(applicationClientSecretPurpose)
	return base64.RawURLEncoding.EncodeToString(derived), nil
}
