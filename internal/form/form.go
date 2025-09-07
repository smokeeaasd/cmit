package form

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
	"github.com/smokeeaasd/cmit/internal/utils"
)

var (
	CommitType string
	Message    string
	Scope      string
	Confirm    bool
)

func ValidateMessage(msg string) error {
	if msg == "" {
		return errors.New("commit message cannot be empty")
	}
	return nil
}

func CreateForm() *huh.Form {
	KeyMap := huh.NewDefaultKeyMap()

	KeyMap.Text = huh.TextKeyMap{
		Prev:    key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "back")),
		Next:    key.NewBinding(key.WithKeys("alt+enter"), key.WithHelp("alt+enter", "next")),
		NewLine: key.NewBinding(key.WithKeys("enter", "ctrl+j"), key.WithHelp("enter / ctrl+j", "new line")),
	}

	KeyMap.Input = huh.InputKeyMap{
		Prev:   key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "back")),
		Next:   key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "next")),
		Submit: key.NewBinding(key.WithKeys("shift+enter"), key.WithHelp("shift+enter", "submit")),
	}

	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("🧐 What kind of commit is this?").
				Description("Select the type of change you are committing").
				Options(
					huh.NewOption("💡 feat - New feature", "feat"),
					huh.NewOption("🐞 fix - Bug fix", "fix"),
					huh.NewOption("📦 build - Build system or dependencies", "build"),
					huh.NewOption("🔧 chore - Other changes that don't modify src or test files", "chore"),
					huh.NewOption("🤖 ci - CI configuration", "ci"),
					huh.NewOption("📝 docs - Documentation changes", "docs"),
					huh.NewOption("🎨 style - Code style changes (formatting, missing semi colons, etc)", "style"),
					huh.NewOption("🔨 refactor - Code refactor", "refactor"),
					huh.NewOption("🚀 perf - Performance improvements", "perf"),
					huh.NewOption("✅ test - Test additions/updates", "test"),
				).
				Value(&CommitType),
		),
		huh.NewGroup(
			huh.NewInput().
				Title("📂 Specify the scope (optional)").
				Description("Press Enter to skip if there's no scope").
				Value(&Scope),
		),
		huh.NewGroup(
			huh.NewText().
				Title("✏️\u00A0 Type your commit message").
				Description("Enter a concise and descriptive commit message").
				CharLimit(1000).
				Validate(ValidateMessage).
				Value(&Message),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("🧐 Are you sure?").
				DescriptionFunc(func() string {
					commitPrefix := utils.CommitLabels[CommitType]
					var description = fmt.Sprintf("%s\nConfirm commit creation.", utils.BuildCommitMessage(commitPrefix, Scope, Message))

					return description
				}, nil).
				Affirmative("Yes!").
				Negative("No.").
				Value(&Confirm),
		),
	).WithKeyMap(KeyMap)
}
