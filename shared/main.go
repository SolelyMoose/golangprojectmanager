package main

import (
	"github.com/solelymoose/golangprojectmanager/cmd"
	"github.com/solelymoose/golangprojectmanager/shared"
)

func main() {
	shared.InitSharedVariables()
	cmd.Execute()
}
