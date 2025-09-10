package form

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
	"github.com/smokeeaasd/cmit/internal/utils"
)

var (
	CommitType  string
	Title       string
	Description string
	Scope       string
	Confirm     bool
)

func ValidateTitle(title string) error {
	if title == "" {
		return errors.New("commit title cannot be empty")
	}
	if len(title) > 72 {
		return errors.New("commit title should be 72 characters or less")
	}
	return nil
}

func CreateForm(detailed bool) *huh.Form {
	if detailed {
		templateLines := []string{
			"<Detailed description explaining what changed and why>",
			"<The body can have multiple lines, usage examples, context, links, etc.>",
			"",
			"BREAKING CHANGE: <description of the breaking change>",
			"",
			"Closes #<issue number>",
			"Fixes #<issue number>",
			"Refs #<issue number>",
			"Co-authored-by: Name <email>",
			"Co-authored-by: Another Name <email>",
			"Reviewed-by: Name <email>",
			"Signed-off-by: Name <email>",
		}

		Description = ""
		for i, line := range templateLines {
			Description += line
			if i < len(templateLines)-1 {
				Description += "\n"
			}
		}
	}

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
			huh.NewInput().
				Title("✏️\u00A0 Commit title").
				Description("A short summary of the change (max 72 characters)").
				Validate(ValidateTitle).
				Value(&Title),
		),
		huh.NewGroup(
			huh.NewText().
				Title("📝 Commit description (optional)").
				Description("A more detailed explanation of the change").
				CharLimit(1000).
				Value(&Description).WithHeight(14),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("🧐 Are you sure?").
				DescriptionFunc(func() string {
					commitPrefix := utils.CommitLabels[CommitType]

					fullMessage := utils.BuildCommitMessage(commitPrefix, Scope, Title)

					if Description != "" {
						lines := 1
						for _, c := range Description {
							if c == '\n' {
								lines++
							}
						}
						fullMessage += fmt.Sprintf("\n(+%d lines)", lines)
					}

					return fmt.Sprintf("%s\nConfirm commit creation.", fullMessage)
				}, nil).
				Affirmative("Yes!").
				Negative("No.").
				Value(&Confirm),
		),
	).WithKeyMap(KeyMap)
}
