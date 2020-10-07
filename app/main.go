package main

import (
	"os"

	benchconsole "github.com/pip-benchmark/pip-benchmark-go/console"
)

func main() {
	benchconsole.Run(os.Args)
}
