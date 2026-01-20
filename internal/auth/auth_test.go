package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {

	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name: "Success Scenario when valid auth header is given",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid-test-api-key"},
			},
			want:    "valid-test-api-key",
			wantErr: false,
		},
		{
			name: "Faliure Scenario when valid auth header is given",
			headers: http.Header{
				"Content-Type": []string{"application/json"},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		got, err := GetAPIKey(tt.headers)

		if (err != nil) != tt.wantErr {
			t.Errorf("GetAPIKey() = %v, want %v", err, tt.wantErr)
			return
		}

		if got != tt.want {
			t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
		}
	}

}
