package main

import (
	"fmt"
	"os"
)

var defaultsmessage string = "Usage:\n  shar config\n   \tnew [<name>]\n   \tdelete [<name>]\n   \tselect [<name>]\n   \twrite\n   \tcurrent\n  shar run <iters>\n"

func main() {
	var err error
	homedir, err = os.UserHomeDir()
	if err != nil {
		return
	}
	err = dircheckcurcfg()
	if err != nil {
		return
	}
	if len(os.Args) == 1 {
		fmt.Print(defaultsmessage)
		return
	}
	switch os.Args[1] {
	case "config":
		err = cfgmain()
	case "run":
		err = runmain()
	default:
		fmt.Println(defaultsmessage)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}
