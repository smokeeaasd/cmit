package utils

import (
	"fmt"
)

var CommitLabels = map[string]string{
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

func BuildCommitMessage(prefix string, scope string, message string) string {
	if scope == "" {
		return fmt.Sprintf("%s: %s", prefix, message)
	}
	return fmt.Sprintf("%s(%s): %s", prefix, scope, message)
}
