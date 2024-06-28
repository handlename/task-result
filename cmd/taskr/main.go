package main

import (
	"os"

	"github.com/handlename/task-result/cli"
)

func main() {
	os.Exit(int(cli.Run()))
}
