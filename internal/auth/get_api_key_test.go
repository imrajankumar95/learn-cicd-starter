package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		authHeader string
		want       string
		wantErr    bool
	}{
		// Add your test cases here
		{
			name:       "valid api key",
			authHeader: "ApiKey 12345",
			want:       "12345",
			wantErr:    false,
		},
		{
			name:       "missing authorization header",
			authHeader: "",
			want:       "",
			wantErr:    true,
		},
		{
			name:       "malformed authorization header",
			authHeader: "Bearer 12345",
			want:       "",
			wantErr:    true,
		},
	}

	// Loop goes here
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// create headers
			headers := http.Header{}
			if tc.authHeader != "" {
				headers.Set("Authorization", tc.authHeader)
			}

			// call GetAPIKey
			var got string
			got, err := GetAPIKey(headers)

			// check error expectation
			if err != nil && !tc.wantErr {
				t.Errorf("unexcepted error: %v", err)
			}

			if err == nil && tc.wantErr {
				t.Errorf("expected error, got nil")
			}

			// check got vs want
			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
