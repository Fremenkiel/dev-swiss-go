package handlers

import (
	"fmt"
	"os"
)

type Flag struct {
	Flag				string
	Shorthand		string
	Description	string
}

func (f *Flag) Print() {
	fmt.Fprintf(os.Stdout, "	-%s, --%s - %s\n", f.Shorthand, f.Flag, f.Description)
}
