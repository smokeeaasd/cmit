package commit

import (
	"fmt"
	"testing"

	"github.com/smokeeaasd/cmit/internal/form"
)

func buildCmdArgs(message, scope, commitType string, extraArgs []string) ([]string, error) {
	prefix, ok := commitLabels[commitType]
	if !ok {
		return nil, fmt.Errorf("invalid commit type: %s", commitType)
	}

	var commitMessage string
	if scope == "" {
		commitMessage = fmt.Sprintf("%s: %s", prefix, message)
	} else {
		commitMessage = fmt.Sprintf("%s(%s): %s", prefix, scope, message)
	}

	cmdArgs := []string{"commit", "-m", commitMessage}
	if len(extraArgs) > 0 {
		cmdArgs = append(cmdArgs, extraArgs...)
	}
	return cmdArgs, nil
}

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
		args, err := buildCmdArgs(tt.message, tt.scope, tt.commitType, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if args[2] != tt.expected {
			t.Errorf("expected '%s', got '%s'", tt.expected, args[2])
		}
	}
}

func TestCommitWithExtraArgs(t *testing.T) {
	form.CommitType = "feat"
	form.Scope = "core"
	form.Message = "Add feature"
	extra := []string{"--no-verify", "--amend"}

	args, err := buildCmdArgs(form.Message, form.Scope, form.CommitType, extra)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedMessage := "💡 feat(core): Add feature"
	if args[2] != expectedMessage {
		t.Errorf("expected commit message '%s', got '%s'", expectedMessage, args[2])
	}

	if len(args) != 5 || args[3] != "--no-verify" || args[4] != "--amend" {
		t.Errorf("extra args not appended correctly: %v", args[3:])
	}
}

func TestInvalidCommitType(t *testing.T) {
	form.CommitType = "invalid"
	form.Scope = ""
	form.Message = "Some message"

	_, err := buildCmdArgs(form.Message, form.Scope, form.CommitType, nil)
	if err == nil {
		t.Errorf("expected error for invalid commit type")
	}
}
