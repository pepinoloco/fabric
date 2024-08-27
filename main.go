package main

import (
	"fmt"
	"os"

	"fabric/cli"
)

func main() {
	_, err := cli.Cli()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
