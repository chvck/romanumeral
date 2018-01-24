package main

import (
	"fmt"
	"os"

	"github.com/chvck/romanumeral/run"
)

func main() {
	r := run.NewRunner()
	if roman, err := r.Run(os.Args[1:]...); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*roman)
	}
}
