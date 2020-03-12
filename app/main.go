package main

import (
	benchconsole "github.com/pip-benchmark/pip-benchmark-go/console"
)

func main() {
	//benchconsole.Run(os.Args)
	benchconsole.Run([]string{"-e"})
}
