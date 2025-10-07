package main

import (
	"fmt"
	"os"
	"strings"
	todo "todomanager"
)

const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	if err := l.Get(todoFileName); err != nil {
		// if !os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		// }
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Print(item.Task)
		}
	default:
		item := strings.Join(os.Args[1:], " ")
		l.Add(item)

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	}
}
