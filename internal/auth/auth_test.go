package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectedAPIKey string
		expectError    bool
	}{
		{
			name:           "No Authorization Header",
			headers:        http.Header{},
			expectedAPIKey: "",
			expectError:    true,
		},
		{
			name: "Malformed Authorization Header",
			headers: http.Header{
				"Authorization": []string{"InvalidHeader"},
			},
			expectedAPIKey: "",
			expectError:    true,
		},
		{
			name: "Valid Authorization Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid-api-key"},
			},
			expectedAPIKey: "valid-api-key",
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
				if apiKey != tt.expectedAPIKey {
					t.Errorf("expected API key %s but got %s", tt.expectedAPIKey, apiKey)
				}
			}
		})
	}
}