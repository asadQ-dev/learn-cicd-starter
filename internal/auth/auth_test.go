package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		header  http.Header
		wantKey string
		wantErr bool
	}{
		{
			name:    "valid header",
			header:  http.Header{"Authorization": []string{"ApiKey test-key"}},
			wantKey: "test-key",
			wantErr: false,
		},
		{
			name:    "no space",
			header:  http.Header{"Authorization": []string{"ApiKeyOneString"}},
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "no header",
			header:  http.Header{},
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "wrong prefix",
			header:  http.Header{"Authorization": []string{"Bearer test-key"}},
			wantKey: "",
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(testCase.header)
			if (err != nil) != testCase.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if gotKey != testCase.wantKey {
				t.Errorf("GetAPIKey() = %v, want %v", gotKey, testCase.wantKey)
			}
		})
	}
}
