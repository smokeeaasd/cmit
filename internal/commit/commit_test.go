package commit

import (
	"fmt"
	"testing"

	"github.com/smokeeaasd/cmit/internal/form"
)

func TestCommitMessageFormatting(t *testing.T) {
	tests := []struct {
		commitType string
		scope      string
		message    string
		expected   string
	}{
		{"feat", "", "Add new feature", "💡 feat: Add new feature"},
		{"fix", "ui", "Fix button alignment", "🐞 fix(ui): Fix button alignment"},
		{"docs", "", "Update README", "📝 docs: Update README"},
	}

	for _, tt := range tests {
		form.CommitType = tt.commitType
		form.Scope = tt.scope
		form.Message = tt.message

		prefix, ok := commitLabels[tt.commitType]
		if !ok {
			t.Fatalf("invalid commit type: %s", tt.commitType)
		}

		var commitMessage string
		if tt.scope == "" {
			commitMessage = fmt.Sprintf("%s: %s", prefix, tt.message)
		} else {
			commitMessage = fmt.Sprintf("%s(%s): %s", prefix, tt.scope, tt.message)
		}

		if commitMessage != tt.expected {
			t.Errorf("expected '%s', got '%s'", tt.expected, commitMessage)
		}
	}
}
