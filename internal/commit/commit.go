package commit

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/smokeeaasd/cmit/internal/form"
)

var GitWorkDir string

var commitLabels = map[string]string{
	"feat":     "💡 feat",
	"fix":      "🐞 fix",
	"build":    "📦 build",
	"chore":    "🔧 chore",
	"ci":       "🤖 ci",
	"docs":     "📝 docs",
	"style":    "🎨 style",
	"refactor": "🔨 refactor",
	"perf":     "🚀 perf",
	"test":     "✅ test",
}

var execCommand = exec.Command

func ExecuteCommit(extraArgs []string) {
	commitPrefix, ok := commitLabels[form.CommitType]
	if !ok {
		log.Fatal("Invalid commit type selected")
	}

	var commitMessage string
	if form.Scope == "" {
		commitMessage = fmt.Sprintf("%s: %s", commitPrefix, form.Message)
	} else {
		commitMessage = fmt.Sprintf("%s(%s): %s", commitPrefix, form.Scope, form.Message)
	}

	escapedMessage := strings.ReplaceAll(commitMessage, "\"", "\\\"")
	fmt.Printf("\n✅ Commit message: \n%s\n\n", commitMessage)

	cmdArgs := []string{"commit", "-m", escapedMessage}
	if len(extraArgs) > 0 {
		cmdArgs = append(cmdArgs, extraArgs...)
	}

	cmd := execCommand("git", cmdArgs...)
	if GitWorkDir != "" {
		cmd.Dir = GitWorkDir
	} else {
		cmd.Dir = "."
	}

	out, err := cmd.CombinedOutput()
	outputStr := string(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[31m%s\033[0m", outputStr)
		log.Fatalf("Failed to execute git commit: %v", err)
	}

	fmt.Print(outputStr)
}
