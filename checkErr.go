package checkErr

import (
	"os"
	"fmt"
)

func Exit(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
}

func Notify(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
}

func ExitWithCode(e error, c int) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(c)
	}
}
