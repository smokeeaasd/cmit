package commit

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/smokeeaasd/cmit/internal/form"
	"github.com/smokeeaasd/cmit/internal/utils"
)

var GitWorkDir string

var execCommand = exec.Command

func ExecuteCommit(extraArgs []string) {
	if !form.Confirm {
		fmt.Println("Commit aborted. 👋")
		os.Exit(0)
	}

	commitPrefix, ok := utils.CommitLabels[form.CommitType]
	if !ok {
		log.Fatal("Invalid commit type selected")
	}

	commitMessage := utils.BuildCommitMessage(commitPrefix, form.Scope, form.Title)
	if form.Description != "" {
		commitMessage += "\n\n" + form.Description
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
