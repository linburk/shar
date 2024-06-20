package main

import (
	"fmt"
	"os"
)

var defaultsmessage string = "Shar - program for stress-testing\n\t-config\n\t\t-new\n\t\tcreate new config file\n\t\t-delete\n\t\tdelete config file\n\t\t-select\n\t\tselect config file\n\t\t-write\n\t\tinput current config file\n\t\t-current\n\t\tprint current config\n\n\t-run [iters]\n\tstart stress-test with current config"

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
