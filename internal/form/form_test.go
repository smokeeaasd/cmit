package form

import "testing"

func TestValidateMessage(t *testing.T) {
	tests := []struct {
		message string
		wantErr bool
	}{
		{"", true},
		{"Add new feature", false},
	}

	for _, tt := range tests {
		err := ValidateMessage(tt.message)
		if (err != nil) != tt.wantErr {
			t.Errorf("ValidateMessage(%q) error = %v, wantErr %v", tt.message, err, tt.wantErr)
		}
	}
}
