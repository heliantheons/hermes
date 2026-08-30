package dto

import (
	"encoding/json"
	"testing"
)

func TestCreateResourceRequestsPreserveLogoURL(t *testing.T) {
	t.Parallel()

	const logoURL = "https://assets.example.com/logo.png"
	tests := []struct {
		name    string
		request any
	}{
		{name: "application", request: &ApplicationCreateRequest{}},
		{name: "service", request: &ServiceCreateRequest{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := json.Unmarshal([]byte(`{"logo_url":"`+logoURL+`"}`), tt.request); err != nil {
				t.Fatalf("unmarshal create request: %v", err)
			}

			encoded, err := json.Marshal(tt.request)
			if err != nil {
				t.Fatalf("marshal create request: %v", err)
			}
			var fields map[string]any
			if err := json.Unmarshal(encoded, &fields); err != nil {
				t.Fatalf("unmarshal encoded request: %v", err)
			}
			if got, ok := fields["logo_url"]; !ok || got != logoURL {
				t.Fatalf("logo_url = %v, present = %t; want %q", got, ok, logoURL)
			}
		})
	}
}
