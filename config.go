package main

import (
	"fmt"
	"os"
)

func cfgwrite() (err error) {
	var cfg config
	err = readcurcfg()
	if err != nil {
		return
	}
	fmt.Println("Input 1st solution's filepath: ")
	fmt.Scan(&cfg.Solve1)
	fmt.Println("Input checker/2nd solution's filepath: ")
	fmt.Scan(&cfg.Solve2)
	fmt.Println("Checker? Y/n ")
	{
		var temp string
		fmt.Scan(&temp)
		if temp == "Y" || temp == "y" {
			cfg.Checker = true
		}
	}
	fmt.Println("Input generator's filepath: ")
	fmt.Scan(&cfg.Generator)
	fmt.Println("Input your compiler (clang++ or g++): ")
	fmt.Scan(&cfg.Compiler)

	_, err = os.OpenFile(string(homedir+"/shar/"+curcfg), os.O_TRUNC, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	cfgfile, err := os.OpenFile(string(homedir+"/shar/"+curcfg), os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = cfgwritefile(cfgfile, cfg)
	return
}

func cfgnew() (err error) {
	var newcfg string
	if len(os.Args) < 4 {
		fmt.Println("Input name of config")
		fmt.Scan(&newcfg)
	} else {
		newcfg = os.Args[3]
	}
	_, err = os.Create(string(homedir + "/shar/" + newcfg))
	if os.IsExist(err) {
		fmt.Println("File is exist")
	}
	if err != nil {
		fmt.Println(err)
	}
	return
}

func cfgdelete() (err error) {
	var delcfg string
	if len(os.Args) < 4 {
		fmt.Println("Input name of config")
		fmt.Scan(&delcfg)
	} else {
		delcfg = os.Args[3]
	}
	err = os.Remove(string(homedir + "/shar/" + delcfg))
	if os.IsNotExist(err) {
		fmt.Println("File isn't exist")
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func cfgselect() (err error) {
	var selcfg string
	if len(os.Args) < 4 {
		fmt.Println("Input name of config")
		fmt.Scan(&selcfg)
	} else {
		selcfg = os.Args[3]
	}
	err = writecurcfg(selcfg)
	return
}

func cfgmain() (err error) {
	if len(os.Args) < 3 {
		fmt.Println("Not enough args")
		return
	}
	switch os.Args[2] {
	case "new":
		err = cfgnew()
	case "delete":
		err = cfgdelete()
	case "select":
		err = cfgselect()
	case "write":
		err = cfgwrite()
	case "current":
		err = readcurcfg()
		if err != nil {
			return
		}
		fmt.Println("Current config is", curcfg)
	default:
		fmt.Println("Unknown command")
	}
	if err != nil {
		return
	}
	return
}
