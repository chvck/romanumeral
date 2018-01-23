package main

import (
	"fmt"
	"os"

	"github.com/chvck/romanumeral"
)

func main() {
	r := romanumeral.NewRunner()
	if roman, err := r.Run(os.Args[1:]...); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(*roman)
	}
}
