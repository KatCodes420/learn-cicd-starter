package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		header    string
		want      string
		wantError bool
	}{
		{
			name:   "valid api key",
			header: "ApiKey abc123",
			want:   "abc123",
		},
		{
			name:      "missing authorization header",
			wantError: true,
		},
		{
			name:      "malformed authorization header",
			header:    "Bearer abc123",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}

			if tt.header != "" {
				headers.Set("Authorization", tt.header)
			}

			got, err := GetAPIKey(headers)

			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tt.want)
			}

			if tt.wantError && err == nil {
				t.Errorf("GetAPIKey() expected an error, got nil")
			}

			if !tt.wantError && err != nil {
				t.Errorf("GetAPIKey() unexpected error: %v", err)
			}
		})
	}
}
