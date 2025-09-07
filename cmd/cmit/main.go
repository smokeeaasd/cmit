package main

import (
	"fmt"
	"log"

	"github.com/smokeeaasd/cmit/internal/commit"
	"github.com/smokeeaasd/cmit/internal/form"
	"github.com/smokeeaasd/cmit/internal/version"
)

func main() {
	fmt.Println(version.Version)
	form := form.CreateForm()

	if err := form.Run(); err != nil {
		log.Fatal(err)
	}

	commit.ExecuteCommit()
}
