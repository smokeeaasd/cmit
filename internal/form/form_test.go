package form

import (
	"testing"
)

func TestValidateTitle(t *testing.T) {
	tests := []struct {
		title   string
		wantErr bool
	}{
		{"", true},
		{"Add new feature", false},
		{"This title is intentionally made longer than seventy-two characters to test the limit", true},
	}

	for _, tt := range tests {
		err := ValidateTitle(tt.title)
		if (err != nil) != tt.wantErr {
			t.Errorf("ValidateTitle(%q) error = %v, wantErr %v", tt.title, err, tt.wantErr)
		}
	}
}
