package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	input1 := make(http.Header)
	input1.Set("Authorization", "ApiKey testtok")
	input2 := make(http.Header)
	input2.Set("Authorization", "testtok")
	input3 := make(http.Header)

	tests := map[string]struct {
		input   http.Header
		want    string
		wantErr error
	}{
		"contains auth": {
			input:   input1,
			want:    "testtok",
			wantErr: nil,
		},
		"malformed header": {
			input:   input2,
			want:    "",
			wantErr: ErrMalformedAuthHeader,
		},
		"missing auth": {
			input:   input3,
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)

			// unexpected err
			if !errors.Is(err, tc.wantErr) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
			}

			// unexpected got
			if tc.wantErr == nil && got != tc.want {
				t.Errorf("GetAPIKey() got = %v, want %v", got, tc.want)
			}
		})
	}
}
