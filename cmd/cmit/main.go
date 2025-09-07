package main

import (
	"log"

	"github.com/smokeeaasd/cmit/internal/commit"
	"github.com/smokeeaasd/cmit/internal/form"
)

func main() {
	form := form.CreateForm()

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	commit.ExecuteCommit()
}
